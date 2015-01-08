package samplers

import (
	"github.com/barnex/fmath"
)

type UniformSampler struct {
	samples [][2]float32
}

/*
* Expects n to be a square number.
 */
func (r *UniformSampler) Get2DSamples(n int) [][2]float32 {
	k := fmath.Ceil(fmath.Sqrt(float32(n)))
	k_int := int(k)
	dist := 1 / k
	offset := 1 / (k * 2)
	r.samples = r.samples[0 : k_int*k_int] // adapt size of slice
	for i := range r.samples {
		r.samples[i][0] = offset + dist*float32(i%k_int)
		r.samples[i][1] = offset + dist*float32(i/k_int)
	}
	return r.samples
}

func NewUniformSampler(seed int64, maxSampleCount int) Sampler {
	s := new(UniformSampler)
	s.samples = make([][2]float32, maxSampleCount)
	return s
}
