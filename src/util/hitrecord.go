package util

import (
	"github.com/ungerik/go3d/vec3"
)

type Hitrecord struct {
	T             float32
	Position      vec3.T
	Normal        vec3.T
	W_in          vec3.T
	U             float32
	V             float32
	Intersectable Intersectable
	Material      Material
}

type Intersectable interface {
	Intersect(r *Ray) (*Hitrecord, bool)
}

type Material interface {
	EvaluateEmission(hit *Hitrecord, wOut *vec3.T) vec3.T
	GetEmissionSample(hit *Hitrecord, sample [2]float32) *vec3.T
	GetShadingSample(hit *Hitrecord, sample [2]float32) *vec3.T
	DoesCastShadows() bool
	HasSpecularReflection() bool
	HasSpecularRefraction() bool

	// these pass a copy back
	EvaluateSpecularReflection(hit *Hitrecord) vec3.T
	EvaluateBRDF(hit *Hitrecord, wOut, wIn *vec3.T) vec3.T
}
