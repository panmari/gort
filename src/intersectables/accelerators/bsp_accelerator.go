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
}

type BSPNode struct {
	
}

func newBSPNode(box vec3.Box, axis Axis) *BSPNode {
	return new(BSPNode)
}

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
	acc.root = newBSPNode(a.GetBoundingBox(), Axis.X)
	buildTree(acc.root, a.GetIntersectables(), 0)
}

func buildTree(node *BSPNode, inters []util.Intersectable, depth int) *BSPNode {
	return new(BSPNode)
}