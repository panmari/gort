// ray.go
package util

import (
	"github.com/ungerik/go3d/vec3"
	"sync"
)

type Ray struct {
	Origin    vec3.T
	Direction vec3.T
}

func (R *Ray) PointAt(t float32) vec3.T {
	scaled_direction := R.Direction.Scaled(t)
	return vec3.Add(&R.Origin, &scaled_direction)
}

func MakeEpsilonRay(origin, direction *vec3.T) *Ray {
	epsilonOrig := direction.Scaled(0.0001)
	epsilonOrig.Add(origin)
	return NewRay(origin, &epsilonOrig)
}

func NewRay(origin, direction *vec3.T) *Ray {
	r := Raypool.Get().(*Ray)
	r.Origin = *origin
	r.Direction = *direction
	return r
}

var Raypool = sync.Pool{New: func() interface{} { return new(Ray) } }