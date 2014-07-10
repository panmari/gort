package samplers

import (
	"math/rand"
)

type RandomSampler struct {
	rand rand.Rand
}

func (r *RandomSampler) Get2DSamples(n int) [][2]float32 {
	samples := make([][2]float32, n)
	for i := range samples {
		samples[i] = [2]float32{r.rand.Float32(), r.rand.Float32()}
	}
	return samples
}

//TODO: seed?
func MakeRandomSampler(seed int64) Sampler {
	return &RandomSampler{*rand.New(rand.NewSource(seed))}
}
