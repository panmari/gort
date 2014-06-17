package integrators

import (
	"util"
)
type Integrator interface {
	
	Integrate(r *util.Ray) util.Spectrum
}