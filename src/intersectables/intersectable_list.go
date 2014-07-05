package intersectables

import (
	"util"
)

type IntersectableList struct {
	intersectables []util.Intersectable
}

func (list *IntersectableList) Add(intersectables ...util.Intersectable) {
	list.intersectables = append(list.intersectables, intersectables...)
}

func (list *IntersectableList) GetIntersectables() []util.Intersectable {
	return list.intersectables
}

func (list *IntersectableList) Size() int {
	return len(list.intersectables)
}

func NewIntersectableList(initialSize int) *IntersectableList {
	i := new(IntersectableList)
	i.intersectables = make([]util.Intersectable, 0, initialSize)
	return i
}
