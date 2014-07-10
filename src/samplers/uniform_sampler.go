package samplers

import (
	"github.com/barnex/fmath"
)

type UniformSampler struct {
	
}

/*
* Expects n to be a square number.
*/
func (r *UniformSampler) Get2DSamples(n int) [][2]float32 {
	k := fmath.Ceil(fmath.Sqrt(float32(n)))
	k_int := int(k)
	dist := 1/k
	offset := 1/(k*2)
	samples := make([][2]float32, k_int*k_int)
	for i := range samples {
		samples[i][0] = offset + dist*float32(i % k_int)
		samples[i][1] = offset + dist*float32(i / k_int)
	}
	return samples
}

func MakeUniformSampler(seed int64) Sampler {
	return new(UniformSampler)
}
