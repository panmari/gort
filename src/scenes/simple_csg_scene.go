package scenes

import (
	"cameras"
	"films"
	"github.com/ungerik/go3d/vec3"
	"integrators"
	"intersectables"
	"intersectables/csg"
	"lights"
	"samplers"
)

func NewSimpleCSGScene() Scene {
	eye := vec3.T{0, 0, 5}
	lookAt := vec3.T{0, 0, 0}
	up := vec3.T{0, 1, 0}
	var fov float32 = 60.0
	var aspect float32 = 16.0 / 9.0
	width := 640
	height := 360
	c := cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
	// pass function to create new samplers
	s := samplers.MakeRandomSampler
	//s := samplers.MakeOneSampler()
	f := films.MakeBoxFilterFilm(width, height)

	root := intersectables.NewIntersectableList(2)
	sphere := csg.NewDiffuseSphere(vec3.T{0, 0, 0}, 1.0)
	cutNormal := vec3.T{0.3,.5,.1}
	cutNormal.Normalize()
	plane := csg.NewDiffusePlane(cutNormal, -.5)
	node1 := csg.NewNode(sphere, plane, csg.INTERSECT)
	cut2Normal := cutNormal.Scaled(-1)
	plane2 := csg.NewDiffusePlane(cut2Normal, -.5)
	node2 := csg.NewNode(node1, plane2, csg.INTERSECT)
	root.Add(node2)
	root.Add(intersectables.MakeDiffuseSphere(vec3.T{2, 1, 0}, 1.0))
	sphere1 := csg.NewDiffuseSphere(vec3.T{-3, 0, 0}, 1.0)
	sphere2 := csg.NewDiffuseSphere(vec3.T{-3, 1, 0}, 1.0)
	root.Add(csg.NewNode(sphere1, sphere2, csg.INTERSECT))
	root.Add(intersectables.MakeDiffusePlane(vec3.T{0, 0, 1}, 4))
	root.Add(intersectables.MakeDiffusePlane(vec3.T{0, 1, 0}, 2))

	//i := integrators.MakeDebugIntegrator(root)
	l := make([]lights.LightGeometry, 0, 2)
	l = append(l, lights.MakePointLight(vec3.T{0, 2, 0}, vec3.T{10, 10, 10}))
	l = append(l, lights.MakePointLight(vec3.T{-3, 2, 0}, vec3.T{10, 10, 10}))
	i := integrators.MakePointLightIntegrator(root, l)

	return Scene{Camera: c, Sampler: s, Integrator: i, Film: f, Root: root, SPP: 8, Filename: "test_scene_csg"}
}
