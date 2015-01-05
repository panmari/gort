package accelerators

import (
	math "github.com/barnex/fmath"
	"github.com/ungerik/go3d/vec3"
	"intersectables"
	"util"
)

type BSPAccelerator struct {
	n, max_depth int
	root         *BSPNode
	box          vec3.Box
}

type BSPNode struct {
	inters      []util.Intersectable
	box         vec3.Box
	splitAxis   Axis
	splitPos    float32
	left, right *BSPNode
}

func (n *BSPNode) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func newBSPNode(box vec3.Box, axis Axis) *BSPNode {
	return new(BSPNode)
}

const MIN_NR_PRIMITIVES = 5

type Axis int

const (
	X Axis = iota
	Y
	Z
)

func NewBSPAccelerator(a *intersectables.Aggregate) *BSPAccelerator {
	acc := new(BSPAccelerator)
	acc.n = a.Size()
	acc.max_depth = int(8 + 1.3*math.Log(float32(acc.n)) + 0.5)
	acc.root = new(BSPNode)
	acc.root.box = *a.BoundingBox()
	acc.buildTree(acc.root, a.GetIntersectables(), X, 0)
	return acc
}

func (acc *BSPAccelerator) buildTree(node *BSPNode, inters []util.Intersectable, splitAxis Axis, depth int) *BSPNode {
	if depth > acc.max_depth || len(inters) < MIN_NR_PRIMITIVES {
		node.inters = inters
		return node
	}
	leftInters := make([]util.Intersectable, 0, len(inters)/2)
	rightInters := make([]util.Intersectable, 0, len(inters)/2)
	// Save split axis in node and prepare bounding boxes.
	node.splitAxis = splitAxis
	b := node.box
	node.splitPos = b.Center()[splitAxis]
	leftBoxMax := b.Max
	leftBoxMax[splitAxis] = node.splitPos
	rightBoxMin := b.Min
	rightBoxMin[splitAxis] = node.splitPos

	var leftNode, rightNode BSPNode
	leftNode.box = vec3.Box{b.Min, leftBoxMax}
	rightNode.box = vec3.Box{rightBoxMin, b.Max}

	for i := range inters {
		if leftNode.box.Intersects(inters[i].BoundingBox()) {
			leftInters = append(leftInters, inters[i])
		}
		if rightNode.box.Intersects(inters[i].BoundingBox()) {
			rightInters = append(rightInters, inters[i])
		}
	}
	nextSplitAxis := Axis((int(node.splitAxis) + 1) % 3)
	node.left = acc.buildTree(&leftNode, leftInters, nextSplitAxis, depth+1)
	node.right = acc.buildTree(&rightNode, rightInters, nextSplitAxis, depth+1)
	return node
}

// A helper struct for intersect method.
type BSPStackNode struct {
	node       *BSPNode
	tmin, tmax float32
}

// A wrapper around ray that precomputes some stuff.
type BSPRay struct {
	*util.Ray
	invDirection vec3.T
}

func NewBSPRay(ray *util.Ray) *BSPRay {
	r := BSPRay{ray, vec3.T{1 / ray.Direction[0], 1 / ray.Direction[1], 1 / ray.Direction[2]}}
	return &r
}

func (acc *BSPAccelerator) Intersect(ray *util.Ray) *util.Hitrecord {
	r := NewBSPRay(ray)
	tmin, tmax, doesIntersect := doesRayIntersectBox(r, &acc.root.box)
	if !doesIntersect {
		return nil
	}
	nodeStack := make([]*BSPStackNode, 0, 5)

	var nearestHit *util.Hitrecord
	nearestT := float32(math.MaxFloat32)
	var node *BSPNode = acc.root
	for node != nil {
		if nearestT < tmin {
			break
		}
		if !node.isLeaf() {
			tSplitAxis := (node.splitPos - r.Origin[node.splitAxis]) / r.Direction[node.splitAxis]
			var first, second *BSPNode

			if r.Origin[node.splitAxis] < node.splitPos {
				first = node.left
				second = node.right
			} else {
				first = node.right
				second = node.left
			}
			// process children
			if _, _, doesIntersect := doesRayIntersectBox(r, &first.box); tSplitAxis > tmax || tSplitAxis < 0 || (math.Abs(tSplitAxis) < 1e-5 && doesIntersect) {
				node = first
			} else if _, _, doesIntersect := doesRayIntersectBox(r, &second.box); tSplitAxis < tmin || (math.Abs(tSplitAxis) < 1e-5 && doesIntersect) {
				node = second
			} else {
				node = first
				laterNode := BSPStackNode{second, tSplitAxis, tmax}
				nodeStack = append(nodeStack, &laterNode)
				tmax = tSplitAxis
			}
		} else {
			for i := range node.inters {
				hit := node.inters[i].Intersect(r.Ray)
				if hit != nil && hit.T < nearestT && hit.T > 0 {
					nearestT = hit.T
					nearestHit = hit
				}
			}
			if len(nodeStack) > 0 {
				s := nodeStack[len(nodeStack)-1]
				nodeStack = nodeStack[:len(nodeStack)-1]
				node = s.node
				tmin = s.tmin
				tmax = s.tmax
			} else {
				break
			}
		}
	}
	return nearestHit
}

func (acc *BSPAccelerator) IntersectSlow(ray *util.Ray) *util.Hitrecord {
	r := NewBSPRay(ray)
	nodeStack := make([]*BSPNode, 1)
	nodeStack[0] = acc.root
	var nearestHit *util.Hitrecord
	nearestT := float32(math.MaxFloat32)
	var n *BSPNode
	for len(nodeStack) > 0 {
		n, nodeStack = nodeStack[len(nodeStack)-1], nodeStack[:len(nodeStack)-1]
		if n.inters != nil {
			for i := range n.inters {
				currentHit := n.inters[i].Intersect(r.Ray)
				if currentHit != nil && nearestT > currentHit.T && currentHit.T > 0 {
					nearestT = currentHit.T
					nearestHit = currentHit
				}
			}
		} else {
			if _, _, doesIntersect := doesRayIntersectBox(r, &n.left.box); doesIntersect {
				nodeStack = append(nodeStack, n.left)
			}
			if _, _, doesIntersect := doesRayIntersectBox(r, &n.right.box); doesIntersect {
				nodeStack = append(nodeStack, n.right)
			}
		}
	}
	return nearestHit
}

// For description of algorithm, see
// http://www.scratchapixel.com/lessons/3d-basic-lessons/lesson-7-intersecting-simple-shapes/ray-box-intersection/
func doesRayIntersectBox(r *BSPRay, b *vec3.Box) (tmin, tmax float32, doesIntersect bool) {
	bounds := [2]*vec3.T{&b.Min, &b.Max}
	signs := [3]uint32{0, 0, 0}
	for i := 0; i < 3; i++ {
		if r.Direction[i] < 0 {
			signs[i] = 1
		}
		// If statement is unexplicably faster than this:
		//signs[i] = *(*uint32)(unsafe.Pointer(&r.Direction[i])) >> 31
	}
	// Intersections on X planes.
	tmin = (bounds[signs[X]][X] - r.Origin[X]) * r.invDirection[X]
	tmax = (bounds[1-signs[X]][X] - r.Origin[X]) * r.invDirection[X]
	// Intersections on Y planes.
	tymin := (bounds[signs[Y]][Y] - r.Origin[Y]) * r.invDirection[Y]
	tymax := (bounds[1-signs[Y]][Y] - r.Origin[Y]) * r.invDirection[Y]
	if tmin > tymax || tymin > tmax {
		return 0, 0, false
	}
	if tymin > tmin {
		tmin = tymin
	}
	if tymax < tmax {
		tmax = tymax
	}
	// Intersections on Z planes.
	tzmin := (bounds[signs[Z]][Z] - r.Origin[Z]) * r.invDirection[Z]
	tzmax := (bounds[1-signs[Z]][Z] - r.Origin[Z]) * r.invDirection[Z]
	if tmin > tzmax || tzmin > tmax {
		return 0, 0, false
	}
	if tzmin > tmin {
		tmin = tzmin
	}
	if tzmax < tmax {
		tmax = tzmax
	}
	return tmin, tmax, true
}

func (acc *BSPAccelerator) BoundingBox() *vec3.Box {
	return &acc.box
}
