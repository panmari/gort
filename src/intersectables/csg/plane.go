package csg

import (
	"github.com/ungerik/go3d/vec3"
	"intersectables"
	"util"
	"math"
)

type Plane struct {
	plane  intersectables.Plane
}

func (p *Plane) GetIntervalBoundaries(r *util.Ray) *ByT {
	boundaries := make(ByT, 0, 2)
	hit, didHit := p.plane.Intersect(r)
	if (didHit) {
		b1 := IntervalBoundary{t: hit.T, hit: hit}
		b2 := IntervalBoundary{}
		if vec3.Dot(&p.plane.Normal, &r.Direction) < 0 {
			b1.isStart = true
			b2.isStart = false
			if hit.T > 0 {
				b2.t = math.MaxFloat32
			} else {
				b2.t = -math.MaxFloat32
			}
		} else {
			b1.isStart = false
			b2.isStart = true
			if hit.T > 0 {
				b2.t = -math.MaxFloat32
			} else {
				b2.t = math.MaxFloat32
			}
		}
		boundaries = append(boundaries, &b1, &b2)
	} 
	return &boundaries
}

func NewDiffusePlane(normal vec3.T, dist float32) *Solid {
	p := Solid{&Plane{*intersectables.MakeDiffusePlane(normal, dist)}}
	return &p
}
