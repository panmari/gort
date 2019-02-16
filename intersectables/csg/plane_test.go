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

	testCases := []struct {
		name string
		ray  util.Ray
		wantHit bool
	}{
		{
			name: "parallel ray",
			ray:  util.Ray{vec3.Zero, vec3.UnitY},
			wantHit: false,
		},
		{
			name: "pointing away from plane ray",
			ray:  util.Ray{vec3.Zero, vec3.UnitX},
			wantHit: false,
		},
		{
			name: "pointing towards plane ray",
			ray:  util.Ray{vec3.Zero, vec3.T{-1, 0, 0}},
			wantHit: true,
		},
	}
	for _, tc := range testCases {
	  got := s.Intersect(&tc.ray)
		if gotHit := got != nil; gotHit != tc.wantHit {
			t.Errorf("s.Intersect(%q), got %v, want %v", tc.name, got, tc.wantHit)
		}
		// TODO(panmari): Also check attributes of hitrecord.
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
