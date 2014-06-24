package intersectables

import (
	"github.com/ungerik/go3d/mat4"
	"util"
)

type Instance struct {
}

func (i *Instance) Intersect(r *util.Ray) *util.Hitrecord {
	return new(util.Hitrecord)
}

func MakeInstance(intersectable *util.Intersectable, transform *mat4.T) {
	//TODO
}
