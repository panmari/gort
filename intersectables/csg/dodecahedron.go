package csg

import (
	"github.com/barnex/fmath"
	"github.com/panmari/gort/materials"
	"github.com/ungerik/go3d/vec3"
)

// This is actually not a type, since it only is a helper that intersects planes and nodes
func NewDodecahedron() *Solid {
	var planes [12]*Solid
	m := materials.MakeDiffuseMaterial(vec3.Red)
	// Bottom half
	planes[0] = NewPlane(vec3.T{0, -1, 0}, -1, m)
	for i := 0; i < 5; i++ {
		// Make face normals, using facts that in a dodecahedron
		// - top and bottom faces are uniform pentagons
		// - dihedral angles between all faces are pi - arctan(2)
		theta := float32(i) * 2 * fmath.Pi / 5
		x := fmath.Sin(theta) * fmath.Sin(fmath.Atan(2))
		y := -fmath.Cos(fmath.Atan(2))
		z := fmath.Cos(theta) * fmath.Sin(fmath.Atan(2))
		normal := vec3.T{x, y, z}
		planes[i+1] = NewPlane(normal, -1, m)
	}

	// Top half
	planes[6] = NewPlane(vec3.T{0, 1, 0}, -1, m)
	for i := 0; i < 5; i++ {
		// Make face normals
		theta := (float32(i) + 0.5) * 2 * fmath.Pi / 5
		x := fmath.Sin(theta) * fmath.Sin(fmath.Atan(2))
		y := fmath.Cos(fmath.Atan(2))
		z := fmath.Cos(theta) * fmath.Sin(fmath.Atan(2))
		normal := vec3.T{x, y, z}
		planes[i+7] = NewPlane(normal, -1, m)
	}

	// Build CSG tree
	var nodes [6]*Solid
	for i := 0; i < 6; i++ {
		nodes[i] = NewNode(planes[2*i], planes[2*i+1], INTERSECT)
	}

	var nodes2 [3]*Solid
	for i := 0; i < 3; i++ {
		nodes2[i] = NewNode(nodes[2*i], nodes[2*i+1], INTERSECT)
	}

	almostRoot := NewNode(nodes2[0], nodes2[1], INTERSECT)

	root := NewNode(almostRoot, nodes2[2], INTERSECT)
	return root
}
