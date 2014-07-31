package intersectables

import (
	"fmt"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
	"materials"
	"util"
)

type Instance struct {
	t             mat4.T
	tinverse      mat4.T
	tinverseT     mat4.T
	Material      util.Material
	intersectable util.Intersectable
}

// Transform the given ray into the coordinate frame of the instance and returns the resulting intersection.
func (i *Instance) Intersect(r *util.Ray) *util.Hitrecord {
	//transform ray into coordinate system of instance
	rTransformed := util.Ray{r.Origin, r.Direction}
	i.tinverse.TransformVec3(&rTransformed.Origin, 1)
	i.tinverse.TransformVec3(&rTransformed.Direction, 0)
	h := i.intersectable.Intersect(&rTransformed)
	if h == nil {
		return nil
	}
	//transform back
	i.t.TransformVec3(&h.Position, 1)
	i.t.TransformVec3(&h.W_in, 0)
	h.W_in.Normalize()

	//use transpose of inverse for normal
	i.tinverseT.TransformVec3(&h.Normal, 0)
	//normalize again, bc may contain scaling
	h.Normal.Normalize()
	h.Material = i.Material
	return h
}

func (i *Instance) String() string {
	return fmt.Sprintf("Instance around %v", i.intersectable)
}

func NewDiffuseInstance(intersectable util.Intersectable, transformation mat4.T) *Instance {
	i := new(Instance)
	i.t = transformation
	i.tinverse = transformation
	i.tinverse.Invert()
	i.tinverseT = i.tinverse
	i.tinverseT.Transpose()
	i.Material = materials.MakeDiffuseMaterial(vec3.T{1, 1, 1})
	i.intersectable = intersectable
	return i
}
