package samplers

import (
	"testing"
)

const (
	EPSILON = 0.00001
)

func BenchmarkLdSamplerOneSample(b *testing.B) {
	s := NewLowDiscrepancySampler(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(1)
	}
}

func BenchmarkLdSamplerFourSamples(b *testing.B) {
	s := NewLowDiscrepancySampler(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(4)
	}
}

func TestLdSamplerSamplings(t *testing.T) {
	s := NewLowDiscrepancySampler(4)
	samples := s.Get2DSamples(4)
	expected_samples := [][2]float32{[2]float32{0.5, 0.33333334}, [2]float32{0.25, 0.6666667},
		[2]float32{0.75, 0.11111111}, [2]float32{0.125, 0.44444445}}
	for i := range samples {
		if samples[i] != expected_samples[i] {
			t.Error(samples[i], expected_samples[i])
		}
	}
}
