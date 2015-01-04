package intersectables

import (
	"fmt"
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
		min := vec3.MaxVal
		max := vec3.MinVal
		for j := 0; j < 3; j++ {
			t.vertices[j] = data.Vertices[face.VertexIds[j]]
			t.normals[j] = data.Normals[face.NormalIds[j]]
			min = vec3.Min(&min, t.vertices[j])
			max = vec3.Max(&max, t.vertices[j])
		}
		if data.HasTexCoords {
			for j := 0; j < 2; j++ {
				t.texCoords[j] = data.TexCoords[face.TexCoordIds[j]]
			}
		}
		// precompute edges used for intersection algorithm
		t.e1 = vec3.Sub(t.vertices[1], t.vertices[0])
		t.e2 = vec3.Sub(t.vertices[2], t.vertices[0])
		t.Box = vec3.Box{min, max}
		mesh.triangles[i] = &t
	}
	return &mesh
}

func NewMeshAggregate(data *obj.Data, m util.Material) util.Intersectable {
	return NewAggregate(NewMesh(data, m))
}