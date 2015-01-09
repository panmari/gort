package samplers

import (
	"math/rand"
)

type RandomSampler struct {
	rand    rand.Rand
	samples [][2]float32
}

func (r *RandomSampler) Get2DSamples(n int) [][2]float32 {
	r.samples = r.samples[0:n] // adapt size of slice
	for i := range r.samples {
		r.samples[i][0] = r.rand.Float32()
		r.samples[i][0] = r.rand.Float32()
	}
	return r.samples
}

func NewRandomSampler(seed int64, maxSampleCount int) Sampler {
	s := new(RandomSampler)
	s.rand = *rand.New(rand.NewSource(seed))
	s.samples = make([][2]float32, maxSampleCount)
	return s
}
