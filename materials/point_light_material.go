package materials

import (
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type PointLightMaterial struct {
	Emission vec3.T
}

func (m *PointLightMaterial) EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) vec3.T {
	return m.Emission
}

func (m *PointLightMaterial) GetEmissionSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *PointLightMaterial) GetShadingSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *PointLightMaterial) DoesCastShadows() bool {
	return false
}

func (m *PointLightMaterial) HasSpecularReflection() bool {
	return false
}

func (m *PointLightMaterial) HasSpecularRefraction() bool {
	return false
}

func (m *PointLightMaterial) EvaluateSpecularReflection(hit *util.Hitrecord) (util.ShadingSample, bool) {
	return util.ShadingSample{}, false
}

func (m *PointLightMaterial) EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) vec3.T {
	return vec3.T{}
}

func MakePointLightMaterial(emission vec3.T) *PointLightMaterial {
	m := new(PointLightMaterial)
	m.Emission = emission
	return m
}
