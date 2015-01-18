package renderer

import (
	"cameras"
	"encoding/gob"
	"films"
	"integrators"
	"intersectables"
	"lights"
	"materials"
	"samplers"
)

func registerTypes() {
	gob.Register(&cameras.PinholeCamera{})
	gob.Register(&samplers.OneSampler{})
	gob.Register(&samplers.RandomSampler{})
	gob.Register(&integrators.PointLightIntegrator{})
	gob.Register(&films.BoxFilterFilm{})
	gob.Register(&intersectables.Aggregate{})
	gob.Register(&intersectables.Sphere{})
	gob.Register(&intersectables.Plane{})
	gob.Register(&lights.PointLight{})
	gob.Register(&materials.Diffuse{})
	gob.Register(&materials.PointLightMaterial{})
	gob.Register(&intersectables.IntersectableList{})
}
