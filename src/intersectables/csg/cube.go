package csg

import (
	"github.com/ungerik/go3d/vec3"
)

func NewDiffuseCube() *Solid {
	p1 := NewDiffusePlane(vec3.T{1,0,0},-1)
	p2 := NewDiffusePlane(vec3.T{-1,0,0},-1)
	p3 := NewDiffusePlane(vec3.T{0,1,0},-1)
	p4 := NewDiffusePlane(vec3.T{0,-1,0},-1)
	p5 := NewDiffusePlane(vec3.T{0,0,1},-1)
	p6 := NewDiffusePlane(vec3.T{0,0,-1},-1)
	
	n1 := NewNode(p1, p2, INTERSECT)
	n2 := NewNode(p3, p4, INTERSECT)
	n3 := NewNode(p5, p6, INTERSECT)
	n4 := NewNode(n1, n2, INTERSECT)
	root := NewNode(n3, n4, INTERSECT)
	return root
}
