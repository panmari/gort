package intersectables

import (
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type Plane struct {
	Normal         vec3.T
	DistanceOrigin float32
	Material       util.Material
}

func (s *Plane) Intersect(r *util.Ray) *util.Hitrecord {
	return s.IntersectHelper(r, false)
}

// Returns MaxBox in all cases, even though we could do better if the
// plane is axis aligned.
func (s *Plane) BoundingBox() *vec3.Box {
	return &vec3.MaxBox
}

func (s *Plane) IntersectHelper(r *util.Ray, allowNegative bool) *util.Hitrecord {
	tmp := vec3.Dot(&r.Direction, &s.Normal)

	if tmp == 0 { // parallel to plane, does not hit
		return nil
	}
	t := -(vec3.Dot(&s.Normal, &r.Origin) + s.DistanceOrigin) / tmp
	if !allowNegative && t <= 0 { // only allow negative for CSG objects
		return nil
	}
	hit := new(util.Hitrecord)
	hit.T = t
	hit.Position = r.PointAt(t)
	hit.Normal = s.Normal
	w := r.Direction.Scaled(-1)
	w.Normalize()
	hit.W_in = w
	hit.Material = s.Material
	hit.Intersectable = s
	return hit
}

// Creates a new plane with the given normal and the given distance to origin (measured along normal).
func NewPlane(normal vec3.T, distanceOrigin float32, m util.Material) *Plane {
	p := new(Plane)
	normal.Normalize()
	p.Normal = normal
	p.DistanceOrigin = distanceOrigin
	p.Material = m
	return p
}

// Same as NewPlane, but with default diffuse material.
func MakeDiffusePlane(normal vec3.T, distanceOrigin float32) *Plane {
	return NewPlane(normal, distanceOrigin, materials.DiffuseDefault)
}
