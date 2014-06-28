package csg

import (
	"util"
	"sort"
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

func (n *Node) GetIntervalBoundaries(r *util.Ray) *ByT {
		leftIntervals := n.left.GetIntervalBoundaries(r)
		rightIntervals := n.right.GetIntervalBoundaries(r)
		
		combined := combineIntervals(leftIntervals, rightIntervals)
		sort.Sort(combined)
		
		inLeft, inRight := false, false
		for _, b := range *combined {
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
		}
		// clean up
		previousWasStart := false
		for _, b := range *combined {
			if previousWasStart == b.isStart {
				//remove
			}
			previousWasStart = b.isStart
		}
		return combined;
}

func combineIntervals(left, right *ByT) *ByT {
	for _, ib := range *left {
		ib.belongsToLeft = true
	}
	for _, ib := range *right {
		ib.belongsToLeft = false
	}	
	combined := append(*left, *right...)
	return &combined
}

func NewNode(left, right *Solid, o Operation) {
	
}