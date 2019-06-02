package csg

import (
	"testing"

	"github.com/panmari/gort/materials"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

func TestNodeAdd(t *testing.T) {

	p := NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)
	s := NewSphere(vec3.Zero, 2, materials.DiffuseDefault)
	n := NewNode(p, s, ADD)

	testCases := []struct {
		name    string
		ray     util.Ray
		wantHit bool
		wantT   float32
	}{
		{
			name:    "parallel ray",
			ray:     util.Ray{Origin: vec3.T{0, 3, 0}, Direction: vec3.UnitY},
			wantHit: false,
		},
		{
			name:    "parallel ray hitting sphere",
			ray:     util.Ray{Origin: vec3.T{0, -3, 0}, Direction: vec3.UnitY},
			wantHit: true,
			wantT:   1,
		},
		{
			name:    "pointing away ray inside sphere",
			ray:     util.Ray{Origin: vec3.Zero, Direction: vec3.UnitX},
			wantHit: true,
			wantT:   2,
		},
		{
			name:    "pointing towards plane ray",
			ray:     util.Ray{Origin: vec3.T{4, 0, 0}, Direction: vec3.T{-1, 0, 0}},
			wantHit: true,
			wantT:   2,
		},
		{
			name:    "pointing towards plane ray far from origin",
			ray:     util.Ray{Origin: vec3.T{10, 10, 10}, Direction: vec3.T{-1, 0, 0}},
			wantHit: true,
			wantT:   11,
		},
	}
	for _, tc := range testCases {
		got := n.Intersect(&tc.ray)
		if gotHit := got != nil; gotHit != tc.wantHit {
			t.Errorf("n.Intersect(%q), got %v, want %v", tc.name, got, tc.wantHit)
		}
		if got != nil && got.T != tc.wantT {
			t.Errorf("n.Intersect(%q) unexpeded T value, got %v, want %v", tc.name, got.T, tc.wantT)
		}
	}
}

func TestNodeSubtract(t *testing.T) {
	p := NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)
	s := NewSphere(vec3.Zero, 2, materials.DiffuseDefault)
	n := NewNode(s, p, SUBTRACT)
	r := util.Ray{vec3.T{4, 0, 0}, vec3.T{-1, 0, 0}}
	ibs := n.GetIntervalBoundaries(&r)

	if len(ibs) < 2 {
		t.Errorf("There were fewer than two intersections: %d", len(ibs))
		t.FailNow()
	}
	if len(ibs) > 2 {
		t.Errorf("More than two intersections: %d", len(ibs))
	}
	if ibs[0].t != 2 {
		t.Errorf("Entering shape at wrong t: %f", ibs[0].t)
	}
	if !ibs[0].isStart {
		t.Error("First intersection is not entering")
	}

	if ibs[1].t != 5 {
		t.Errorf("Exiting shape at wrong t: %f", ibs[1].t)
	}
	if ibs[1].isStart {
		t.Error("Second intersection is not exiting")
	}
}

func TestNodeIntersect(t *testing.T) {
	p := NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)
	s := NewSphere(vec3.Zero, 2, materials.DiffuseDefault)
	n := NewNode(s, p, INTERSECT)
	r := util.Ray{Origin: vec3.T{4, 0, 0}, Direction: vec3.T{-1, 0, 0}}
	ibs := n.GetIntervalBoundaries(&r)

	if len(ibs) < 2 {
		t.Errorf("There were fewer than two intersections: %d", len(ibs))
		t.FailNow()
	}
	if len(ibs) > 2 {
		t.Errorf("More than two intersections: %d", len(ibs))
	}
	if ibs[0].t != 5 {
		t.Errorf("Entering shape at wrong t: %f", ibs[0].t)
	}
	if !ibs[0].isStart {
		t.Error("First intersection is not entering")
	}

	if ibs[1].t != 6 {
		t.Errorf("Exiting shape at wrong t: %f", ibs[1].t)
	}
	if ibs[1].isStart {
		t.Error("Second intersection is not exiting")
	}
}

func TestInfinitePlaneIntersection(t *testing.T) {
	p1 := NewDiffusePlane(vec3.T{1, 0, 0}, -1)
	p2 := NewDiffusePlane(vec3.T{-1, 0, 0}, -1)

	c := NewNode(p1, p2, INTERSECT)

	testCases := []struct {
		name    string
		ray     util.Ray
		wantHit bool
	}{
		{
			name:    "ray from left",
			ray:     util.Ray{Origin: vec3.T{-3, 0, 0}, Direction: vec3.T{1, 0, 0}},
			wantHit: true,
		},
		{
			name:    "ray from right",
			ray:     util.Ray{Origin: vec3.T{3, 0, 0}, Direction: vec3.T{-1, 0, 0}},
			wantHit: true,
		},
		{
			name:    "ray from within parallel to planes",
			ray:     util.Ray{Origin: vec3.T{0, 5, 0}, Direction: vec3.T{0, -1, 0}},
			wantHit: false,
		},
		{
			name:    "ray diagonal",
			ray:     util.Ray{Origin: vec3.T{5, 5, 5}, Direction: vec3.T{-1, -1, -1}},
			wantHit: true,
		},
		{
			name:    "ray from within",
			ray:     util.Ray{Origin: vec3.Zero, Direction: vec3.T{-1, 0, 0}},
			wantHit: true,
		},
	}

	for _, tc := range testCases {
		got := c.Intersect(&tc.ray)
		// t.Log(tc.name, got)
		if gotHit := got != nil; gotHit != tc.wantHit {
			t.Errorf("s.Intersect(%q), got %v, want %v", tc.name, got, tc.wantHit)
		}
		// TODO(panmari): Also check attributes of hitrecord.
	}
}

func BenchmarkNodeIntersection(b *testing.B) {
	p := NewPlane(vec3.UnitX, 1, materials.DiffuseDefault)
	s := NewSphere(vec3.Zero, 2, materials.DiffuseDefault)
	nInt := NewNode(s, p, INTERSECT)
	nSub := NewNode(s, p, SUBTRACT)
	nAdd := NewNode(s, p, ADD)
	r := &util.Ray{vec3.T{4, 0, 0}, vec3.T{-1, 0, 0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nInt.Intersect(r)
		nSub.Intersect(r)
		nAdd.Intersect(r)
	}
}
