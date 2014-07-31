package obj

import (
	"bufio"
	"github.com/ungerik/go3d/vec3"
	"strings"
	"testing"
)

func TestParseVertex(t *testing.T) {
	line := "v 3.4 1.2 4.1"
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	vector, _ := parseVec3(scanner)

	expected := vec3.T{3.4, 1.2, 4.1}
	if *vector != expected {
		t.Errorf("Dat vörtex is sömthing else actually... %v", vector)
	}
}

func TestParseFacePoint(t *testing.T) {
	line := "10/5/30"
	v_id, tc_id, n_id := parseFacePoint([]byte(line))
	if v_id != 9 {
		t.Error(v_id)
	}
	if tc_id != 4 {
		t.Error(tc_id)
	}
	if n_id != 29 {
		t.Error(n_id)
	}
}

func TestParceFacePointNoNormal(t *testing.T) {
	line := "32/32" 
	v_id, tc_id, n_id := parseFacePoint([]byte(line))
	if v_id != 31 {
		t.Error(v_id)
	}
	if tc_id != 31 {
		t.Error(tc_id)
	}
	if n_id != -1 {
		t.Error(n_id)
	}
}

func TestParseFacePointNoTc(t *testing.T) {
	line := "10//30"
	v_id, tc_id, n_id := parseFacePoint([]byte(line))
	if v_id != 9 {
		t.Error(v_id)
	}
	if tc_id != -1 {
		t.Error(tc_id)
	}
	if n_id != 29 {
		t.Error(n_id)
	}
}

func TestParseFace(t *testing.T) {
	o := Data{}
	line := "f 1/2/3 5/6/7 11/12/13"
	o.InsertLine(line)
	if len(o.Faces) != 1 {
		t.Fatal("Triangle should have 1 face, instead:", len(o.Faces))
	}
	f := o.Faces[0]
	if f.VertexIds != [3]int{0, 4, 10} {
		t.Error(f.VertexIds)
	}
	if f.TexCoordIds != [3]int{1, 5, 11} {
		t.Error(f.TexCoordIds)
	}
	if f.NormalIds != [3]int{2, 6, 12} {
		t.Error(f.NormalIds)
	}
}

func TestParseQuadrangleFace(t *testing.T) {
	o := Data{}
	line := "f 1/2/3 5/6/7 11/12/13 21/22/23"
	o.InsertLine(line)
	if len(o.Faces) != 2 {
		t.Fatal("Quadrangle should have 2 faces, instead:", len(o.Faces))
	}
	f := o.Faces[0]
	if f.VertexIds != [3]int{0, 4, 10} {
		t.Error(f.VertexIds)
	}
	if f.TexCoordIds != [3]int{1, 5, 11} {
		t.Error(f.TexCoordIds)
	}
	if f.NormalIds != [3]int{2, 6, 12} {
		t.Error(f.NormalIds)
	}
	f2 := o.Faces[1]
	if f2.VertexIds != [3]int{0, 10, 20} {
		t.Error(f2.VertexIds)
	}
	if f2.TexCoordIds != [3]int{1, 11, 21} {
		t.Error(f2.TexCoordIds)
	}
	if f2.NormalIds != [3]int{2, 12, 22} {
		t.Error(f2.NormalIds)
	}
}

const EPSILON = 1e-6

func TestNormalInterpolation(t *testing.T) {
	o := Data{}
	o.InsertLine("f 1/1 2/2 3/3 ")
	o.InsertLine("v 170.897228 77.953877 -1.212034")
	o.InsertLine("v 170.893054 198.184678 -0.928379")
	o.InsertLine("v 173.441402 79.129056 -1.212072")
	if len(o.Normals) != 0 {
		t.Fatalf("Expected 0 normals before interpolating, but had", len(o.Normals))
	}
	o.interpolateNormals()
	if len(o.Normals) != 3 {
		t.Fatalf("Expected 3 normals, but had", len(o.Normals))
	}
	expected := &vec3.T{-0.0011047, 0.0023592, -0.9999966}
	if n := o.Normals[0]; vec3.Distance(n, expected) > EPSILON {
		t.Error(n)
	}
	if n := o.Normals[1]; vec3.Distance(n, expected) > EPSILON {
		t.Error(n)
	}
	if n := o.Normals[2]; vec3.Distance(n, expected) > EPSILON {
		t.Error(n)
	}
}

func BenchmarkParseLine(b *testing.B) {
	line := "v 3.4 1.2 4.1"
	for i := 0; i < b.N; i++ {
		o := Data{}
		o.InsertLine(line)
	}
}

func BenchmarkTeapot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read("../../../obj/teapot.obj", 1)
	}
}

func BenchmarkHeart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read("../../../obj/Heart.obj", 1)
	}
}

func BenchmarkDragon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read("../../../obj/dragon.obj", 1)
	}
}