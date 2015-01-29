package renderer

// TODO: Write a script to automatically generate this file via `go generate`.
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
	gob.Register(&integrators.DebugIntegrator{})
	gob.Register(&integrators.PointLightIntegrator{})
	gob.Register(&films.BoxFilterFilm{})
	gob.Register(&intersectables.Aggregate{})
	gob.Register(&intersectables.Mesh{})
	gob.Register(&intersectables.MeshTriangle{})
	gob.Register(&intersectables.Plane{})
	gob.Register(&intersectables.Sphere{})
	gob.Register(&lights.PointLight{})
	gob.Register(&materials.Diffuse{})
	gob.Register(&materials.PointLightMaterial{})
	gob.Register(&intersectables.Instance{})
	gob.Register(&intersectables.IntersectableList{})
}
