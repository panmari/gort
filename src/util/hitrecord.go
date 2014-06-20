package util

import (
	"github.com/ungerik/go3d/vec3"
)

type Hitrecord struct {
	T             float32
	Point         vec3.T
	Normal        vec3.T
	W_in          vec3.T
	U             float32
	V             float32
	Intersectable Intersectable
}

type Intersectable interface {
	Intersect(r *Ray) (*Hitrecord, bool)
}
