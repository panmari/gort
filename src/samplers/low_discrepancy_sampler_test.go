package samplers

import (
	math "github.com/barnex/fmath"
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

func TestModFunction(t *testing.T) {
	r := math.Mod(5, 2)
	if r != 1 {
		t.Error("5 % 2 should be 1, but is ", r)
	}

	r = math.Mod(5, 0.5)
	if r != 0 {
		t.Error("5 % 0.5 should be 0, but is ", r)
	}

	r = math.Mod(0.5, 0.2)
	if math.Abs(r-0.1) > EPSILON {
		t.Error("0.5 % 0.2 should be 0.1, but is ", r)
	}

	r = math.Mod(0.2, 1)
	if math.Abs(r-0.2) > EPSILON {
		t.Error("0.2 % 1 should be 0.2, but is ", r)
	}

	// This is an issue!
	r = math.Mod(2, 3)
	if math.Abs(r-2) > EPSILON {
		t.Error("2 % 3 should be 2, but is ", r)
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
