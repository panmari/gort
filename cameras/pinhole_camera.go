package cameras

import (
	"github.com/barnex/fmath"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
)

type PinholeCamera struct {
	M   mat4.T
	Eye vec3.T
}

func MakePinholeCamera(eye *vec3.T, lookat *vec3.T, up *vec3.T,
	fov float32, aspect float32,
	width int, height int) *PinholeCamera {
	cam := new(PinholeCamera)
	cam.Eye = *eye
	w := vec3.Sub(eye, lookat)
	w.Normalize()
	u := vec3.Cross(up, &w)
	u.Normalize()
	v := vec3.Cross(&w, &u)

	var camMatrix mat4.T
	camMatrix[0][0] = u[0]
	camMatrix[0][1] = u[1]
	camMatrix[0][2] = u[2]
	camMatrix[1][0] = v[0]
	camMatrix[1][1] = v[1]
	camMatrix[1][2] = v[2]
	camMatrix[2][0] = w[0]
	camMatrix[2][1] = w[1]
	camMatrix[2][2] = w[2]
	camMatrix.SetTranslation(eye)

	var vp mat4.T
	t := fmath.Tan(util.ToRadians(fov) / 2)
	r := aspect * t
	vp[0][0] = 2 * r / float32(width)
	vp[2][0] = r
	vp[1][1] = 2 * t / float32(height)
	vp[2][1] = t
	vp[2][2] = 1
	vp[3][3] = 1
	cam.M.AssignMul(&camMatrix, &vp)
	return cam
}

func (c *PinholeCamera) MakeWorldSpaceRay(i, j int, samples [2]float32) *util.Ray {
	d := vec4.T{float32(i) + samples[0], float32(j) + samples[1], -1.0, 1.0}

	// Transform it back to world coordinates
	dTransformed := c.M.MulVec4(&d)
	// Make ray consisting of origin and direction in world coordinates
	dir := vec3.T{dTransformed[0], dTransformed[1], dTransformed[2]}
	dir.Sub(&c.Eye) //.Normalize()
	return &util.Ray{c.Eye, dir}
}
