package scenes

import (
	"cameras"
	"samplers"
	"integrators"
	"films"
	"util"
)

type Scene struct {
	Camera 		cameras.Camera
	Sampler 	func(seed int64) samplers.Sampler
	Integrator 	integrators.Integrator
	Film		films.Film
	Root 		util.Intersectable
	SPP 		int
	Filename 	string
}