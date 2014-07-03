package intersectables

import (
	"math"
	"util"
)

type IntersectableList struct {
	intersectables []util.Intersectable
}

func (list *IntersectableList) Add(intersectables ...util.Intersectable) {
	list.intersectables = append(list.intersectables, intersectables...)
}

func (list *IntersectableList) Intersect(ray *util.Ray) (*util.Hitrecord) {
	var closestHit *util.Hitrecord
	closestT := float32(math.MaxFloat32)
	for _, i := range list.intersectables {
		if hit := i.Intersect(ray); hit != nil && hit.T < closestT && hit.T > 0 {
			closestHit = hit
			closestT = hit.T
		}
	}
	return closestHit
}

func NewIntersectableList(initialSize int) *IntersectableList {
	i := new(IntersectableList)
	i.intersectables = make([]util.Intersectable, 0, initialSize)
	return i
}
