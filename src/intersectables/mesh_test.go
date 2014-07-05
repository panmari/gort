package intersectables

import (
    "testing"
    "util"
    "github.com/ungerik/go3d/vec3"
    "util/obj"
)

func TestSingleTrianglesAdded(t *testing.T) {
	data := new(obj.Data)
	data.InsertLine("v 0 0 0")
	data.InsertLine("v 1 0 0")
	data.InsertLine("v 0 1 0")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("vn 0 0 1")
	data.InsertLine("f 0//0 1//1 2//2")
	t.Log(data)
	
	m := NewMeshAggregate(data)
	r := util.Ray{vec3.T{.2,.2, 1}, vec3.T{0,0,-1}}
	t.Log(m)
	hit := m.Intersect(&r)
	if hit == nil { 
		t.Errorf("Did not hit")
	}
	if hit.T != 1 {
		t.Errorf("Hit at wrong t: %f", hit.T)
	}
	expectedNormal := vec3.T{0,0,1}
	if hit.Normal != expectedNormal {
		t.Errorf("Wrong normal: %v", hit.Normal)
	}
	rNoHit := util.Ray{vec3.T{1, 1, 1}, vec3.T{0,0,-1}}
	if noHit := m.Intersect(&rNoHit); noHit != nil { 
		t.Errorf("Should not hit: %v", noHit)
	}
	skewed := util.Ray{vec3.T{0,0,1}, vec3.T{.1,.05,-1}}
	skewedHit := m.Intersect(&skewed)
	if skewedHit == nil { 
		t.Errorf("Did not hit")
	}
}

