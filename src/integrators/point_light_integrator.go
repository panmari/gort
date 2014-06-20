package integrators

import (
	"util"
	"github.com/ungerik/go3d/vec3"
	"lights"
)

type PointLightIntegrator struct {
	lights []lights.LightGeometry
}

func (d *PointLightIntegrator) Integrate(r *util.Ray, root util.Intersectable) *vec3.T {
	if hit, doesHit := root.Intersect(r); doesHit {
		return &hit.Normal
	}
	return &vec3.T{0,0,0}
}

func MakePointLightIntegrator(pointLights []lights.LightGeometry) (*PointLightIntegrator) {
	return new(PointLightIntegrator)
}