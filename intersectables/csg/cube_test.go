package csg

import (
	"testing"

	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

func TestCubeIntersection(t *testing.T) {
	c := NewDiffuseCube()

	testCases := []struct {
		name    string
		ray     util.Ray
		wantHit bool
	}{
		{
			name:    "ray from left",
			ray:     util.Ray{vec3.T{-3, 0, 0}, vec3.T{1, 0, 0}},
			wantHit: true,
		},
		{
			name:    "ray from right",
			ray:     util.Ray{vec3.T{3, 0, 0}, vec3.T{-1, 0, 0}},
			wantHit: true,
		},
		{
			name:    "ray from above",
			ray:     util.Ray{vec3.T{0, 5, 0}, vec3.T{0, -1, 0}},
			wantHit: true,
		},
		{
			// Hits corner, which might have weird normal.
			name:    "ray diagonal",
			ray:     util.Ray{vec3.T{5, 5, 5}, vec3.T{-1, -1, -1}},
			wantHit: true,
		},
		{
			name:    "ray from within",
			ray:     util.Ray{vec3.Zero, vec3.T{-1, 0, 0}},
			wantHit: true,
		},
	}

	for _, tc := range testCases {
		got := c.Intersect(&tc.ray)
		// t.Log(tc.name, got)
		if gotHit := got != nil; gotHit != tc.wantHit {
			t.Errorf("s.Intersect(%q), got %v, want %v", tc.name, got, tc.wantHit)
		}
		// TODO(panmari): Also check attributes of hitrecord.
	}
}
