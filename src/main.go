package main

import (
	"fmt"
	"github.com/ungerik/go3d/vec3"
	"intersectables"
	"util"
	"cameras"
)

func main() {
	r := util.Ray{vec3.Zero, vec3.UnitXYZ}
	fmt.Println(r.PointAt(5.0))
	s := intersectables.Sphere{vec3.Zero, 5}
	fmt.Println(s)
	fmt.Println(s.Intersect(&r))
	t0, t1, hasSolution := util.SolveQuadratic(1, 0, -4)
	fmt.Println(t0, t1, hasSolution)
	eye := vec3.T{0, 0, 5}
	lookAt := vec3.T{0, -5, 0}
	up := vec3.T{0, 1, 0}
	var fov float32 = 60.0
	var aspect float32 = 16.0 / 9.0
	width := 640
	height := 360
	cameras.MakePinholeCamera(&eye, &lookAt, &up, fov, aspect, width, height)
}
