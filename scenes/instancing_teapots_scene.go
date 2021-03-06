package scenes

import (
	"github.com/panmari/gort/cameras"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/integrators"
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/lights"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/samplers"
	"github.com/panmari/gort/tonemappers"
	"github.com/panmari/gort/util"
	"github.com/panmari/gort/util/obj"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
)

func NewInstancingTeapotsScene() Scene {
	var s Scene
	s.Filename = "instancing_teapots_scene"
	s.SPP = 1

	width := 256
	height := 256

	eye := vec3.T{0, 0, 2}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	fov := float32(60.0)
	aspect := float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.Sampler = samplers.NewOneSampler(s.SPP)
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	p1 := intersectables.NewPlane(vec3.UnitY, 1, materials.NewDiffuseMaterial(0, 0.8, 0.8))
	p2 := intersectables.NewPlane(vec3.UnitZ, 1, materials.NewDiffuseMaterial(0.3, 0.8, 0.8))
	p3 := intersectables.NewPlane(vec3.T{-1, 0, 0}, 1, materials.NewDiffuseMaterial(1, 0.8, 0.8))
	p4 := intersectables.NewPlane(vec3.UnitX, 1, materials.NewDiffuseMaterial(0, 0.8, 0))
	p5 := intersectables.NewPlane(vec3.T{0, -1, 0}, 1, materials.NewDiffuseMaterial(0.8, 0.8, 0.8))

	data := obj.Read("obj/teapot.obj", 1)
	teapot := intersectables.NewMeshAggregate(data, materials.DiffuseDefault)

	t1 := mat4.Ident
	t1.Scale(0.5).SetTranslation(&vec3.T{0, -0.35, 0})
	redBlinn := materials.NewBlinn(vec3.T{0.5, 0, 0}, vec3.T{0.6, 0.6, 0}, 40.0)
	teapotInstance := intersectables.NewInstance(teapot, t1, redBlinn)

	t2 := mat4.Ident
	t2.Scale(0.5).SetTranslation(&vec3.T{0, 0.25, 0})
	rot := mat4.Zero
	rot.AssignXRotation(util.ToRadians(30))
	t2.MultMatrix(&rot)
	grayDiffuse := materials.NewDiffuseMaterial(0.5, 0.5, 0.5)
	teapotInstance2 := intersectables.NewInstance(teapot, t2, grayDiffuse)

	list := intersectables.NewIntersectableList(6)
	list.Add(p1, p2, p3, p4, p5, teapotInstance, teapotInstance2)

	s.Root = intersectables.NewAggregate(list)

	l := []lights.LightGeometry{lights.MakePointLight(vec3.T{0, 0.8, 0.8}, vec3.T{3, 3, 3}),
		lights.MakePointLight(vec3.T{-0.8, 0.2, 1}, vec3.T{1.5, 1.5, 1.5})}

	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	return s
}
