package renderer

// TODO: Write a script to automatically generate this file via `go generate`.
import (
	"encoding/gob"
	"github.com/panmari/gort/cameras"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/integrators"
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/lights"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/samplers"
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
