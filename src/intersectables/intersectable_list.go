package intersectables

import (
	"util"
	"math"
)

type IntersectableList struct {
	intersectables []util.Intersectable
}

func (list *IntersectableList) Add(intersectable util.Intersectable) {
	list.intersectables = append(list.intersectables, intersectable)
}

func (list *IntersectableList) Intersect(ray *util.Ray) (*util.Hitrecord, bool) {
	closestHit := new(util.Hitrecord)
	closestHit.T = math.MaxFloat32
	hadHit := false
	for _, i := range list.intersectables {
		if hit, doesHit := i.Intersect(ray); doesHit && hit.T < closestHit.T && hit.T > 0 {
			closestHit = hit
			hadHit = true
		}
	}
	return closestHit, hadHit
}

func MakeIntersectableList(initialSize int) *IntersectableList {
	i := new(IntersectableList)
	i.intersectables = make([]util.Intersectable, 0, initialSize)
	return i
}