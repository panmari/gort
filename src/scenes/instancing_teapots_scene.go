package scenes

import (
	"util/obj"
	"cameras"
	"films"
	"github.com/ungerik/go3d/vec3"
	"integrators"
	"intersectables"
	"samplers"
	"tonemappers"
	"lights"
	"materials"
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
	var fov float32 = 60.0
	var aspect float32 = float32(width) / float32(height)
	s.Camera = cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s.Sampler = samplers.MakeOneSampler
	s.Film = films.MakeBoxFilterFilm(width, height, tonemappers.ClampToneMap)
	
	data := obj.Read("obj/teapot.obj", 1)
	s.Root = intersectables.NewMeshAggregate(data, &materials.DiffuseDefault)
	
	l := []lights.LightGeometry{ 	lights.MakePointLight(vec3.T{0,0.8,0.8}, vec3.T{3, 3, 3}),
									lights.MakePointLight(vec3.T{-0.8,0.2,1}, vec3.T{1.5, 1.5, 1.5})}
		
	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	return s
}

