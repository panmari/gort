package materials

import (
	"testing"

	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

func TestGridMaterial(t *testing.T) {
	g := NewGrid(NewDiffuseMaterial(0.2, 0.2, 0.2), DiffuseDefault, 0.9, 0.1)

	testCases := []struct {
		hit  util.Hitrecord
		want vec3.T
	}{
		{
			hit:  util.Hitrecord{Position: vec3.Zero},
			want: vec3.T{0.06366198, 0.06366198, 0.06366198},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{0, 0.5, 0.5}},
			want: vec3.T{0.06366198, 0.06366198, 0.06366198},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{0.95, 0.95, 0.95}},
			want: vec3.T{0.31830987, 0.31830987, 0.31830987},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{0, 0.95, 0.95}},
			want: vec3.T{0.31830987, 0.31830987, 0.31830987},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{3, 3.95, 3.95}},
			want: vec3.T{0.31830987, 0.31830987, 0.31830987},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{-1.5, -1.5, 0}},
			want: vec3.T{0.06366198, 0.06366198, 0.06366198},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{-3, -3.05, -3.05}},
			want: vec3.T{0.31830987, 0.31830987, 0.31830987},
		},
		{
			hit:  util.Hitrecord{Position: vec3.T{0, -0.05, -0.05}},
			want: vec3.T{0.31830987, 0.31830987, 0.31830987},
		},
	}
	for _, tc := range testCases {
		got := g.EvaluateBRDF(&tc.hit, nil, nil)
		if got != tc.want {
			t.Errorf("g.EvaluateBRDF(%v), got %v, want %v", tc.hit, got, tc.want)
		}
	}
}
