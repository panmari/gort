package intersectables

import (
	"fmt"
	"github.com/ungerik/go3d/vec3"
	"materials"
	"util"
	"util/obj"
)

type Mesh struct {
	Triangles []util.Intersectable
}

func (m *Mesh) Size() int {
	return len(m.Triangles)
}

func (m *Mesh) GetIntersectables() []util.Intersectable {
	return m.Triangles
}

func (m *Mesh) String() string {
	return fmt.Sprintf("Mesh: %v", m.Triangles)
}

func NewDiffuseMesh(data *obj.Data) *Mesh {
	return NewMesh(data, materials.MakeDiffuseMaterial(vec3.UnitXYZ))
}

func NewMesh(data *obj.Data, m util.Material) *Mesh {
	mesh := Mesh{make([]util.Intersectable, len(data.Faces))}
	for i, face := range data.Faces {
		t := MeshTriangle{Material: m}
		min := vec3.MaxVal
		max := vec3.MinVal
		for j := 0; j < 3; j++ {
			t.Vertices[j] = *data.Vertices[face.VertexIds[j]]
			t.Normals[j] = *data.Normals[face.NormalIds[j]]
			min = vec3.Min(&min, &t.Vertices[j])
			max = vec3.Max(&max, &t.Vertices[j])
		}
		if data.HasTexCoords {
			for j := 0; j < 2; j++ {
				t.TexCoords[j] = *data.TexCoords[face.TexCoordIds[j]]
			}
		}
		// TODO: Set flag if TexCoords are not available.
		// Precompute edges used for intersection algorithm.
		t.E1 = vec3.Sub(&t.Vertices[1], &t.Vertices[0])
		t.E2 = vec3.Sub(&t.Vertices[2], &t.Vertices[0])
		t.Box = vec3.Box{min, max}
		mesh.Triangles[i] = &t
	}
	return &mesh
}

func NewMeshAggregate(data *obj.Data, m util.Material) *Aggregate {
	return NewAggregate(NewMesh(data, m))
}
