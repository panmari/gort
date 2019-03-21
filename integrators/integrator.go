package integrators

import (
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type Integrator interface {
	Integrate(r *util.Ray, depth int) *vec3.T
}
