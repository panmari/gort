package test

import (
	"github.com/ungerik/go3d/vec3"
	"intersectables"
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
	if !HasTinyDifference(t0, -2) {
		t.Fail()
	}
	if !HasTinyDifference(t1, 2) {
		t.Fail()
	}

}

func TestSphereIntersection(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitX}
	s := intersectables.Sphere{vec3.Zero, 2}
	hit := s.Intersect(&r)
	expected := vec3.T{2, 0, 0}
	if !hit.DoesHit() {
		t.Errorf("Did not hit sphere!")
	}
	if hit.Point != expected {
		t.Errorf("Did hit something unexpected: %v", hit.Point)
	}
}

