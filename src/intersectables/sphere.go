package intersectables

import (
	"github.com/barnex/fmath"
	"github.com/ungerik/go3d/vec3"
	"util"
)

type Sphere struct {
	Center vec3.T
	Radius float32
}

func (s *Sphere) Intersect(r *util.Ray) *util.Hitrecord {
	a := r.Direction.LengthSqr()
	originCenter := vec3.Sub(&r.Origin, &s.Center)
	b := 2 * vec3.Dot(&r.Direction, &originCenter)
	c := originCenter.LengthSqr() - s.Radius*s.Radius
	t0, t1, hasSolution := util.SolveQuadratic(a, b, c)
	if hasSolution {
		if t0 > 0 {
			return s.makeHitrecord(t0, r)
		}
		if t1 > 0 {
			return s.makeHitrecord(t1, r)
		}
	}
	return &util.Hitrecord{}
}

func (s *Sphere) makeHitrecord(t float32, r *util.Ray) *util.Hitrecord {
	hitPoint := r.PointAt(t)
	normal := vec3.Sub(&hitPoint, &s.Center)
	normal.Normalize()

	//TODO: use a copy of this?
	wIn := r.Direction
	wIn.Normalize().Invert()

	u := 0.5 + fmath.Atan2(hitPoint[2], hitPoint[0])/(2*fmath.Pi)
	v := 0.5 - fmath.Asin(hitPoint[1])/fmath.Pi

	return &util.Hitrecord{t, hitPoint, normal, wIn, u, v, s}
}
