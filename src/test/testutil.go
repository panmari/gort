package test

import (
	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/vec3"
)

func HasTinyDifferenceVec3(a, b *vec3.T) bool {
	return vec3.Distance(a, b) < 0.0001
}

func HasTinyDifference(a, b float32) bool {
	return fmath.Abs(a-b) < 0.0001
}
