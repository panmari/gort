package scenes

import (
	"github.com/panmari/gort/cameras"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/integrators"
	"github.com/panmari/gort/intersectables"
	"github.com/panmari/gort/intersectables/accelerators"
	"github.com/panmari/gort/lights"
	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/samplers"
	"github.com/panmari/gort/tonemappers"
	"github.com/panmari/gort/util"
	"github.com/panmari/gort/util/obj"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
)

const USE_ACCELERATOR = true

func NewAcceleratorTestScene() Scene {
	var s Scene
	s.Filename = "accelerator_test_scene"
	s.SPP = 1

	width := 1024
	height := 1024

	eye := vec3.T{0, 0, 2}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	var fov float32 = 60.0
	var aspect float32 = float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.Sampler = samplers.NewOneSampler(s.SPP)
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	p1 := intersectables.NewPlane(vec3.UnitY, 1, materials.NewDiffuseMaterial(0, 0.8, 0.8))
	p2 := intersectables.NewPlane(vec3.UnitZ, 1, materials.NewDiffuseMaterial(0.3, 0.8, 0.8))
	p3 := intersectables.NewPlane(vec3.T{-1, 0, 0}, 1, materials.NewDiffuseMaterial(1, 0.8, 0.8))
	p4 := intersectables.NewPlane(vec3.UnitX, 1, materials.NewDiffuseMaterial(0, 0.8, 0))
	p5 := intersectables.NewPlane(vec3.T{0, -1, 0}, 1, materials.NewDiffuseMaterial(0.8, 0.8, 0.8))

	dataHeart := obj.Read("obj/Heart.obj", 1)
	heart := intersectables.NewMeshAggregate(dataHeart, materials.DiffuseDefault)
	var heart2 util.Intersectable
	if USE_ACCELERATOR {
		heart2 = accelerators.NewBSPAccelerator(heart)
	} else {
		heart2 = heart
	}

	t2 := mat4.Ident
	t2.Scale(0.3).SetTranslation(&vec3.T{0, 0.25, 0})
	rot := mat4.Zero
	rot.AssignYRotation(util.ToRadians(90))
	t2.MultMatrix(&rot)

	heartInst := intersectables.NewDiffuseInstance(heart2, t2)
	heartInst.Material = materials.MakeDiffuseMaterial(vec3.T{1, 0.1, 0.1})

	list := intersectables.NewIntersectableList(6)
	list.Add(p1, p2, p3, p4, p5, heartInst)

	s.Root = intersectables.NewAggregate(list)

	l := []lights.LightGeometry{lights.MakePointLight(vec3.T{0, 0.8, 0.8}, vec3.T{3, 3, 3}),
		lights.MakePointLight(vec3.T{-0.8, 0.2, 1}, vec3.T{1.5, 1.5, 1.5})}

	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	return s
}
