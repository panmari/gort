package accelerators

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"util"
)

func TestDoesRayIntersectBox(t *testing.T) {
	r := util.Ray{vec3.Zero, vec3.UnitX}
	b := vec3.Box{vec3.T{2, 0, 0}, vec3.T{3, 1, 1}}
	min, max, hits := doesRayIntersectBox(&r, &b)

	if !hits {
		t.Error("Should intersect")
	}

	if min != 2 {
		t.Errorf("Min at %f", min)
	}

	if max != 3 {
		t.Errorf("Max at %f", max)
	}

	r2 := util.Ray{vec3.Zero, vec3.UnitY}
	_, _, hits = doesRayIntersectBox(&r2, &b)
	if hits {
		t.Error("Second ray should not hit")
	}
}

// Performance of BoundingBoxIntersection should be faster than other intersection tests.
func BenchmarkBoundingBoxIntersectionIntersectionHit(b *testing.B) {
	r := util.Ray{vec3.Zero, vec3.UnitX} // hits
	box := vec3.Box{vec3.T{2, 0, 0}, vec3.T{3, 1, 1}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doesRayIntersectBox(&r, &box)
	}
}

// This should be faster since we can return early
func BenchmarkBoundingBoxIntersectionIntersectionNoHit(b *testing.B) {
	r2 := util.Ray{vec3.Zero, vec3.UnitY} // does not hit
	box := vec3.Box{vec3.T{2, 0, 0}, vec3.T{3, 1, 1}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doesRayIntersectBox(&r2, &box)
	}
}
