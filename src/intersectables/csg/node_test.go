package csg

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"util"
	)

func TestSimpleNodeIntersection(t *testing.T) {
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	
	p := NewDiffusePlane(vec3.UnitX, 1)
	s := NewDiffuseSphere(vec3.Zero, 2)
	
	n := NewNode(p, s, ADD)
	
	if hit, doesHit := n.Intersect(&parallelRay); doesHit {
		t.Errorf("Parallel ray hit plane: %v", hit)
	}

	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	if hit, doesHit := n.Intersect(&pointingAwayRay); doesHit {
		t.Errorf("Ray pointing away from plane hit plane: %v", hit)
	}

	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	
	hit, doesHit := n.Intersect(&pointingTowardsRay)
	if !doesHit {
		t.Errorf("Ray pointing towards plane does not hit plane: %v", hit)
	}
	if hit.T != 1 {
		t.Errorf("Ray pointing towards plane hits at strange T: %f", hit.T)
	}
}