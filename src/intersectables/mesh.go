package intersectables

import (
	"fmt"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
	"materials"
	"util"
	"util/obj"
)

type Mesh struct {
	triangles []util.Intersectable
}

func (m *Mesh) Size() int {
	return len(m.triangles)
}

func (m *Mesh) GetIntersectables() []util.Intersectable {
	return m.triangles
}

func (m *Mesh) String() string {
	return fmt.Sprintf("Mesh: %v", m.triangles)
}

func NewDiffuseMesh(data *obj.Data) *Mesh {
	return NewMesh(data, materials.MakeDiffuseMaterial(vec3.UnitXYZ))
}

func NewMesh(data *obj.Data, m util.Material) *Mesh {
	mesh := Mesh{make([]util.Intersectable, len(data.Faces))}
	for i, face := range data.Faces {
		t := MeshTriangle{material: m}
		for j := 0; j < 3; j++ {
			t.vertices[j] = data.Vertices[face.VertexIds[j]]
			t.normals[j] = data.Normals[face.NormalIds[j]]
			if data.HasTexCoords {
				t.texCoords[j] = data.TexCoords[face.TexCoordIds[j]]
			}
		}
		// precompute matrix used for cramers rule
		t.e1 = vec3.Sub(t.vertices[1], t.vertices[0])
		t.e2 = vec3.Sub(t.vertices[2], t.vertices[0])
		mesh.triangles[i] = &t
	}
	return &mesh
}

func NewMeshAggregate(data *obj.Data, m util.Material) util.Intersectable {
	return NewAggregate(NewMesh(data, m))
}

type MeshTriangle struct {
	vertices  [3]*vec3.T
	normals   [3]*vec3.T
	texCoords [2]*vec2.T
	e1, e2    vec3.T
	material  util.Material
}

const (
	EPSILON = 1e-6
)

// Using Möller-Trumbore intersectin algorithm from
// http://en.wikipedia.org/wiki/M%C3%B6ller%E2%80%93Trumbore_intersection_algorithm
// to compute intersections with mesh triangle.
func (tr *MeshTriangle) Intersect(r *util.Ray) *util.Hitrecord {
	parameter := vec3.Cross(&r.Direction, &tr.e2)
	det := vec3.Dot(&tr.e1, &parameter)
	if det > -EPSILON && det < EPSILON {
		return nil
	}
	inv_det := 1 / det

	dist := vec3.Sub(&r.Origin, tr.vertices[0])
	u := vec3.Dot(&dist, &parameter) * inv_det
	if u < 0 || u > 1 {
		return nil
	}
	vParameter := vec3.Cross(&dist, &tr.e1)
	v := vec3.Dot(&r.Direction, &vParameter) * inv_det
	if v < 0 || u+v > 1 {
		return nil
	}

	t := vec3.Dot(&tr.e2, &vParameter) * inv_det

	if t > EPSILON {
		h := new(util.Hitrecord)
		h.T = t
		h.Position = r.PointAt(h.T)
		h.Normal = *tr.makeNormal(1-u-v, u, v)
		h.W_in = r.Direction
		h.W_in.Normalize().Scale(-1)
		//TODO: texture coordinates
		h.Material = tr.material
		h.Intersectable = tr
		return h
	}
	return nil
}

func (t *MeshTriangle) makeNormal(alpha, beta, gamma float32) *vec3.T {
	var normal vec3.T
	n0 := t.normals[0].Scaled(alpha)
	normal.Add(&n0)
	n1 := t.normals[1].Scaled(beta)
	normal.Add(&n1)
	n2 := t.normals[2].Scaled(gamma)
	normal.Add(&n2)
	// this should not be needed, but most meshes suck...
	normal.Normalize()
	return &normal
}

func (t *MeshTriangle) String() string {
	return fmt.Sprintf("v: %v, \nn: %v, \ntc: %v", t.vertices, t.normals, t.texCoords)
}
