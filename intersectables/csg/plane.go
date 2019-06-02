package csg

import (
	"math"

	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type Plane struct {
	plane intersectables.Plane
}

func (p *Plane) GetIntervalBoundaries(r *util.Ray) ByT {
	if hit := p.plane.IntersectHelper(r, true); hit != nil {
		b1 := IntervalBoundary{t: hit.T, hit: hit}
		b2 := IntervalBoundary{}
		if vec3.Dot(&p.plane.Normal, &r.Direction) < 0 {
			// The ray is pointing towards the plane, so it will exit at infinity.
			b1.isStart = true
			b2.isStart = false
			b2.t = float32(math.Inf(1))
		} else {
			// The ray is pointing away from the plane, so it entered at -infinity.
			b1.isStart = false
			b2.isStart = true
			b2.t = float32(math.Inf(-1))
		}
		return ByT{b1, b2}
	} else {
		// The ray is parallel to the plane. Check if it's fully contained.
	}
	return ByT{}
}

func (p *Plane) BoundingBox() *vec3.Box {
	return p.plane.BoundingBox()
}

// NewPlane creates a new plane with the given normal and the given distance to origin (measured along normal).
func NewPlane(normal vec3.T, dist float32, m util.Material) *Solid {
	plane := *intersectables.NewPlane(normal, dist, m)
	return &Solid{&Plane{plane}}
}

// NewDiffusePlane a plane the same way as NewPlane, with a default diffuse material.
func NewDiffusePlane(normal vec3.T, dist float32) *Solid {
	return NewPlane(normal, dist, materials.DiffuseDefault)
}
