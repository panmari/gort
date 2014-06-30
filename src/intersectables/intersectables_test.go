package intersectables

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"util"
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
	s := MakeDiffuseSphere(vec3.Zero, 2)
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
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := MakeDiffusePlane(vec3.UnitX, 1)

	if hit := s.Intersect(&parallelRay); hit != nil {
		t.Errorf("Parallel ray hit plane!", hit)
	}

	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	if hit := s.Intersect(&pointingAwayRay); hit != nil {
		t.Errorf("Ray pointing away from plane hit plane!", hit)
	}

	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	hit := s.Intersect(&pointingTowardsRay)
	if hit == nil {
		t.Errorf("Ray pointing towards plane does not hit plane!", hit)
	}
	if hit.T != 1 {
		t.Errorf("Ray pointing towards plane hits at strange T: %f", hit.T)
	}
}

func BenchmarkPlaneIntersection(b *testing.B) {
	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := MakeDiffusePlane(vec3.UnitX, 1)
	
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
	s := MakeDiffuseSphere(vec3.UnitX, 1)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(&pointingTowardsRay)
		s.Intersect(&pointingAwayRay)
		s.Intersect(&parallelRay)
	}
}