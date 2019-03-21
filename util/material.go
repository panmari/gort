package util

import "github.com/ungerik/go3d/vec3"

type Material interface {
	EvaluateEmission(hit *Hitrecord, wOut *vec3.T) vec3.T
	GetEmissionSample(hit *Hitrecord, sample [2]float32) *vec3.T
	GetShadingSample(hit *Hitrecord, sample [2]float32) *vec3.T
	DoesCastShadows() bool

	EvaluateSpecularReflection(hit *Hitrecord) (ShadingSample, bool)
	// EvaluateSpecularRefraction(hit *Hitrecord) (ShadingSample, bool)
	EvaluateBRDF(hit *Hitrecord, wOut, wIn *vec3.T) vec3.T
}
