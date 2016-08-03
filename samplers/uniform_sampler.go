package samplers

import (
	"github.com/barnex/fmath"
)

type UniformSampler struct {
	Samples [][2]float32
}

/*
* Expects n to be a square number.
 */
func (r *UniformSampler) Get2DSamples(n int) [][2]float32 {
	k := fmath.Ceil(fmath.Sqrt(float32(n)))
	k_int := int(k)
	dist := 1 / k
	offset := 1 / (k * 2)
	r.Samples = r.Samples[0 : k_int*k_int] // adapt size of slice
	for i := range r.Samples {
		r.Samples[i][0] = offset + dist*float32(i%k_int)
		r.Samples[i][1] = offset + dist*float32(i/k_int)
	}
	return r.Samples
}

func (r *UniformSampler) DuplicateAndSeed(seed int64) Sampler {
	return NewUniformSampler(len(r.Samples))
}

func NewUniformSampler(maxSampleCount int) Sampler {
	s := new(UniformSampler)
	s.Samples = make([][2]float32, maxSampleCount)
	return s
}
