package csg

import (
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
	"math"
	"testing"
)

func TestPlaneIntersection(t *testing.T) {
	s := NewDiffusePlane(vec3.UnitX, 1)

	parallelRay := util.Ray{vec3.Zero, vec3.UnitY}
	if hit := s.Intersect(&parallelRay); hit != nil {
		t.Errorf("s.Intersect(parallelRay): want nil, got %v", hit)
	}

	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	if hit := s.Intersect(&pointingAwayRay); hit != nil {
		t.Errorf("s.Intersect(pointingAwayRay): want nil, got %v", hit)
	}

	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	hit := s.Intersect(&pointingTowardsRay)
	if hit == nil {
		t.Errorf("s.Intersect(pointingTowardsRay): want not nil, got %v", hit)
	}
	if hit.T != 1 {
		t.Errorf("Ray pointing towards plane hits at strange T: %f", hit.T)
	}
}

func TestPlaneGetIntervalboundaries(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	s := Plane{*intersectables.MakeDiffusePlane(vec3.UnitX, 1)}
	ibs := s.GetIntervalBoundaries(&r)
	if ibs[0].t != 1 {
		t.Errorf("First intersection not correct: %f", ibs[0].t)
	}
	if ibs[1].t != float32(math.Inf(1)) {
		t.Errorf("Second intersection not correct: %f", ibs[1].t)
	}

}

func BenchmarkPlaneIntersection(b *testing.B) {
	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewDiffusePlane(vec3.UnitX, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(&pointingTowardsRay)
		s.Intersect(&pointingAwayRay)
		s.Intersect(&parallelRay)
	}
}
