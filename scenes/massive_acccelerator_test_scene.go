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
	"github.com/panmari/gort/util/obj"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
)

func NewMassiveAcceleratorTestScene() Scene {
	var s Scene
	s.Filename = "massive_accelerator_test_scene"
	s.SPP = 1

	width := 400
	height := 800

	eye := vec3.T{0, 0, 7}
	lookAt := vec3.T{0, 1, 0}
	up := vec3.T{0, 1, 0}
	fov := float32(60.0)
	aspect := float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.Sampler = samplers.NewOneSampler(s.SPP)
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	groundPlane := intersectables.NewPlane(vec3.UnitY, 1.5, materials.NewDiffuseMaterial(0, 0.8, 0.8))
	backPlane := intersectables.NewPlane(vec3.UnitZ, 3.15, materials.NewDiffuseMaterial(0.3, 0.8, 0.8))

	dataHeart := obj.Read("obj/xyzrgb_statuette.obj", 3)
	heart := intersectables.NewMeshAggregate(dataHeart, materials.DiffuseDefault)
	heart2 := accelerators.NewBSPAccelerator(heart)

	t2 := mat4.Ident
	t2.SetTranslation(&vec3.T{0, 1.5, 0})

	heartInst := intersectables.NewDiffuseInstance(heart2, t2)
	heartInst.Material = materials.DiffuseDefault

	list := intersectables.NewIntersectableList(6)
	list.Add(backPlane, groundPlane, heartInst)

	s.Root = intersectables.NewAggregate(list)

	pl1Pos := vec3.Add(&eye, &vec3.T{-1, 0, 0})
	pl1 := lights.MakePointLight(pl1Pos, vec3.T{14, 14, 14})
	pl2Pos := vec3.Add(&eye, &vec3.T{1, 0, 0})
	pl2 := lights.MakePointLight(pl2Pos, vec3.T{14, 14, 14})
	pl3 := lights.MakePointLight(vec3.T{0, 5, 1}, vec3.T{24, 24, 24})
	l := []lights.LightGeometry{pl1, pl2, pl3}
	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	return s
}
