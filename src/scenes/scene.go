package scenes

import (
	"cameras"
	"samplers"
	"integrators"
	"films"
)

type Scene interface {
	GetCamera() (*cameras.Camera)
	GetSampler() (*samplers.Sampler)
	GetIntegrator() (*integrators.Integrator)
	getFilm() (*films.Film)
}