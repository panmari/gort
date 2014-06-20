package integrators

import (
	"util"
	"github.com/ungerik/go3d/vec3"
)
type Integrator interface {
	Integrate(r *util.Ray) *vec3.T
}