package integrators

import (
	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/vec3"
	"lights"
	"util"
)

type PointLightIntegrator struct {
	root        util.Intersectable
	pointLights []lights.LightGeometry
}

func (d *PointLightIntegrator) Integrate(r *util.Ray) *vec3.T {
	outgoing := vec3.T{}
	if hit := d.root.Intersect(r); hit != nil {
		for _, light := range d.pointLights {
			lightHit := light.Sample([2]float32{0, 0})
			lightDir := vec3.Sub(&lightHit.Position, &hit.Position)
			dist2 := lightDir.LengthSqr()
			lightDir.Normalize()

			dist := fmath.Sqrt(dist2)
			shadowRay := util.MakeEpsilonRay(&hit.Position, &lightDir)
			defer util.Raypool.Put(shadowRay)
			if shadowHit := d.root.Intersect(shadowRay); shadowHit != nil && shadowHit.T < dist {
				continue
			}

			brdfValue := hit.Material.EvaluateBRDF(hit, &hit.W_in, &lightDir)

			inverseLightDir := lightDir.Scaled(-1)
			lightValue := lightHit.Material.EvaluateEmission(lightHit, &inverseLightDir)
			brdfValue.Mul(&lightValue)
			ndotl := fmath.Max(vec3.Dot(&hit.Normal, &lightDir), 0)
			brdfValue.Scale(ndotl)
			brdfValue.Scale(1 / dist2)
			outgoing.Add(&brdfValue)
		}
	}
	return &outgoing
}

func MakePointLightIntegrator(root util.Intersectable, pointLights []lights.LightGeometry) *PointLightIntegrator {
	integrator := new(PointLightIntegrator)
	integrator.root = root
	integrator.pointLights = pointLights
	return integrator
}
