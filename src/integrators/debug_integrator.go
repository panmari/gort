package integrators

import (
	"util"
	"github.com/ungerik/go3d/vec3"
)

type DebugIntegrator struct {
}

func (d *DebugIntegrator) Integrate(r *util.Ray, root util.Intersectable) *vec3.T {
	if _, doesHit := root.Intersect(r); doesHit {
		return &vec3.T{1,0,0}
	}
	return &vec3.T{0,0,0}
}

func MakeDebugIntegrator() (*DebugIntegrator) {
	return new(DebugIntegrator)
}