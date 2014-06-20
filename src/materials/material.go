package materials

import (
	"util"
	"github.com/ungerik/go3d/vec3"
)

type Material interface {
	EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) (*vec3.T)
	GetEmissionSample(hit *util.Hitrecord, sample [2]float32) (*vec3.T)
	GetShadingSample(hit *util.Hitrecord, sample [2]float32) (*vec3.T)
	DoesCastShadows() bool
	HasSpecularReflection() bool
	HasSpecularRefraction() bool
	
	// these pass a copy back
	EvaluateSpecularReflection(hit *util.Hitrecord) (vec3.T)
	EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) (vec3.T)
}