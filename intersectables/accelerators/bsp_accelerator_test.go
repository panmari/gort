package accelerators

import (
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
	"testing"
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
func BenchmarkBoundingBoxIntersectionHit(b *testing.B) {
	r := util.Ray{vec3.Zero, vec3.UnitX} // hits
	box := vec3.Box{vec3.T{2, 0, 0}, vec3.T{3, 1, 1}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doesRayIntersectBox(&r, &box)
	}
}

// This should be faster since we can return early
func BenchmarkBoundingBoxIntersectionNoHit(b *testing.B) {
	r2 := util.Ray{vec3.Zero, vec3.UnitY} // does not hit
	box := vec3.Box{vec3.T{2, 0, 0}, vec3.T{3, 1, 1}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doesRayIntersectBox(&r2, &box)
	}
}

// Negative direction.
func BenchmarkBoundingBoxIntersectionHitNegative(b *testing.B) {
	r := util.Ray{vec3.Zero, vec3.T{-1, 0, 0}} // hits
	box := vec3.Box{vec3.T{-3, -1, -1}, vec3.T{-2, 0, 0}}

	min, max, hits := doesRayIntersectBox(&r, &box)
	if !hits {
		b.Error("Should intersect")
	}

	if min != 2 {
		b.Errorf("Min at %f", min)
	}

	if max != 3 {
		b.Errorf("Max at %f", max)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doesRayIntersectBox(&r, &box)
	}
}
