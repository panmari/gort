package scenes

import (
	"cameras"
	"films"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
	"integrators"
	"intersectables"
	"intersectables/csg"
	"samplers"
	"tonemappers"
)

func NewInstancingTestScene() Scene {
	var s Scene
	s.Filename = "instancing_test_scene"
	width := 1280
	height := 720
	eye := vec3.T{0, 0, 3}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	var fov float32 = 60.0
	var aspect float32 = float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.SPP = 1
	s.Sampler = samplers.MakeOneSampler
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	list := intersectables.NewIntersectableList(2)
	sphere := csg.NewDiffuseSphere(vec3.T{0, 0, 0}, 1.0)
	//assemble the first instance
	trans := mat4.Ident
	trans.SetTranslation(&vec3.T{2, 0, 0})
	sphere2 := intersectables.NewDiffuseInstance(sphere, trans)
	//assemble the second instance
	trans2 := mat4.Ident
	trans2.SetTranslation(&vec3.T{-2, 0, 0})
	sphere3 := intersectables.NewDiffuseInstance(sphere, trans2)

	list.Add(sphere, sphere2, sphere3)
	s.Root = intersectables.NewAggregate(list)

	s.Integrator = integrators.MakeDebugIntegrator(s.Root)
	return s
}
