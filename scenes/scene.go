package scenes

import (
	"github.com/panmari/gort/cameras"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/integrators"
	"github.com/panmari/gort/samplers"
	"github.com/panmari/gort/util"
)

type Scene struct {
	Camera     cameras.Camera
	Sampler    samplers.Sampler
	Integrator integrators.Integrator
	Film       films.Film
	Root       util.Intersectable
	SPP        int
	Filename   string
}
