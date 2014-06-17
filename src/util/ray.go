// ray.go
package util

import (
	"github.com/ungerik/go3d/vec3"
)

type Ray struct {
	Origin    vec3.T
	Direction vec3.T
}

func (R *Ray) PointAt(t float32) vec3.T {
	scaled_direction := R.Direction.Scaled(t)
	return vec3.Add(&R.Origin, &scaled_direction)
}
