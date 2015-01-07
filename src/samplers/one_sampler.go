package samplers

import ()

type OneSampler struct{
	samples [][2]float32
}

func (o *OneSampler) Get2DSamples(n int) [][2]float32 {
	for i := range o.samples {
		o.samples[i] = [2]float32{0.5, 0.5}
	}
	return o.samples
}

func NewOneSampler(seed int64, maxSampleCount int) Sampler {
	s := new(OneSampler)
	s.samples = make([][2]float32, maxSampleCount)
	return s
}
