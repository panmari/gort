package cameras

import (
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
	"util"
)

type FixedCamera struct {
	M   mat4.T
	Eye vec3.T
}

func NewFixedCamera(width int, height int) *FixedCamera {
	cam := new(FixedCamera)
	// Fixed eye position in world coordinates
	cam.Eye = vec3.T{0, 0, 3.0}

	// Fixed camera to world transform, just a translation
	camMatrix := mat4.Ident
	camMatrix.SetTranslation(&cam.Eye)

	// Fixed projection matrix, the viewing frustum defined here goes through
	// the points [-1,-1,-1], [1,-1,-1], [-1,1,-1],[1,1,-1] in camera coordinates.
	p := mat4.Ident
	near, far := float32(1), float32(10)

	p[2][2] = -(far + near) / (far - near)
	p[3][2] = -(2 * far * near) / (far - near)
	p[2][3] = -1
	p[3][3] = 0
	p.Invert()

	// Make viewport matrix given image resolution
	v := mat4.Ident
	v[0][0] = float32(width) / 2
	v[3][0] = float32(width) / 2
	v[1][1] = float32(height) / 2
	v[3][1] = float32(height) / 2
	v[2][2] = 1
	v[3][2] = 0
	v.Invert()
	//TODO: invert!
	vp := new(mat4.T)
	vp.AssignMul(&v, &p)
	cam.M.AssignMul(&camMatrix, vp)
	return cam
}

func (c *FixedCamera) MakeWorldSpaceRay(i, j int, samples [2]float32) *util.Ray {
	d := vec4.T{float32(i) + samples[0], float32(j) + samples[1], -1.0, 1.0}

	// Transform it back to world coordinates
	dTransformed := c.M.MulVec4(&d)
	// Make ray consisting of origin and direction in world coordinates
	dir := vec3.T{dTransformed[0], dTransformed[1], dTransformed[2]}
	dir.Sub(&c.Eye) //.Normalize()
	return &util.Ray{c.Eye, dir}
}
