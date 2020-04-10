package intersectables

import (
	"testing"

	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/panmari/gort/util/obj"
	"github.com/ungerik/go3d/vec3"
)

func TestSingleTrianglesAdded(t *testing.T) {
	data := new(obj.Data)
	data.InsertLine("v 0 0 0")
	data.InsertLine("v 1 0 0")
	data.InsertLine("v 0 1 0")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("f 1//1 2//2 3//3")

	m := NewMeshAggregate(data, materials.DiffuseDefault)
	r := util.Ray{Origin: vec3.T{.2, .2, 1}, Direction: vec3.T{0, 0, -1}}
	hit := m.Intersect(&r)
	if hit == nil {
		t.Errorf("Did not hit")
	}
	if hit.T != 1 {
		t.Errorf("Hit at wrong t: %f", hit.T)
	}
	expectedNormal := vec3.T{0, 0, 1}
	if hit.Normal != expectedNormal {
		t.Errorf("Wrong normal: %v", hit.Normal)
	}
	rNoHit := util.Ray{Origin: vec3.T{1, 1, 1}, Direction: vec3.T{0, 0, -1}}
	if noHit := m.Intersect(&rNoHit); noHit != nil {
		t.Errorf("Should not hit: %v", noHit)
	}
	skewed := util.Ray{Origin: vec3.T{0, 0, 1}, Direction: vec3.T{.1, .05, -1}}
	skewedHit := m.Intersect(&skewed)
	if skewedHit == nil {
		t.Errorf("Did not hit")
	}
}

func TestTeapotMesh(t *testing.T) {
	data := obj.Read("../obj/teapot.obj", 1)
	if len(data.Vertices) != 302 {
		t.Errorf("Wrong number of vertices: %d", len(data.Vertices))
	}
	if len(data.Normals) != 317 {
		t.Errorf("Wrong number of normals: %d", len(data.Normals))
	}

	m := NewMeshAggregate(data, materials.DiffuseDefault)
	r := util.Ray{Origin: vec3.T{.1, .1, 1}, Direction: vec3.T{0, 0, -2}}
	hit := m.Intersect(&r)
	if hit == nil {
		t.Fatalf("Did not hit")
	}
	if hit.T-0.404918 > 0.0001 {
		t.Errorf("Hit at wrong t: %f", hit.T)
	}
	expectedNormal := vec3.T{0.28876302, 0.3982444, 0.87064195}
	if hit.Normal != expectedNormal {
		t.Errorf("Wrong normal: %v", hit.Normal)
	}

}

func BenchmarkTeapotMesh(b *testing.B) {
	data := obj.Read("../obj/teapot.obj", 1)

	m := NewMeshAggregate(data, materials.DiffuseDefault)
	r := util.Ray{Origin: vec3.T{.1, .1, 1}, Direction: vec3.T{0, 0, -2}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Intersect(&r)
	}
}
