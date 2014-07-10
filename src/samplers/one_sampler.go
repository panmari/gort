package samplers

import ()

type OneSampler struct {}

func (o *OneSampler) Get2DSamples(n int) [][2]float32 {
	samples := make([][2]float32, n)
	for i := range samples {
		samples[i] = [2]float32{0.5, 0.5}
	}
	return samples
}

func MakeOneSampler(seed int64) Sampler {
	return new(OneSampler)
}
