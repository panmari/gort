package csg

import (
	"github.com/ungerik/go3d/vec3"
	"intersectables"
	"testing"
	"util"
)

func TestSphereIntersection(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewDiffuseSphere(vec3.Zero, 2)
	hit := s.Intersect(&r)
	expected := vec3.T{2, 0, 0}
	if hit == nil {
		t.Errorf("Did not hit sphere!")
	}
	if hit.Position != expected {
		t.Errorf("Did hit something unexpected: %v", hit.Position)
	}
}

func TestSphereGetIntervalboundaries(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitX}
	s := Sphere{*intersectables.MakeDiffuseSphere(vec3.Zero, 2)}
	ibs := s.GetIntervalBoundaries(&r)
	if ibs[0].t != -2 {
		t.Errorf("First intersection not correct: %f", ibs[0].t)
	}
	if ibs[1].t != 2 {
		t.Errorf("Second intersection not correct: %f", ibs[1].t)
	}
}

func BenchmarkSphereIntersection(b *testing.B) {
	pointingTowardsRay := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}}
	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	parallelRay := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewDiffuseSphere(vec3.UnitX, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(&pointingTowardsRay)
		s.Intersect(&pointingAwayRay)
		s.Intersect(&parallelRay)
	}
}
