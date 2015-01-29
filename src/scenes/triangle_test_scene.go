package scenes

import (
	"cameras"
	"films"
	"github.com/ungerik/go3d/vec3"
	"integrators"
	"intersectables"
	"materials"
	"samplers"
	"tonemappers"
	"util/obj"
)

func NewTriangleTestScene() Scene {
	var s Scene
	s.Filename = "triangle_test_scene"
	width := 512
	height := 512
	s.SPP = 1

	eye := vec3.T{0, 0, 3}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	var fov float32 = 60.0
	var aspect float32 = float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.Sampler = samplers.NewOneSampler(s.SPP)
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)

	data := new(obj.Data)
	data.InsertLine("v 0 0 0")
	data.InsertLine("v 1 0 0")
	data.InsertLine("v 0 1 0")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("vn 0 0 1")
	// The expected .obj format uses 1 based counting.
	data.InsertLine("f 1//1 2//2 3//3") 
	s.Root = intersectables.NewMeshAggregate(data, materials.DiffuseDefault)

	s.Integrator = integrators.MakeDebugIntegrator(s.Root)

	return s
}
