package integrators

import (
	"github.com/ungerik/go3d/vec3"
	"util"
)

type Integrator interface {
	Integrate(r *util.Ray) *vec3.T
}
