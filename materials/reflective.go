package materials

import (
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type reflective struct {
	ks vec3.T
}

func (m *reflective) EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) vec3.T {
	return vec3.T{}
}
func (m *reflective) GetEmissionSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *reflective) GetShadingSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *reflective) DoesCastShadows() bool {
	return true
}

func (m *reflective) EvaluateSpecularReflection(hit *util.Hitrecord) (util.ShadingSample, bool) {
	r := util.Reflect(&hit.Normal, &hit.W_in)
	return util.ShadingSample{
		BRDF:        m.ks,
		W:           r,
		Probability: 1,
	}, true
}

func (m *reflective) EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) vec3.T {
	return m.ks
}

// NewReflective TODO(panmari): Documentation.
func NewReflective(ks vec3.T) util.Material {
	return &reflective{ks}
}
