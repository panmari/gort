package scenes

import (
	"util/obj"
	"cameras"
	"films"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/mat4"
	"github.com/barnex/fmath"
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
	
	p1 := intersectables.NewPlane(vec3.UnitY, 1, materials.DiffuseDefault)
	p2 := intersectables.NewPlane(vec3.UnitZ, 1, materials.DiffuseDefault)
	p3 := intersectables.NewPlane(vec3.T{-1,0,0}, 1, materials.DiffuseDefault)
	p4 := intersectables.NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)
	p5 := intersectables.NewPlane(vec3.T{0,-1,0}, 1, materials.DiffuseDefault)
	data := obj.Read("obj/teapot.obj", 1)
	teapot := intersectables.NewMeshAggregate(data, materials.DiffuseDefault)
	
	t1 := mat4.Ident
	t1.Mul(0.5).SetTranslation(&vec3.T{0, -0.35, 0})
	teapotInstance := intersectables.NewDiffuseInstance(teapot, t1)
	
	t2 := mat4.Ident
	t2.Mul(0.5).SetTranslation(&vec3.T{0, 0.25, 0})
	rot := mat4.Zero
	rot.AssignXRotation(30*fmath.Pi/180)
	t2.MultMatrix(&rot)
	teapotInstance2 := intersectables.NewDiffuseInstance(teapot, t2)
	
	list := intersectables.NewIntersectableList(6)
	list.Add(p1, p2, p3, p4, p5, teapotInstance, teapotInstance2)
	
	s.Root = intersectables.NewAggregate(list)
	
	l := []lights.LightGeometry{ 	lights.MakePointLight(vec3.T{0,0.8,0.8}, vec3.T{3, 3, 3}),
									lights.MakePointLight(vec3.T{-0.8,0.2,1}, vec3.T{1.5, 1.5, 1.5})}
		
	s.Integrator = integrators.MakePointLightIntegrator(s.Root, l)
	return s
}

