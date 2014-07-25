package samplers

import (
	"testing"
)

func BenchmarkRandomSamplerOneSample(b *testing.B) {
	s := MakeRandomSampler(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(1)
	}
}

func BenchmarkRandomSamplerFourSamples(b *testing.B) {
	s := MakeRandomSampler(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(4)
	}
}
