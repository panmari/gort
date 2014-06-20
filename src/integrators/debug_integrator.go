package integrators

import (
	"util"
	"github.com/ungerik/go3d/vec3"
)

type DebugIntegrator struct {
	root util.Intersectable
}

func (d *DebugIntegrator) Integrate(r *util.Ray) *vec3.T {
	if _, doesHit := d.root.Intersect(r); doesHit {
		return &vec3.T{1,0,0}
	}
	return &vec3.T{0,0,0}
}

func MakeDebugIntegrator(root util.Intersectable) (*DebugIntegrator) {
	integrator := new(DebugIntegrator)
	integrator.root = root
	return integrator
}