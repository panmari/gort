package samplers

import (
	math "github.com/barnex/fmath"
)

// Generates a (2, 3) Halton sequence, ignoring the seed.
// see http://en.wikipedia.org/wiki/Halton_sequence.
type LowDiscrepancySampler struct {
	Samples      [][2]float32
	BaseX, BaseY int
}

func (ld *LowDiscrepancySampler) Get2DSamples(n int) [][2]float32 {
	ld.Samples = ld.Samples[0:n] // adapt size of slice
	for i := range ld.Samples {
		ld.Samples[i][0] = halton(i+1, ld.BaseX)
		ld.Samples[i][1] = halton(i+1, ld.BaseY)
	}
	return ld.Samples
}

func halton(index, base int) float32 {
	result := float32(0)
	f := 1 / float32(base)
	for i := float32(index); i > 0; f /= float32(base) {
		result += f * math.Mod(i, float32(base))
		i = math.Floor(i / float32(base))
	}
	return result
}

func (ld *LowDiscrepancySampler) DuplicateAndSeed(seed int64, maxSampleCount int) Sampler {
	return NewLowDiscrepancySampler(maxSampleCount)
}

func NewLowDiscrepancySampler(maxSampleCount int) Sampler {
	s := new(LowDiscrepancySampler)
	s.BaseX = 2
	s.BaseY = 3
	s.Samples = make([][2]float32, maxSampleCount)
	return s
}
