package util

import (
	"github.com/barnex/fmath"
)

func SolveQuadratic(a, b, c float32) (float32, float32, bool) {
	disc := b*b - 4*a*c
	if disc <= 0 {
		return 0, 0, false
	}

	rootDisc := fmath.Sqrt(disc)

	// numerical magic copied from PBRT:
	var q float32
	if b < 0 {
		q = (b - rootDisc) / -2
	} else {
		q = (b + rootDisc) / -2
	}

	t0 := q / a
	t1 := c / q

	//make t0 always the intersection closer to the camera
	if t0 > t1 {
		swap := t0
		t0 = t1
		t1 = swap
	}
	return t0, t1, true
}

func ToRadians(degrees float32) float32 {
	return fmath.Pi * degrees / 180.0
}
