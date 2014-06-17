package test
import (
	"github.com/ungerik/go3d/vec3"
	"github.com/barnex/fmath"
	
)
func HasTinyDifferenceVec3(a, b vec3.T) bool {
	return a.Distance(&b) < 0.0001
}

func HasTinyDifference(a, b float32) bool {
	return fmath.Abs(a-b) < 0.0001
}