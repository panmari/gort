package lights

import (
	"github.com/ungerik/go3d/vec3"
	"util"
	"materials"
)

type PointLight struct {
	Position vec3.T
	Material materials.PointLightMaterial 
}

func (p PointLight) Sample(sample [2]float32) *util.Hitrecord {
	h := new(util.Hitrecord)
	h.Intersectable = &p
	h.Position = p.Position
	h.Material = p.Material
	return h
}

// can not be hit
func (p PointLight) Intersect(r *util.Ray) (*util.Hitrecord, bool) {
	return nil, false
}

func MakePointLight(position, emission vec3.T) PointLight {
	l := PointLight{}
	l.Position = position
	l.Material = materials.MakePointLightMaterial(emission)
	return l
}