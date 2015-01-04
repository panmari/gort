package accelerators

import (
	"intersectables"
	"math"
	"github.com/ungerik/go3d/vec3"
	"util"
)

type BSPAccelerator struct {
	n, max_depth int
	root *BSPNode
	box vec3.Box
}

type BSPNode struct {
	inters []util.Intersectable
	box vec3.Box
	splitAxis Axis
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

func NewBSPAccelerator(a intersectables.Aggregate) {
	acc := new(BSPAccelerator)
	acc.n = a.Size()
	acc.max_depth = int(8 + 1.3 * math.Log(float64(acc.n)) + 0.5)
	acc.root = new(BSPNode)
	acc.root.box = *a.BoundingBox()
	acc.root.splitAxis = X
	acc.buildTree(acc.root, a.GetIntersectables(), 0)
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
	node.left = acc.buildTree(&leftNode, leftInters, depth + 1)
	node.right = acc.buildTree(&rightNode, rightInters, depth + 1)
	return node
}

func (acc *BSPAccelerator) BoundingBox() *vec3.Box {
	return &acc.box;
}
