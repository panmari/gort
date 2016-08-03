package lights

import (
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type PointLight struct {
	Position vec3.T
	Material util.Material
}

func (p *PointLight) Sample(sample [2]float32) *util.Hitrecord {
	h := new(util.Hitrecord)
	h.Intersectable = p
	h.Position = p.Position
	h.Material = p.Material
	return h
}

// Can not be hit, so this returns always nil.
func (p *PointLight) Intersect(r *util.Ray) *util.Hitrecord {
	return nil
}

// Can not be hit, so no bounding box is necessary.
func (p *PointLight) BoundingBox() *vec3.Box {
	return nil
}

func MakePointLight(position, emission vec3.T) *PointLight {
	l := new(PointLight)
	l.Position = position
	l.Material = materials.MakePointLightMaterial(emission)
	return l
}
