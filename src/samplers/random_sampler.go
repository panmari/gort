package samplers

import (
	"math/rand"
)

type RandomSampler struct {
	rand    rand.Rand
	Samples [][2]float32
}

func (r *RandomSampler) Get2DSamples(n int) [][2]float32 {
	r.Samples = r.Samples[0:n] // adapt size of slice
	for i := range r.Samples {
		r.Samples[i][0] = r.rand.Float32()
		r.Samples[i][1] = r.rand.Float32()
	}
	return r.Samples
}

func (r *RandomSampler) DuplicateAndSeed(seed int64) Sampler {
	return NewRandomSampler(seed, len(r.Samples))
}

func NewRandomSampler(seed int64, maxSampleCount int) Sampler {
	s := new(RandomSampler)
	s.rand = *rand.New(rand.NewSource(seed))
	s.Samples = make([][2]float32, maxSampleCount)
	return s
}
