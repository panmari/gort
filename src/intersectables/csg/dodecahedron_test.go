package csg

import (
	"testing"
	"util"
	"github.com/ungerik/go3d/vec3"
	"github.com/barnex/fmath"	
	"materials"
	"math"
)

func TestDodecahedronIntersection(t *testing.T) {
	d := NewDodecahedron()
	
	exactlyParallelRay := util.Ray{vec3.T{0,0,5}, vec3.T{0,0,-1}}
	hit := d.Intersect(&exactlyParallelRay)
	if hit != nil {
		t.Error("This does (strangely) not happen, since the ray was exactly parallel.")
	}
}

func TestFirstNodeIntersection(t *testing.T) {
	var planes [2]*Solid
	m := materials.MakeDiffuseMaterial(vec3.Red)
	// Bottom half
	planes[0] = NewPlane(vec3.T{0, -1, 0}, -1, m)
	for i := 0; i < 1; i++ {
		// Make face normals, using facts that in a dodecahedron
		// - top and bottom faces are uniform pentagons
		// - dihedral angles between all faces are pi - arctan(2)
		theta := float32(i) * 2 * fmath.Pi / 5
		x := fmath.Sin(theta) * fmath.Sin(fmath.Atan(2))
		y := -fmath.Cos(fmath.Atan(2))
		z := fmath.Cos(theta) * fmath.Sin(fmath.Atan(2))
		normal := vec3.T{x,y,z}
		planes[i + 1] = NewPlane(normal, -1, m)
	}
	n := NewNode(planes[0], planes[1], INTERSECT)
	r := util.Ray{vec3.T{0,0,5}, vec3.T{0,0.01,-1}}
	
	// ray is starting inside and ending inside this halfspace
	insidePlaneIbs := planes[0].GetIntervalBoundaries(&r)
	if len(insidePlaneIbs) != 2 {
		t.Fatalf("Not inside of halfspace defined by first plane: %d", len(insidePlaneIbs) ) 
	}
	if insidePlaneIbs[0].t != -100 {
		t.Errorf("Wrong t: %f", insidePlaneIbs[0].t) 
	}
	if insidePlaneIbs[1].t != float32(math.Inf(-1))  {
		t.Errorf("Wrong t: %f", insidePlaneIbs[1].t) 
	}	
	hit := n.Intersect(&r)
	if hit == nil {
		t.Error("Did not hit node") 
	}
}