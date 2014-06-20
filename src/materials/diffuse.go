package materials

import (
	"github.com/ungerik/go3d/vec3"
	"github.com/barnex/fmath"
	"util"
)

type Diffuse struct {
	kd *vec3.T	
}

func (m *Diffuse) EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) (*vec3.T) {
	return nil
}
func (m *Diffuse) GetEmissionSample(hit *util.Hitrecord, sample [2]float32) (*vec3.T) {
	return nil
}

func (m *Diffuse) GetShadingSample(hit *util.Hitrecord, sample [2]float32) (*vec3.T) {
	return nil
}

func (m *Diffuse) DoesCastShadows() bool {
	return true
}

func (m *Diffuse) HasSpecularReflection() bool {
	
	return false
}
func (m *Diffuse) HasSpecularRefraction() bool {
	return false
}

func (m *Diffuse) EvaluateSpecularReflection(hit *util.Hitrecord) (*vec3.T) {
	return new(vec3.T)
}

func (m *Diffuse) EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) (vec3.T) {
	return *m.kd
}

func MakeDiffuseMaterial(diffuseReflectance *vec3.T) *Diffuse {
	m := new(Diffuse)
	kd := diffuseReflectance.Scaled(1/fmath.Pi)
	m.kd = &kd 
	return m
}