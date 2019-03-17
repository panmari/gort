package intersectables

import (
	"testing"

	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

func TestRayScaling(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitXYZ}
	extrapolated := r.PointAt(5.0)
	expected := vec3.T{5, 5, 5}
	if expected != extrapolated {
		t.Fail()
	}
}

func TestSolvingQuadraticFunction(t *testing.T) {
	t0, t1, hasSolution := util.SolveQuadratic(1, 0, -4)
	if !hasSolution {
		t.Fail()
	}
	if t0 != -2 {
		t.Fail()
	}
	if t1 != 2 {
		t.Fail()
	}

}

func TestSphereIntersection(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewSphere(vec3.Zero, 2, materials.DiffuseDefault)
	hit := s.Intersect(&r)
	expected := vec3.T{2, 0, 0}
	if hit == nil {
		t.Errorf("Did not hit sphere!")
	}
	if hit.Position != expected {
		t.Errorf("Did hit something unexpected: %v", hit.Position)
	}
}

func TestPlaneIntersection(t *testing.T) {
	s := NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)

	testCases := []struct {
		name    string
		ray     util.Ray
		wantHit bool
	}{
		{
			name:    "parallel ray",
			ray:     util.Ray{vec3.Zero, vec3.UnitY},
			wantHit: false,
		},
		{
			name:    "pointing away from plane ray",
			ray:     util.Ray{vec3.Zero, vec3.UnitX},
			wantHit: false,
		},
		{
			name:    "pointing towards plane ray",
			ray:     util.Ray{vec3.Zero, vec3.T{-1, 0, 0}},
			wantHit: true,
		},
		{
			name:    "ray from behind plane",
			ray:     util.Ray{vec3.T{-2, 0, 0}, vec3.T{1, 0, 0}},
			wantHit: true,
		},
	}
	for _, tc := range testCases {
		got := s.Intersect(&tc.ray)
		if gotHit := got != nil; gotHit != tc.wantHit {
			t.Errorf("s.Intersect(%q), got %v, want %v", tc.name, got, tc.wantHit)
			t.Logf("Hitrecord: %v", got)
		}
		// TODO(panmari): Also check attributes of hitrecord.
	}
}

func BenchmarkPlaneIntersection(b *testing.B) {
	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(&pointingTowardsRay)
		s.Intersect(&pointingAwayRay)
		s.Intersect(&parallelRay)
	}
}

func BenchmarkSphereIntersection(b *testing.B) {
	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewSphere(vec3.UnitX, 1, materials.DiffuseDefault)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(&pointingTowardsRay)
		s.Intersect(&pointingAwayRay)
		s.Intersect(&parallelRay)
	}
}
