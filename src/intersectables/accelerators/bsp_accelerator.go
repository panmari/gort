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
	left, right *BSPNode
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
	acc.root.splitAxis = X
	acc.buildTree(acc.root, a.GetIntersectables(), 0)
	return acc
}

func (acc *BSPAccelerator) buildTree(node *BSPNode, inters []util.Intersectable, depth int) *BSPNode {
	if depth > acc.max_depth || len(inters) < MIN_NR_PRIMITIVES {
		node.inters = inters
		return node
	}
	leftInters := make([]util.Intersectable, 0, len(inters)/2)
	rightInters := make([]util.Intersectable, 0, len(inters)/2)
	b := node.box
	splitDist := b.Center()[node.splitAxis]
	leftBoxMax := b.Max
	leftBoxMax[node.splitAxis] = splitDist
	rightBoxMin := b.Min
	rightBoxMin[node.splitAxis] = splitDist

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
	node.left = acc.buildTree(&leftNode, leftInters, depth+1)
	node.right = acc.buildTree(&rightNode, rightInters, depth+1)
	return node
}

func (acc *BSPAccelerator) Intersect(r *util.Ray) *util.Hitrecord {
	nodeStack := make([]*BSPNode, 1)
	nodeStack[0] = acc.root
	var nearestHit *util.Hitrecord
	nearestT := float32(math.MaxFloat32)
	var n *BSPNode
	for len(nodeStack) > 0 {
		n, nodeStack = nodeStack[len(nodeStack)-1], nodeStack[:len(nodeStack)-1]
		if n.inters != nil {
			for i := range n.inters {
				currentHit := n.inters[i].Intersect(r)
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
func doesRayIntersectBox(r *util.Ray, b *vec3.Box) (tmin, tmax float32, doesIntersect bool) {
	bounds := [2]*vec3.T{&b.Min, &b.Max}
	invdir := vec3.T{1 / r.Direction[0], 1 / r.Direction[1], 1 / r.Direction[2]}
	signs := [3]int{0, 0, 0}
	for i := 0; i < 3; i++ {
		if invdir[i] < 0 {
			signs[i] = 1
		}
	}
	tmin = (bounds[signs[0]][0] - r.Origin[0]) * invdir[0]
	tmax = (bounds[1-signs[0]][0] - r.Origin[0]) * invdir[0]
	tymin := (bounds[signs[1]][1] - r.Origin[1]) * invdir[1]
	tymax := (bounds[1-signs[1]][1] - r.Origin[1]) * invdir[1]
	if tmin > tymax || tymin > tmax {
		return 0, 0, false
	}
	if tymin > tmin {
		tmin = tymin
	}
	if tymax < tmax {
		tmax = tymax
	}
	tzmin := (bounds[signs[2]][2] - r.Origin[2]) * invdir[2]
	tzmax := (bounds[1-signs[2]][2] - r.Origin[2]) * invdir[2]
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
