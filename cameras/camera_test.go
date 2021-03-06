package cameras

import (
	"testing"

	"github.com/ungerik/go3d/vec3"
)

func TestCameraRay(t *testing.T) {
	eye := vec3.T{0, 0, 5}
	lookAt := vec3.T{0, -5, 0}
	up := vec3.T{0, 1, 0}
	fov := float32(60.0)
	aspect := float32(16.0 / 9.0)
	width := 640
	height := 360
	c := MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	sample := [2]float32{0.0, 0.0}
	r := c.MakeWorldSpaceRay(width/2, height/2, sample)
	expectedOrig := vec3.T{0, 0, 5}
	if r.Origin != expectedOrig {
		t.Error("wrong origin")
	}
	expectedDir := vec3.T{1.192093e-07, -0.7071069, -0.70710665}
	if vec3.Distance(&expectedDir, &r.Direction) > 0.001 {
		t.Errorf("wrong direction: %v vs %v, dist: %f", r.Direction, expectedDir, vec3.Distance(&expectedDir, &r.Direction))
	}
}
