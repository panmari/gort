package csg

import (
	"fmt"
	"github.com/ungerik/go3d/vec3"
	"sort"
	"util"
)

type Node struct {
	left, right Shape
	box         vec3.Box
	operation   Operation
}

type Operation int

const (
	INTERSECT Operation = iota
	ADD
	SUBTRACT
)

func (n *Node) GetIntervalBoundaries(r *util.Ray) ByT {
	leftIntervals := n.left.GetIntervalBoundaries(r)
	rightIntervals := n.right.GetIntervalBoundaries(r)
	combined := combineIntervals(leftIntervals, rightIntervals)
	sort.Sort(combined)
	inLeft, inRight := false, false
	previousWasStart := false

	cleanCount := 0
	for i := range combined {
		b := &combined[i]
		if b.belongsToLeft {
			inLeft = b.isStart
		} else {
			inRight = b.isStart
		}
		switch n.operation {
		case INTERSECT:
			b.isStart = inLeft && inRight
		case ADD:
			b.isStart = inLeft || inRight
		case SUBTRACT:
			b.isStart = inLeft && !inRight
			// In a subtract operation, the subtracted solid is turned
			// inside out,
			// or it "switches sign", so we need to flip its normal
			// direction
			if !b.belongsToLeft && b.hit != nil {
				b.hit.Normal.Scale(-1)
			}
		}
		// remove start - start or end - end combinations by only adding good ones back
		if previousWasStart != b.isStart {
			combined[cleanCount] = *b
			cleanCount++
		}
		previousWasStart = b.isStart
	}
	// only return 0..cleanCount entries, the others are garbage
	return combined[:cleanCount]
}

func (i *Node) BoundingBox() *vec3.Box {
	return &i.box
}

func (i *Node) String() string {
	return fmt.Sprintf("left: %v, right: %v", i.left, i.right)
}

//combines the two intervals and tags all containing intervalboundaries as left resp. right
func combineIntervals(left, right ByT) ByT {
	for i := range left {
		left[i].belongsToLeft = true
	}
	for i := range right {
		right[i].belongsToLeft = false
	}
	combined := append(left, right...)
	return combined
}

func NewNode(left, right *Solid, o Operation) *Solid {
	n := new(Node)
	n.left = left
	n.right = right
	n.operation = o
	// TODO: Do correct thing depending on operation, not extend for every operation.
	n.box = *left.BoundingBox()
	n.box.ExtendBox(right.BoundingBox())
	return &Solid{n}
}
