package scenes

import (
	"cameras"
	"samplers"
	"integrators"
	"films"
	"github.com/ungerik/go3d/vec3"
	"intersectables"
)

func MakeSimpleScene() *Scene {
	eye := vec3.T{0, 0, 5}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	var fov float32 = 60.0
	var aspect float32 = 16.0 / 9.0
	width := 640
	height := 360
	c := cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	s := samplers.MakeRandomSampler()
	i := integrators.MakeDebugIntegrator()
	f := films.MakeBoxFilterFilm(width, height)
	
	root := intersectables.MakeIntersectableList(2)
	root.Add(intersectables.Sphere{Center: vec3.T{0,0,0}, Radius: 1.0})
	root.Add(intersectables.Sphere{Center: vec3.T{2,0,0}, Radius: 1.0})
	root.Add(intersectables.Sphere{Center: vec3.T{-3,0,0}, Radius: 1.0})
	
	return &Scene{Camera: c, Sampler: s, Integrator: i, Film: f, Root: root, SPP: 8}
}