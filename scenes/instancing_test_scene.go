package scenes

import (
	"github.com/panmari/gort/cameras"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/integrators"
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/intersectables/csg"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/samplers"
	"github.com/panmari/gort/tonemappers"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
)

func NewInstancingTestScene() Scene {
	var s Scene
	s.Filename = "instancing_test_scene"
	width := 1280
	height := 720
	eye := vec3.T{0, 0, 3}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	fov := float32(60.0)
	aspect := float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.SPP = 1
	s.Sampler = samplers.NewOneSampler(s.SPP)
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	list := intersectables.NewIntersectableList(2)
	sphere := csg.NewSphere(vec3.T{0, 0, 0}, 1.0, materials.DiffuseDefault)
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
