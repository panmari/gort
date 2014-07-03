package csg

import (
	"testing"
	"util"
	"github.com/ungerik/go3d/vec3"
)

func TestDodecahedronIntersection(t *testing.T) {
	d := NewDodecahedron()
	
	r := util.Ray{vec3.T{0,0,5}, vec3.T{0,0,-1}}
	hit := d.Intersect(&r)
	if hit == nil {
		t.Errorf("Ray did not hit")
	}
}

