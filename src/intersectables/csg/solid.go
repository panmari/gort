package csg

import (
	"util"
)

type Shape interface {
	GetIntervalBoundaries(r *util.Ray) *ByT
}

type Solid struct {
	shape Shape
	}

func (s *Solid) Intersect(r *util.Ray) (*util.Hitrecord, bool) {
	for _, ib := range *s.shape.GetIntervalBoundaries(r) {
		hit := ib.hit
		if hit != nil && hit.T > 0 {
			hit.Intersectable = s
			return hit, true
		} 
	}
	return nil, false
}

type IntervalBoundary struct {
	t              float32
	isStart        bool
	belongsToLeft  bool
	hit            *util.Hitrecord
}

// A slice of interval boundaries that is sortable by the T of their hitpoints
type ByT []IntervalBoundary

func (a ByT) Len() int				{return len(a) }
func (a ByT) Swap(i, j int)			{ a[i], a[j] = a[j], a[i] }
func (a ByT) Less(i, j int)	bool	{ return a[i].t < a[j].t }