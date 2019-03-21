package scenes

import (
	"github.com/panmari/gort/cameras"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/integrators"
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/intersectables/csg"
	"github.com/panmari/gort/lights"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/samplers"
	"github.com/panmari/gort/tonemappers"
	"github.com/ungerik/go3d/vec3"
)

func NewSimpleCSGScene() Scene {
	var s Scene
	s.Filename = "test_scene_csg"
	s.SPP = 8
	eye := vec3.T{0, 0, 5}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	fov := float32(60.0)
	var aspect float32 = 16.0 / 9.0
	width := 640
	height := 360
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.Sampler = samplers.NewRandomSampler(42, s.SPP)
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	redMaterial := materials.MakeDiffuseMaterial(vec3.T{1, 0, 0})
	blueMaterial := materials.MakeDiffuseMaterial(vec3.T{0, 0, 1})
	yellowMaterial := materials.MakeDiffuseMaterial(vec3.T{1, 1, 0})
	reflectiveMaterial := materials.NewReflective(vec3.T{1, 1, 1})

	list := intersectables.NewIntersectableList(2)
	sphere1 := csg.NewSphere(vec3.T{-3, 0, 0}, 1.0, blueMaterial)
	sphere2 := csg.NewSphere(vec3.T{-3, 1, 0}, 1.0, blueMaterial)
	pill := csg.NewNode(sphere1, sphere2, csg.INTERSECT)
	list.Add(pill)

	sphere := csg.NewSphere(vec3.T{0, 0, 0}, 1.0, redMaterial)
	cutNormal := vec3.T{0.3, .5, .1}
	cutNormal.Normalize()
	plane := csg.NewPlane(cutNormal, -.5, redMaterial)
	node1 := csg.NewNode(sphere, plane, csg.INTERSECT)
	cut2Normal := cutNormal.Scaled(-1)
	plane2 := csg.NewPlane(cut2Normal, -.5, redMaterial)
	drum := csg.NewNode(node1, plane2, csg.INTERSECT)
	list.Add(drum)

	list.Add(intersectables.NewSphere(vec3.T{2.5, -0.1, 0}, 1.0, reflectiveMaterial))
	list.Add(intersectables.NewSphere(vec3.T{2.6, -0.8, 2}, 0.5, yellowMaterial))
	porcelanMaterial := materials.NewBlinn(vec3.T{0.5, 0.5, 0.5}, vec3.T{0.5, 0.6, 0.6}, 80)
	gridMaterial := materials.NewGrid(porcelanMaterial, materials.NewDiffuseMaterial(1, 1, 1), 0.95, 0.05)
	list.Add(intersectables.NewPlane(vec3.T{0, 0, 1}, 4, gridMaterial))
	list.Add(intersectables.NewPlane(vec3.T{0, 1, 0}, 2, gridMaterial))
	s.Root = intersectables.NewAggregate(list)

	// s.Integrator = integrators.MakeDebugIntegrator(s.Root)
	l := make([]lights.LightGeometry, 0, 2)
	l = append(l, lights.MakePointLight(vec3.T{0, 2, 0}, vec3.T{10, 10, 10}))
	l = append(l, lights.MakePointLight(vec3.T{-3, 2, 0}, vec3.T{10, 10, 10}))
	s.Integrator = integrators.MakeWhittedIntegrator(s.Root, l)

	return s
}
