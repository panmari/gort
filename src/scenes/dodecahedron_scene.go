package scenes

import (
	"cameras"
	"films"
	"github.com/ungerik/go3d/vec3"
	"integrators"
	"intersectables/csg"
	"lights"
	"samplers"
	"tonemappers"
)

func NewDodecahedronScene() Scene {
	var s Scene
	s.Filename = "dodecahedron_scene"
	width := 512
	height := 512
	s.Camera = cameras.NewFixedCamera(width, height)
	s.SPP = 1
	s.Sampler = samplers.NewOneSampler
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	n1 := csg.NewNode(csg.NewDiffusePlane(vec3.T{0, 1, 0}, 1), csg.NewDodecahedron(), csg.ADD)
	s.Root = csg.NewNode(n1, csg.NewDiffusePlane(vec3.T{0, 0, 1}, 1), csg.ADD)

	l := make([]lights.LightGeometry, 0, 1)
	l = append(l, lights.MakePointLight(vec3.T{0, 0, 3}, vec3.T{15, 15, 15}))
	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	//s.Integrator = integrators.MakeDebugIntegrator(s.Root)

	return s
}
