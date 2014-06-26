package intersectables

import (
	"math"
	"util"
)

type IntersectableList struct {
	intersectables []util.Intersectable
}

func (list *IntersectableList) Add(intersectable util.Intersectable) {
	list.intersectables = append(list.intersectables, intersectable)
}

func (list *IntersectableList) Intersect(ray *util.Ray) (*util.Hitrecord, bool) {
	var closestHit *util.Hitrecord
	closestT := float32(math.MaxFloat32)
	hadHit := false
	for _, i := range list.intersectables {
		if hit, doesHit := i.Intersect(ray); doesHit && hit.T < closestT && hit.T > 0 {
			closestHit = hit
			closestT = hit.T
			hadHit = true
		}
	}
	return closestHit, hadHit
}

func MakeIntersectableList(initialSize int) IntersectableList {
	i := IntersectableList{}
	i.intersectables = make([]util.Intersectable, 0, initialSize)
	return i
}
