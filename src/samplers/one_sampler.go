package samplers

import ()

type OneSampler struct {
	Samples [][2]float32
}

func (o *OneSampler) Get2DSamples(n int) [][2]float32 {
	o.Samples = o.Samples[0:n] // adapt size of slice
	for i := range o.Samples {
		o.Samples[i][0] = 0.5
		o.Samples[i][1] = 0.5
	}
	return o.Samples
}

func (o *OneSampler) DuplicateAndSeed(seed int64) Sampler {
	return NewOneSampler(len(o.Samples))
}

func NewOneSampler(maxSampleCount int) Sampler {
	s := new(OneSampler)
	s.Samples = make([][2]float32, maxSampleCount)
	return s
}
