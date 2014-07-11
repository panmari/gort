package intersectables

import (
	"fmt"
	"math"
	"util"
)

type Aggregate struct {
	Aggregator
}

type Aggregator interface {
	Size() int
	GetIntersectables() []util.Intersectable
}

func (a *Aggregate) Intersect(ray *util.Ray) *util.Hitrecord {
	var closestHit *util.Hitrecord
	closestT := float32(math.MaxFloat32)
	for _, i := range a.GetIntersectables() {
		if hit := i.Intersect(ray); hit != nil && hit.T < closestT && hit.T > 0 {
			closestHit = hit
			closestT = hit.T
		}
	}
	return closestHit
}

func (a *Aggregate) String() string {
	return fmt.Sprint(a.GetIntersectables())
}

func NewAggregate(a Aggregator) util.Intersectable {
	return &Aggregate{a}
}
