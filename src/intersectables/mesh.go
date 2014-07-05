package intersectables

import (
	"util"
	"util/obj"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/mat3"
)

type Mesh struct {
	triangles []MeshTriangle
}

//TODO: iterate over all triangles
func (m *Mesh) Intersect(r *util.Ray) *util.Hitrecord {
	return nil
}

func NewMesh(data *obj.Data) {
	mesh := Mesh{make([]MeshTriangle, len(data.Faces))}
	for i, face := range data.Faces {
		for j := 0; j < 3; j++ {
			mesh.triangles[i].vertices[j] = &data.Vertices[face.VertexIds[j]]
			mesh.triangles[i].texCoords[j] = &data.TexCoords[face.TexCoordIds[j]]
			mesh.triangles[i].normals[j] = &data.Normals[face.NormalIds[j]]
		}
	}
}


type MeshTriangle struct {
	vertices [3]*vec3.T
	normals [3]*vec3.T
	texCoords [2]*vec2.T
}

func (t *MeshTriangle) Intersect(r *util.Ray) *util.Hitrecord {
	col0 := vec3.Sub(t.vertices[0], t.vertices[1])
	col1 := vec3.Sub(t.vertices[0], t.vertices[2])
	m := mat3.T{col0, col1, r.Direction}
	b := vec3.Sub(t.vertices[0], &r.Origin)
	
	betaGammaT := getBetaGammaTCramer(&m, b)
	if isInside(betaGammaT) {
		//TODO: assemble hitrecord
		return new(util.Hitrecord)
	}
	return nil
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

func getBetaGammaTCramer(m *mat3.T, b vec3.T) *vec3.T {
	detA := m.Determinant()
	m0 := mat3.T{b, m[1], m[2]}
	detA0 := m0.Determinant()
	m1 := mat3.T{m[0], b, m[2]}
	detA1 := m1.Determinant()
	m2 := mat3.T{m[0], m[1], b}
	detA2 := m2.Determinant()
	// alpha, beta gamma in one vector
	return &vec3.T{detA0/detA, detA1/detA, detA2/detA}
}