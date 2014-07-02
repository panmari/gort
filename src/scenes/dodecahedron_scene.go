package scenes

import (
	"cameras"
	"films"
	"github.com/ungerik/go3d/vec3"
	"integrators"
	"intersectables/csg"
	"lights"
	"samplers"
)

func NewDodecahedronScene() Scene {
	width := 512
	height := 512
	c := cameras.NewFixedCamera(width, height)
	spp := 1
	s := samplers.MakeOneSampler
	f := films.MakeBoxFilterFilm(width, height)

	n1 := csg.NewNode(csg.NewDiffusePlane(vec3.T{0, 1, 0}, 1), 
		csg.NewDodecahedron(), csg.ADD)
	root := csg.NewNode(n1, csg.NewDiffusePlane(vec3.T{0, 0, 1}, 1), csg.ADD)

	//i := integrators.MakeDebugIntegrator(root)
	l := make([]lights.LightGeometry, 0, 2)
	l = append(l, lights.MakePointLight(vec3.T{0, 2, 0}, vec3.T{10, 10, 10}))
	l = append(l, lights.MakePointLight(vec3.T{-3, 2, 0}, vec3.T{10, 10, 10}))
	i := integrators.MakePointLightIntegrator(root, l)

	return Scene{Camera: c, Sampler: s, Integrator: i, Film: f, Root: root, SPP: spp, Filename: "test_scene_csg"}
}
