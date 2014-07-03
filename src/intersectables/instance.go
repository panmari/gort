package intersectables

import (
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
	"util"
	"materials"
)

type Instance struct {
	t 				mat4.T 
	tinverse		mat4.T
	tinverseT		mat4.T
	material 		util.Material
	intersectable	util.Intersectable
}

// Transform the given ray into the coordinate frame of the instance and returns the resulting intersection.
func (i *Instance) Intersect(r *util.Ray) *util.Hitrecord {
	//transform ray into coordinate system of instance
	rTransformed := util.Ray{r.Origin, r.Direction}
	i.tinverse.MulVec3(&rTransformed.Origin, 1)
	i.tinverse.MulVec3(&rTransformed.Direction, 0)
	h := i.intersectable.Intersect(&rTransformed)
	if h == nil {
		return nil
	}
	
	//transform back
	i.t.MulVec3(&h.Position, 1)
	i.t.MulVec3(&h.W_in, 0)
	h.W_in.Normalize()
	
	//use transpose of inverse for normal
	i.tinverseT.MulVec3(&h.Normal, 0)
	//normalize again, bc may contain scaling
	h.Normal.Normalize()
	return h
}

func NewDiffuseInstance(intersectable util.Intersectable, transformation mat4.T) util.Intersectable {
	i := new(Instance)
	i.t = transformation
	i.tinverse = transformation
	i.tinverse.Invert()
	i.tinverseT = i.tinverse
	i.tinverseT.Transpose()
	i.material = materials.MakeDiffuseMaterial(vec3.T{1,1,1})
	i.intersectable = intersectable
	return i
}
