package csg

import (
	"util"
	"sort"
	"fmt"
)

type Node struct {
	left, right 	Shape
	operation       Operation
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
	for _, b := range combined {
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
			combined[cleanCount] = b
			cleanCount++
		}
		previousWasStart = b.isStart
	}
	// only return 0..cleanCount entries, the others are garbage
	return combined[:cleanCount];
}

func (i *Node) String() string {
	return fmt.Sprintf("left: %v, right: %v", i.left, i.right)
}

//combines the two intervals and tags all containing intervalboundaries as left resp. right
func combineIntervals(left, right ByT) ByT {
	for _, ib := range left {
		ib.belongsToLeft = true
	}
	for _, ib := range right {
		ib.belongsToLeft = false
	}
	combined := append(left, right...)
	return combined
}

func NewNode(left, right *Solid, o Operation) *Solid {
	n := &Node{left, right, o}
	return &Solid{n}
}