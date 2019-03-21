package util

import (
	"fmt"

	"github.com/ungerik/go3d/vec3"
)

type Hitrecord struct {
	T             float32
	Position      vec3.T
	Normal        vec3.T
	W_in          vec3.T
	U             float32
	V             float32
	Intersectable Intersectable
	Material      Material
}

type ShadingSample struct {
	BRDF vec3.T
	// Sampled direction.
	W           vec3.T
	Probability float32
}

func (h *Hitrecord) String() string {
	return fmt.Sprintf("Position: %v", h.Position)
}

type Intersectable interface {
	Intersect(r *Ray) *Hitrecord
	BoundingBox() *vec3.Box
}
