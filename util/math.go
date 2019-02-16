package util

import (
	"github.com/barnex/fmath"
)

// SolveQuadratic solves the quadratic equation
// ax^2 + bx + c = 0
// Returns both solutions and true if a solution exists. Otherwise false.
func SolveQuadratic(a, b, c float32) (float32, float32, bool) {
	disc := b*b - 4*a*c
	if disc <= 0 {
		return 0, 0, false
	}

	rootDisc := fmath.Sqrt(disc)

	// Numerical magic copied from PBRT.
	var q float32
	if b < 0 {
		q = (b - rootDisc) / -2
	} else {
		q = (b + rootDisc) / -2
	}

	t0 := q / a
	t1 := c / q

	// Make t0 always the intersection closer to the camera, i.e. the smaller
	// value.
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	return t0, t1, true
}

func ToRadians(degrees float32) float32 {
	return fmath.Pi * degrees / 180.0
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
