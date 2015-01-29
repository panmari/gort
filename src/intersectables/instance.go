package intersectables

import (
	"fmt"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
	"materials"
	"util"
)

type Instance struct {
	T             mat4.T
	Tinverse      mat4.T
	TinverseTranspose     mat4.T
	Material      util.Material
	Intersectable util.Intersectable
	Box           vec3.Box
}

// Transforms the given ray into the coordinate frame of this instance and returns the resulting intersection.
func (i *Instance) Intersect(r *util.Ray) *util.Hitrecord {
	//transform ray into coordinate system of instance
	rTransformed := util.Ray{r.Origin, r.Direction}
	i.Tinverse.TransformVec3W(&rTransformed.Origin, 1)
	i.Tinverse.TransformVec3W(&rTransformed.Direction, 0)
	h := i.Intersectable.Intersect(&rTransformed)
	if h == nil {
		return nil
	}
	//transform back
	i.T.TransformVec3W(&h.Position, 1)
	i.T.TransformVec3W(&h.W_in, 0)
	h.W_in.Normalize()

	//use transpose of inverse for normal
	i.TinverseTranspose.TransformVec3W(&h.Normal, 0)
	//normalize again, bc may contain scaling
	h.Normal.Normalize()
	h.Material = i.Material
	return h
}

func (i *Instance) BoundingBox() *vec3.Box {
	return &i.Box
}

func (i *Instance) String() string {
	return fmt.Sprintf("Instance around %v", i.Intersectable)
}

func NewInstance(intersectable util.Intersectable, transformation mat4.T, m util.Material) *Instance {
	i := new(Instance)
	i.T = transformation
	i.Tinverse = transformation
	i.Tinverse.Invert()
	i.TinverseTranspose = i.Tinverse
	i.TinverseTranspose.Transpose()
	i.Material = m
	i.Intersectable = intersectable

	// Transform bounding box.
	bb := intersectable.BoundingBox()
	minInstance := bb.Min
	i.T.TransformVec3W(&minInstance, 1)
	maxInstance := bb.Max
	i.T.TransformVec3W(&maxInstance, 1)
	i.Box = vec3.Box{minInstance, maxInstance}
	return i
}

// Same as above with default material.
func NewDiffuseInstance(intersectable util.Intersectable, transformation mat4.T) *Instance {
	return NewInstance(intersectable, transformation, materials.DiffuseDefault)
}
