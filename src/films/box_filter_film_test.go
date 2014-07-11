package films

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"tonemappers"
)

func TestBoxFilterFilm(t *testing.T) {
	f := MakeBoxFilterFilm(100, 100, tonemappers.ClampToneMap)
	f.AddSample(49, 49, &vec3.Red)
	f.AddSample(50, 50, &vec3.Green)
	f.AddSample(51, 51, &vec3.Blue)
	// Y axis is inverted!
	col := f.At(49, 99-49)
	if R, _, _, _ := col.RGBA(); R != 65535 {
		t.Errorf("Got wrong color of first pixel: ", R, col)
	}
	col = f.At(49, 99-49)
	if R, _, _, _ := col.RGBA(); R != 65535 {
		t.Errorf("Got wrong color of first pixel when polling second time: ", R, col)
	}
	col = f.At(50, 99-50)
	if _, G, _, _ := col.RGBA(); G != 65535 {
		t.Errorf("Got wrong color: ", G, col)
	}
	col = f.At(51, 99-51)
	if _, _, B, _ := col.RGBA(); B != 65535 {
		t.Errorf("Got wrong color: ", B, col)
	}
	col = f.At(52, 99-52)
	if R, G, B, A := col.RGBA(); R != 0 || G != 0 || B != 0 || A != 65535 {
		t.Errorf("Is not dark: ", R, G, B, A, col)
	}
	f.AddSample(51, 51, &vec3.T{0, 0, 1})
	f.AddSample(51, 51, &vec3.T{0, 0, 1})
	col = f.At(51, 99-51)
	if _, _, B, _ := col.RGBA(); B != 65535 {
		t.Errorf("Got wrong color after adding 3 samples: ", B, col)
	}

	f.AddSample(50, 50, &vec3.T{0, 0, 0})
	if _, G, _, _ := f.At(50, 99-50).RGBA(); G != 32639 {
		t.Errorf("Got wrong color after adding 2 different samples: ", G, col)
	}
}
