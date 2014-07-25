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