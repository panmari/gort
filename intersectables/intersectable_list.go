package intersectables

import (
	"github.com/panmari/gort/util"
)

type IntersectableList struct {
	Intersectables []util.Intersectable
}

func (list *IntersectableList) Add(intersectables ...util.Intersectable) {
	list.Intersectables = append(list.Intersectables, intersectables...)
}

func (list *IntersectableList) GetIntersectables() []util.Intersectable {
	return list.Intersectables
}

func (list *IntersectableList) Size() int {
	return len(list.Intersectables)
}

func NewIntersectableList(initialSize int) *IntersectableList {
	i := new(IntersectableList)
	i.Intersectables = make([]util.Intersectable, 0, initialSize)
	return i
}
