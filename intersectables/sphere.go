package intersectables

import (
	"github.com/barnex/fmath"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type Sphere struct {
	Center   vec3.T
	Radius   float32
	Material util.Material
	Box      vec3.Box
}

func (s *Sphere) GetIntersections(r *util.Ray) (float32, float32, bool) {
	a := r.Direction.LengthSqr()
	originCenter := vec3.Sub(&r.Origin, &s.Center)
	b := 2 * vec3.Dot(&r.Direction, &originCenter)
	c := originCenter.LengthSqr() - s.Radius*s.Radius
	return util.SolveQuadratic(a, b, c)
}

func (s *Sphere) Intersect(r *util.Ray) *util.Hitrecord {
	t0, t1, ok := s.GetIntersections(r)
	if !ok {
		return nil
	}
	if t0 > 0 {
		return s.MakeHitrecord(t0, r)
	}
	if t1 > 0 {
		return s.MakeHitrecord(t1, r)
	}
	// Hits are behind camera.
	return nil
}

func (s *Sphere) BoundingBox() *vec3.Box {
	return &s.Box
}

func (s *Sphere) MakeHitrecord(t float32, r *util.Ray) *util.Hitrecord {
	hitPoint := r.PointAt(t)
	normal := vec3.Sub(&hitPoint, &s.Center)
	normal.Normalize()

	//TODO: use a copy of this?
	wIn := r.Direction
	wIn.Normalize().Invert()

	u := 0.5 + fmath.Atan2(hitPoint[2], hitPoint[0])/(2*fmath.Pi)
	v := 0.5 - fmath.Asin(hitPoint[1])/fmath.Pi

	return &util.Hitrecord{t, hitPoint, normal, wIn, u, v, s, s.Material}
}

func MakeDiffuseSphere(center vec3.T, radius float32) *Sphere {
	s := new(Sphere)
	s.Center = center
	s.Radius = radius
	s.Material = materials.MakeDiffuseMaterial(vec3.T{1, 1, 1})
	radiusVector := vec3.T{radius, radius, radius}
	minBox := vec3.Sub(&center, &radiusVector)
	maxBox := vec3.Add(&center, &radiusVector)
	s.Box = vec3.Box{minBox, maxBox}
	return s
}
