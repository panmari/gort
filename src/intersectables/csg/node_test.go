package csg

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
	"util"
	)

func TestNodeAdd(t *testing.T) {
	
	p := NewDiffusePlane(vec3.UnitX, 1)
	s := NewDiffuseSphere(vec3.Zero, 2)	
	n := NewNode(p, s, ADD)
	
	parallelRay := util.Ray{vec3.T{0,3,0}, vec3.UnitY}	
	if hit := n.Intersect(&parallelRay); hit != nil {
		t.Errorf("Parallel ray hit plane: %v", hit)
	}

	parallelRaySphereHit := util.Ray{vec3.T{0,-3,0}, vec3.UnitY}
	hitPar := n.Intersect(&parallelRaySphereHit);
	if  hitPar == nil {
		t.Errorf("Does not hit added sphere: %v", hitPar)
	}
	if hitPar.T != 1 {
		t.Errorf("Strange T for hit: %f", hitPar.T)
	}
	expected := vec3.T{0, -2, 0};
	if hitPar.Position != expected {
		t.Errorf("Strange hit Position for hit: %v", hitPar.Position)
	}
	
	pointingAwayRay := util.Ray{vec3.Zero, vec3.UnitX}
	if hit := n.Intersect(&pointingAwayRay);  hit == nil {
		t.Errorf("Ray didn't hit inside of sphere: %v", hit)
	}

	pointingTowardsRay := util.Ray{vec3.T{4,0,0}, vec3.T{-1, 0, 0}}
	hit := n.Intersect(&pointingTowardsRay)
	if hit == nil {
		t.Errorf("Ray pointing towards plane does not hit plane: %v", hit)
	}
	if hit.T != 2 {
		t.Errorf("Ray doesnt hit sphere when shot from inside, T: %f", hit.T)
	}
	
	shouldHitPlane := util.Ray{vec3.T{10, 10, 10}, vec3.T{-1,0,0}}
	hitPlane := n.Intersect(&shouldHitPlane);
	if hitPlane == nil {
		t.Errorf("ray does not hit plane: %v", hit)
	}
	if hitPlane.T != 11 {
		t.Errorf("ray does not hit plane at correct T: %f", hitPlane.T)
	}
	expected = vec3.T{-1, 10, 10}
	if hitPlane.Position != expected {
		t.Errorf("ray does not hit plane at correct Position: %v", hitPlane.Position)
	}
}

func TestNodeSubtract(t *testing.T) {
	p := NewDiffusePlane(vec3.UnitX, 1)
	s := NewDiffuseSphere(vec3.Zero, 2)	
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
	p := NewDiffusePlane(vec3.UnitX, 1)
	s := NewDiffuseSphere(vec3.Zero, 2)	
	n := NewNode(s, p, INTERSECT)
	r := util.Ray{vec3.T{4, 0, 0}, vec3.T{-1, 0, 0}}
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

}