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

func NewBoxScene() Scene {
	var s Scene
	s.Filename = "box_scene"
	width := 512
	height := 512
	s.Camera = cameras.NewFixedCamera(width, height)
	s.SPP = 1
	s.Sampler = samplers.NewOneSampler
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	// Define the root object (an intersectable) of the scene
	// A box
	p1 := csg.NewDiffusePlane(vec3.T{1, 0, 0}, 1)
	p2 := csg.NewDiffusePlane(vec3.T{-1, 0, 0}, 1)
	p3 := csg.NewDiffusePlane(vec3.T{0, 1, 0}, 1)
	p4 := csg.NewDiffusePlane(vec3.T{0, -1, 0}, 1)
	p5 := csg.NewDiffusePlane(vec3.T{0, 0, 1}, 1)

	n1 := csg.NewNode(p1, p2, csg.ADD)
	n2 := csg.NewNode(p3, p4, csg.ADD)
	n3 := csg.NewNode(n2, p5, csg.ADD)
	s.Root = csg.NewNode(n1, n3, csg.ADD)

	l := make([]lights.LightGeometry, 0, 2)
	l = append(l, lights.MakePointLight(vec3.T{0, 0, 3}, vec3.T{10, 10, 10}))
	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	//s.Integrator = integrators.MakeDebugIntegrator(root)

	return s
}
