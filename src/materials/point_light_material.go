package materials

import (
	"github.com/ungerik/go3d/vec3"
)

type PointLightMaterial struct {
	emission *vec3.T
}

func (m *PointLightMaterial) HasSpecularReflection() bool {
	return false
}

func (m *PointLightMaterial) HasSpecularRefraction() bool {
	return false
}

func MakePointLightMaterial(emission *vec3.T) *PointLightMaterial {
	m := new(PointLightMaterial)
	m.emission = emission
	return m
}