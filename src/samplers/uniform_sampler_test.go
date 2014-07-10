package samplers

import (
	"testing"
)


func BenchmarkUniformSamplerOneSample(b *testing.B) {
	s := MakeUniformSampler(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(1)
	}
}

func BenchmarkUniformSamplerEightSamples(b *testing.B) {
	s := MakeUniformSampler(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(4)
	}
}