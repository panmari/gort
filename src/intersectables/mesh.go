package intersectables

import (
	"fmt"
	"util"
	"util/obj"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/mat3"
	"materials"
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
		col0 := vec3.Sub(t.vertices[0], t.vertices[1])
		col1 := vec3.Sub(t.vertices[0], t.vertices[2])
		t.cramerMat = mat3.T{col0, col1, vec3.Zero}
		mesh.triangles[i] = &t
	}
	return &mesh
}

func NewMeshAggregate(data *obj.Data, m util.Material) util.Intersectable {
	return NewAggregate(NewMesh(data, m))
}

type MeshTriangle struct {
	cramerMat  mat3.T
	vertices   [3]*vec3.T
	normals    [3]*vec3.T
	texCoords  [2]*vec2.T
	material   util.Material
}

func (t *MeshTriangle) Intersect(r *util.Ray) *util.Hitrecord {
	m := t.cramerMat
	m[2] = r.Direction
	b := vec3.Sub(t.vertices[0], &r.Origin)
	
	betaGammaT := t.getBetaGammaTCramer(&m, &b)
	if isInside(betaGammaT) {
		h := new(util.Hitrecord)
		h.T = betaGammaT[2]
		h.Position = r.PointAt(h.T)
		h.Normal = *t.makeNormal(betaGammaT)
		h.W_in = r.Direction
		h.W_in.Normalize().Scale(-1)
		//TODO: texture coordinates
		h.Material = t.material
		h.Intersectable = t
		return h
	}
	return nil
}
func (t *MeshTriangle) makeNormal(betaGammaT *vec3.T) *vec3.T {
	var normal vec3.T
	n0 := t.normals[0].Scaled(1 - betaGammaT[0] - betaGammaT[1])
	normal.Add(&n0)
	n1 := t.normals[1].Scaled(betaGammaT[0])
	normal.Add(&n1)
	n2 := t.normals[2].Scaled(betaGammaT[1])
	normal.Add(&n2)
	// this should not be needed, but most meshes suck...
	normal.Normalize()
	return &normal
}
func isInside(betaGammaT *vec3.T) bool {
	if 	(betaGammaT[0] <= 0 ||
		betaGammaT[0] >= 1 ||
		betaGammaT[1] <= 0 ||
		betaGammaT[1] >= 1) {
		return false
	}
	f := betaGammaT[0] + betaGammaT[1]
	return f > 0 && f < 1	
}

// Solves the linear system m*a = b via cramer's rule and returns a
func (t *MeshTriangle) getBetaGammaTCramer(m *mat3.T, b *vec3.T) *vec3.T {
	detA := m.Determinant()
	// set one column to b and get determinant, do for all three columns
	m[0] = *b
	detA0 := m.Determinant()
	m[0] = t.cramerMat[0]
	m[1] = *b
	detA1 := m.Determinant()
	m[1] = t.cramerMat[1]
	m[2] = *b
	detA2 := m.Determinant()
	
	// alpha, beta, gamma in one vector
	// reuse b to return result
	b[0] = detA0/detA
	b[1] = detA1/detA
	b[2] = detA2/detA
	return b
}

func (t *MeshTriangle) String() string {
	return fmt.Sprintf("v: %v, \nn: %v, \ntc: %v", t.vertices, t.normals, t.texCoords)
}