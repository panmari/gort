package intersectables

import (
	"util"
	"math"
)

type IntersectableList struct {
	intersectables []util.Intersectable
}

func (list IntersectableList) Add(intersectable util.Intersectable) {
	list.intersectables = append(list.intersectables, intersectable)
}

func (list IntersectableList) Intersect(ray *util.Ray) *util.Hitrecord {
	closestHit := new(util.Hitrecord)
	closestHit.T = math.MaxFloat32
	for _, i := range list.intersectables {
		if hit := i.Intersect(ray); hit.DoesHit() && hit.T < closestHit.T {
			closestHit = hit
		}
	}
	return closestHit
}

func MakeIntersectableList(initialSize int) *IntersectableList {
	i := new(IntersectableList)
	i.intersectables = make([]util.Intersectable, initialSize)
	return i
}