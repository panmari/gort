package csg

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"util"
)


func TestSphereIntersection(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitX}
	s := NewDiffuseSphere(vec3.Zero, 2)
	hit, doesHit := s.Intersect(&r)
	expected := vec3.T{2, 0, 0}
	if !doesHit {
		t.Errorf("Did not hit sphere!")
	}
	if hit.Position != expected {
		t.Errorf("Did hit something unexpected: %v", hit.Position)
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