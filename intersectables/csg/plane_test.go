package csg

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

func TestCSGPlaneIntersection(t *testing.T) {
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
		}
		// TODO(panmari): Also check attributes of hitrecord.
	}
}

func TestCSGPlaneIntersectionInverse(t *testing.T) {
	s := NewPlane(vec3.T{-1, 0, 0}, 2, materials.DiffuseDefault)

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
			ray:     util.Ray{vec3.Zero, vec3.T{-1, 0, 0}},
			wantHit: false,
		},
		{
			name:    "pointing towards plane ray",
			ray:     util.Ray{vec3.Zero, vec3.T{1, 0, 0}},
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

func TestCSGPlaneGetIntervalboundaries(t *testing.T) {
	s := NewPlane(vec3.UnitX, 0, materials.DiffuseDefault)

	testCases := []struct {
		name string
		ray  util.Ray
		want ByT
	}{
		{
			name: "orthogonal ray starting on edge",
			ray:  util.Ray{Origin: vec3.Zero, Direction: vec3.T{1, 0, 0}},
			want: ByT{
				{t: 0},
				{t: float32(math.Inf(-1))},
			},
		},
		{
			name: "orthogonal ray starting inside",
			ray:  util.Ray{Origin: vec3.T{-1, 0, 0}, Direction: vec3.T{1, 0, 0}},
			want: ByT{
				{t: 1},
				{t: float32(math.Inf(-1))},
			},
		},
		{
			name: "orthogonal ray starting outside",
			ray:  util.Ray{Origin: vec3.T{1, 0, 0}, Direction: vec3.T{1, 0, 0}},
			want: ByT{
				{t: -1},
				{t: float32(math.Inf(-1))},
			},
		},
		{
			name: "orthogonal inverse ray starting on edge",
			ray:  util.Ray{Origin: vec3.Zero, Direction: vec3.T{-1, 0, 0}},
			want: ByT{
				{t: 0},
				{t: float32(math.Inf(1))},
			},
		},
		// TODO(panmari): Needs fixing.
		// {
		// 	name: "parallel ray inside",
		// 	ray:  util.Ray{Origin: vec3.Zero, Direction: vec3.T{0, 1, 0}},
		// 	want: ByT{
		// 		{t: float32(math.Inf(-1))},
		// 		{t: float32(math.Inf(1))},
		// 	},
		// },
		{
			name: "parallel ray outside",
			ray:  util.Ray{Origin: vec3.T{1, 0, 0}, Direction: vec3.T{0, 1, 0}},
			want: ByT{},
		},
	}

	trans := cmp.Transformer("ExtractTs", func(in ByT) []float32 {
		out := make([]float32, 0, len(in))
		for _, ib := range in {
			out = append(out, ib.t)
		}
		return out
	})
	for _, tc := range testCases {
		got := s.GetIntervalBoundaries(&tc.ray)

		if diff := cmp.Diff(got, tc.want, trans); diff != "" {
			t.Errorf("s.GetIntervalBoundaries(%q), got %v, want %v, diff %s", tc.name, got, tc.want, diff)
		}
	}
}

func BenchmarkCSGPlaneIntersection(b *testing.B) {
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
