package csg

import (
	"fmt"

	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type Shape interface {
	// GetIntervalBoundaries returns ALL intersections of a ray with Shape.
	GetIntervalBoundaries(r *util.Ray) ByT
	// BoundingBox returns a bounding box encompassing Shape.
	BoundingBox() *vec3.Box
}

type Solid struct {
	Shape
}

func (s *Solid) Intersect(r *util.Ray) *util.Hitrecord {
	ibs := s.GetIntervalBoundaries(r)
	for i := range ibs {
		hit := ibs[i].hit
		if hit != nil && hit.T > 0 {
			hit.Intersectable = s
			return hit
		}
	}
	return nil
}

func isBoundaryTypeStart(h *util.Hitrecord, r *util.Ray) bool {
	return vec3.Dot(&h.Normal, &r.Direction) < 0
}

type IntervalBoundary struct {
	t             float32
	isStart       bool
	belongsToLeft bool
	hit           *util.Hitrecord
}

func (i *IntervalBoundary) String() string {
	startEnd, belonging := "end", "right"
	if i.isStart {
		startEnd = "start"
	}
	if i.belongsToLeft {
		belonging = "left"
	}
	return fmt.Sprintf("At %f, %s, %s", i.t, startEnd, belonging)
}

// ByT is a slice of interval boundaries that is sortable by the T of their hitpoints
type ByT []IntervalBoundary

func (a ByT) Len() int           { return len(a) }
func (a ByT) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByT) Less(i, j int) bool { return a[i].t < a[j].t }
