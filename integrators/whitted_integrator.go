package integrators

import (
	"github.com/barnex/fmath"
	"github.com/panmari/gort/lights"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

const maxDepth = 10

type WhittedIntegrator struct {
	Root        util.Intersectable
	PointLights []lights.LightGeometry
}

func (d *WhittedIntegrator) Integrate(ray *util.Ray, depth int) *vec3.T {
	outgoing := vec3.T{}
	if hit := d.Root.Intersect(ray); hit != nil {
		if depth < maxDepth {
			reflected := vec3.Zero
			refracted := vec3.Zero
			s, hasReflection := hit.Material.EvaluateSpecularReflection(hit)
			if hasReflection {
				reflected = s.BRDF
				recursiveRay := util.MakeEpsilonRay(&hit.Position, &s.W)
				spec := d.Integrate(recursiveRay, depth+1)
				reflected.Mul(spec)
			}
			// TODO: Refracted part.
			hasRefraction := false
			if hasReflection || hasRefraction {
				return reflected.Add(&refracted)
			}
		}
		for i := range d.PointLights {
			lightHit := d.PointLights[i].Sample([2]float32{0, 0})
			lightDir := vec3.Sub(&lightHit.Position, &hit.Position)
			dist2 := lightDir.LengthSqr()
			dist := fmath.Sqrt(dist2)
			lightDir.Scale(1.0 / dist)

			shadowRay := util.MakeEpsilonRay(&hit.Position, &lightDir)
			if shadowHit := d.Root.Intersect(shadowRay); shadowHit != nil && shadowHit.T < dist {
				continue
			}

			brdfValue := hit.Material.EvaluateBRDF(hit, &hit.W_in, &lightDir)

			inverseLightDir := lightDir.Scaled(-1)
			lightValue := lightHit.Material.EvaluateEmission(lightHit, &inverseLightDir)
			brdfValue.Mul(&lightValue)
			ndotl := fmath.Max(vec3.Dot(&hit.Normal, &lightDir), 0)
			brdfValue.Scale(ndotl / dist2)
			outgoing.Add(&brdfValue)
		}
	}
	return &outgoing
}

func MakeWhittedIntegrator(root util.Intersectable, pointLights []lights.LightGeometry) *WhittedIntegrator {
	return &WhittedIntegrator{Root: root, PointLights: pointLights}
}
