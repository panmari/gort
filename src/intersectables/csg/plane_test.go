package csg

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"util"
)


func TestPlaneIntersection(t *testing.T) {
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewDiffusePlane(vec3.UnitX, 1)

	if hit, doesHit := s.Intersect(&parallelRay); doesHit {
		t.Errorf("Parallel ray hit plane!", hit)
	}

	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	if hit, doesHit := s.Intersect(&pointingAwayRay); doesHit {
		t.Errorf("Ray pointing away from plane hit plane!", hit)
	}

	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	hit, doesHit := s.Intersect(&pointingTowardsRay)
	if !doesHit {
		t.Errorf("Ray pointing towards plane does not hit plane!", hit)
	}
	if hit.T != 1 {
		t.Errorf("Ray pointing towards plane hits at strange T: %f", hit.T)
	}
}