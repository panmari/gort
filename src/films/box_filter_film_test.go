package films

import (
	"testing"
	"github.com/ungerik/go3d/vec3"
)

func TestBoxFilterFilm(t *testing.T) {
	f := MakeBoxFilterFilm(100, 100)
	f.addSample(49, 49, &vec3.T{1,0,0}) 
	f.addSample(50, 50, &vec3.T{0,1,0}) 
	f.addSample(51, 51, &vec3.T{0,0,1}) 
	col := f.At(49, 49)
	if R, _, _, _ := col.RGBA(); R != 65535 {
		t.Errorf("Got wrong color of first pixel: ", R, col)
	}
	col = f.At(49, 49)
	if R, _, _, _ := col.RGBA(); R != 65535 {
		t.Errorf("Got wrong color of first pixel when polling second time: ", R, col)
	}
	col = f.At(50, 50)
	if _, G, _, _ := col.RGBA(); G != 65535 {
		t.Errorf("Got wrong color: ", G, col)
	}
	col = f.At(51, 51)
	if _, _, B, _ := col.RGBA(); B != 65535 {
		t.Errorf("Got wrong color: ", B, col)
	}
	col = f.At(52, 52)
	if R, G, B, A := col.RGBA(); R != 0 || G != 0 || B != 0 || A != 65535 {
		t.Errorf("Is not dark: ", R, G, B, A, col)
	}
	f.addSample(51, 51, &vec3.T{0,0,1}) 
	f.addSample(51, 51, &vec3.T{0,0,1}) 
	col = f.At(51, 51)
	if _, _, B, _ := col.RGBA(); B != 65535 {
		t.Errorf("Got wrong color after adding 3 samples: ", B, col)
	}
	
	f.addSample(50, 50, &vec3.T{0,0,0}) 
	if _, G, _, _ := f.At(50, 50).RGBA(); G != 32639 {
		t.Errorf("Got wrong color after adding 2 different samples: ", G, col)
	}
	
	f.WriteToPng("test_output")
}
