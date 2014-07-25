package samplers

import (
	"testing"
)

func TestUniformSamplerOneSample(t *testing.T) {
	s := MakeUniformSampler(0)
	sample := s.Get2DSamples(1)
	if sample[0] != [2]float32{0.5, 0.5} {
		t.Error(sample)
	}
}

func TestUniformSamplerFourSample(t *testing.T) {
	s := MakeUniformSampler(0)
	sample := s.Get2DSamples(4)
	
	if sample[0] != [2]float32{0.25, 0.25} {
		t.Error(sample)
	}
	if sample[1] != [2]float32{0.75, 0.25} {
		t.Error(sample)
	}
	if sample[2] != [2]float32{0.25, 0.75} {
		t.Error(sample)
	}
	if sample[3] != [2]float32{0.75, 0.75} {
		t.Error(sample)
	}
}

func BenchmarkUniformSamplerOneSample(b *testing.B) {
	s := MakeUniformSampler(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(1)
	}
}

func BenchmarkUniformSamplerFourSamples(b *testing.B) {
	s := MakeUniformSampler(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(4)
	}
}
