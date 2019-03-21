package materials

import (
	"github.com/barnex/fmath"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type blinn struct {
	diffuse   vec3.T
	specular  vec3.T
	shyniness float32
}

func (m *blinn) EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) vec3.T {
	return vec3.Zero
}

func (m *blinn) GetEmissionSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *blinn) GetShadingSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *blinn) DoesCastShadows() bool {
	return true
}

func (m *blinn) EvaluateSpecularReflection(hit *util.Hitrecord) (util.ShadingSample, bool) {
	return util.ShadingSample{}, false
}

func (m *blinn) EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) vec3.T {
	halfwayVector := vec3.Add(wIn, wOut)
	halfwayVector.Normalize()

	total := m.diffuse.Scaled(vec3.Dot(wIn, &hit.Normal)) // Diffuse
	bling := fmath.Pow(vec3.Dot(&halfwayVector, &hit.Normal), m.shyniness)
	specular := m.specular.Scaled(bling)
	total.Add(&specular)  // Specular
	total.Add(&m.diffuse) // Ambient

	return total
}

func NewBlinn(diffuse, specular vec3.T, shyniness float32) util.Material {
	return &blinn{
		diffuse:   diffuse,
		specular:  specular,
		shyniness: shyniness,
	}
}
