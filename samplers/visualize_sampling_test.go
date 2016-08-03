package samplers

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func writeSamplesToImage(s Sampler, nrSamples int, filename string) {
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	samples := s.Get2DSamples(100)
	for i := range samples {
		x := int(samples[i][0] * 100)
		y := int(samples[i][1] * 100)
		img.SetGray(x, y, color.Gray{255})
	}
	fo, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}
	if err = png.Encode(fo, img); err != nil {
		panic(err)
	}

}

func TestVisualizeLowDiscrepancySampling(t *testing.T) {
	s := NewLowDiscrepancySampler(100)
	writeSamplesToImage(s, 100, "low_discrepancy_samples")
}

func TestVisualizeRandomSampling(t *testing.T) {
	s := NewRandomSampler(0, 100)
	writeSamplesToImage(s, 100, "random_samples")
}

func TestVisualizeUniformSampling(t *testing.T) {
	s := NewUniformSampler(100)
	writeSamplesToImage(s, 100, "uniform_samples")
}
