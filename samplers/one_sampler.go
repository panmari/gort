package samplers

import ()

type OneSampler struct {
	samples [][2]float32
}

func (o *OneSampler) Get2DSamples(n int) [][2]float32 {
	o.samples = o.samples[0:n] // adapt size of slice
	for i := range o.samples {
		o.samples[i][0] = 0.5
		o.samples[i][1] = 0.5
	}
	return o.samples
}

func (o *OneSampler) DuplicateAndSeed(seed int64, maxSampleCount int) Sampler {
	return NewOneSampler(maxSampleCount)
}

func NewOneSampler(maxSampleCount int) Sampler {
	s := new(OneSampler)
	s.samples = make([][2]float32, maxSampleCount)
	return s
}
