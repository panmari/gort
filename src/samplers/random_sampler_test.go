package samplers

import (
	"testing"
)

func BenchmarkRandomSamplerOneSample(b *testing.B) {
	s := NewRandomSampler(0, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(1)
	}
}

func BenchmarkRandomSamplerFourSamples(b *testing.B) {
	s := NewRandomSampler(0, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get2DSamples(4)
	}
}

func TestRandomSamplerTwoSamplingsAreDifferent(t *testing.T) {
	s := NewRandomSampler(0, 4)
	samples := s.Get2DSamples(1)
	// samples need to be copied
	samples_copy := make([][2]float32, len(samples))
	copy(samples, samples_copy)
	samples2 := s.Get2DSamples(1)
	if len(samples) != 1 {
		t.Errorf("Expected size 1, but was %d", len(samples))
	}
	if cap(samples) != 4 {
		t.Errorf("Expected cap 4, but was %d", cap(samples))
	}

	if samples_copy[0] == samples2[0] {
		t.Error("Samples are not different:", samples, samples2)
	}
	
	// samples and samples2 should contain the same arrays, since memory
	// was only allocated once.
	if samples[0] != samples2[0] {
		t.Error("Only one slice should be allocated: ", samples, samples2)
	}

}
