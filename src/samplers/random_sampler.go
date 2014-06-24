package samplers

import (
	"math/rand"
)

type RandomSampler struct {
	rand rand.Rand
}

func (r *RandomSampler) Get2DSample() *[2]float32 {
	return &[2]float32{r.rand.Float32(), r.rand.Float32()}
}

//TODO: seed?
func MakeRandomSampler(seed int64) Sampler {
	return &RandomSampler{*rand.New(rand.NewSource(seed))}
}
