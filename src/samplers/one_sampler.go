package samplers

import ()

type OneSampler struct {}

func (o *OneSampler) Get2DSample() *[2]float32 {
	return &[2]float32{0.5, 0.5}
}

func MakeOneSampler(seed int64) Sampler {
	return new(OneSampler)
}
