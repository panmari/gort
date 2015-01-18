package materials

import (
	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/vec3"
	"util"
)

var (
	DiffuseDefault = &Diffuse{vec3.T{1 / fmath.Pi, 1 / fmath.Pi, 1 / fmath.Pi}}
)

type Diffuse struct {
	Kd vec3.T
}

func (m *Diffuse) EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) vec3.T {
	return vec3.T{}
}
func (m *Diffuse) GetEmissionSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *Diffuse) GetShadingSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
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

func (m *Diffuse) EvaluateSpecularReflection(hit *util.Hitrecord) vec3.T {
	return vec3.T{}
}

func (m *Diffuse) EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) vec3.T {
	return m.Kd
}

func NewDiffuseMaterial(r, g, b float32) *Diffuse {
	diffuseReflectance := vec3.T{r, g, b}
	diffuseReflectance.Scale(1 / fmath.Pi)
	m := Diffuse{diffuseReflectance}
	return &m
}

func MakeDiffuseMaterial(diffuseReflectance vec3.T) *Diffuse {
	m := new(Diffuse)
	m.Kd = diffuseReflectance.Scaled(1 / fmath.Pi)
	return m
}
