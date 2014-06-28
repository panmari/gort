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
