package films

import (
	"image/color"
	"testing"

	"github.com/panmari/gort/tonemappers"
	"github.com/ungerik/go3d/vec3"
)

func TestBoxFilterFilmAddsDiagonalOfDifferentValues(t *testing.T) {
	f := MakeBoxFilterFilm(100, 100, tonemappers.ClampToneMap)
	f.AddSample(49, 49, &vec3.Red)
	f.AddSample(50, 50, &vec3.Green)
	f.AddSample(51, 51, &vec3.Blue)

	testCases := []struct {
		x int
		// Y axis is inverted for reading out values.
		y    int
		want color.RGBA
	}{
		{
			x:    49,
			y:    99 - 49,
			want: color.RGBA{255, 0, 0, 255},
		},
		{
			x:    50,
			y:    99 - 50,
			want: color.RGBA{0, 255, 0, 255},
		},
		{
			x:    51,
			y:    99 - 51,
			want: color.RGBA{0, 0, 255, 255},
		},
		{
			x:    52,
			y:    99 - 52,
			want: color.RGBA{0, 0, 0, 255},
		},
	}

	for _, tc := range testCases {
		if got := f.At(tc.x, tc.y); got != tc.want {
			t.Errorf("Unexpected value at (%d, %d), got %v, want %v", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestBoxFilterFilmSumsValuesFromMultipleSamples(t *testing.T) {
	f := MakeBoxFilterFilm(1, 1, tonemappers.ClampToneMap)
	f.AddSample(0, 0, &vec3.T{0, 0, 1})
	f.AddSample(0, 0, &vec3.T{0, 0, 1})
	f.AddSample(0, 0, &vec3.T{1, 0, 0})

	testCases := []struct {
		x int
		// Y axis is inverted for reading out values.
		y    int
		want color.RGBA
	}{
		{
			x:    0,
			y:    0,
			want: color.RGBA{1 * 255 / 3, 0, 2 * 255 / 3, 255},
		},
	}

	for _, tc := range testCases {
		if got := f.At(tc.x, tc.y); got != tc.want {
			t.Errorf("Unexpected value at (%d, %d), got %v, want %v", tc.x, tc.y, got, tc.want)
		}
	}
}

func BenchmarkBoxFilterFilm(b *testing.B) {
	f := MakeBoxFilterFilm(100, 100, tonemappers.ClampToneMap)
	f.AddSample(49, 49, &vec3.Red)
	f.AddSample(50, 50, &vec3.Green)
	f.AddSample(51, 51, &vec3.Blue)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.At(99, 99)
		f.At(30, 30)
		f.At(50, 50)
		f.At(12, 82)
		f.At(33, 1)
	}
}
