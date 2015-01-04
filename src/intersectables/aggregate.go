package intersectables

import (
	"fmt"
	"math"
	"github.com/ungerik/go3d/vec3"
	"util"
)

type Aggregate struct {
	Aggregator
	box vec3.Box
}

type Aggregator interface {
	Size() int
	GetIntersectables() []util.Intersectable
}

func (a *Aggregate) Intersect(ray *util.Ray) *util.Hitrecord {
	var closestHit *util.Hitrecord
	closestT := float32(math.MaxFloat32)
	intersectables := a.GetIntersectables()
	for i := range intersectables {
		if hit := intersectables[i].Intersect(ray); hit != nil && hit.T < closestT && hit.T > 0 {
			closestHit = hit
			closestT = hit.T
		}
	}
	return closestHit
}

func (a *Aggregate) BoundingBox() *vec3.Box {
	return &a.box;
}

func (a *Aggregate) String() string {
	return fmt.Sprint(a.GetIntersectables())
}

func NewAggregate(a Aggregator) util.Intersectable {
	bb := vec3.Box{}
	intersectables := a.GetIntersectables()
	for i := range intersectables {
		bb.Join(intersectables[i].BoundingBox())
	}
	return &Aggregate{a, bb}
}
