package csg

import (
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type Sphere struct {
	sphere intersectables.Sphere
}

func (s *Sphere) GetIntervalBoundaries(r *util.Ray) ByT {
	boundaries := make(ByT, 0, 2)
	if t0, t1, hasSolutions := s.sphere.GetIntersections(r); hasSolutions {
		h0 := s.sphere.MakeHitrecord(t0, r)
		b0 := IntervalBoundary{t: t0, hit: h0, isStart: isBoundaryTypeStart(h0, r)}

		h1 := s.sphere.MakeHitrecord(t1, r)
		b1 := IntervalBoundary{t: t1, hit: h1, isStart: isBoundaryTypeStart(h1, r)}

		boundaries = append(boundaries, b0, b1)
	}
	return boundaries
}

func (s *Sphere) BoundingBox() *vec3.Box {
	return s.sphere.BoundingBox()
}

func NewSphere(center vec3.T, radius float32, m util.Material) *Solid {
	sphere := *intersectables.NewSphere(center, radius, m)
	csgSphere := Solid{&Sphere{sphere}}
	return &csgSphere
}
