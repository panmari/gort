package lights

import (
	"github.com/ungerik/go3d/vec3"
	"util"
)

type PointLight struct {
	Position *vec3.T
	Material *PointLightMaterial 
}

func (p *PointLight) Sample(sample [2]float32) *util.Hitrecord {
	return new(util.Hitrecord)
}

// can not be hit
func (p *PointLight) Intersect(r *util.Ray) *util.HitRecord {
	return new(util.Hitrecord)
}