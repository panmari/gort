package obj

import (
    "testing"
    "bufio"
    "strings"
    "github.com/ungerik/go3d/vec3"
)

func TestParseVertex(t *testing.T) { 
	line := "v 3.4 1.2 4.1"
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	vector := parseVec3(scanner)
	
	expected := vec3.T{3.4, 1.2, 4.1}	
	if vector != expected{
		t.Errorf("Dat vörtex is sömthing else actually... %v", vector)
	}
}

func TestParseFacePoint(t *testing.T) { 
	line := "10/5/30"
	v_id, tc_id, n_id := parseFacePoint([]byte(line))
	if v_id != 10 {
		t.Error(v_id)
	}
	if tc_id != 5 {
		t.Error(tc_id)
	}
	if n_id != 30 {
		t.Error(n_id)
	}
}

func TestParseFacePointNoTc(t *testing.T) { 
	line := "10//30"
	v_id, tc_id, n_id := parseFacePoint([]byte(line))
	if v_id != 10 {
		t.Error(v_id)
	}
	if tc_id != -1 {
		t.Error(tc_id)
	}
	if n_id != 30 {
		t.Error(n_id)
	}
}

func TestParseFace(t *testing.T) { 
	o := ObjData{}
	line := "f 1/2/3 5/6/7 11/12/13"
	o.insertLine(line)
	f := o.faces[0]
	if f.vertexIds != [3]int{1, 5, 11} {
		t.Error(f.vertexIds)
	}
	if f.texCoordIds != [3]int{2, 6, 12} {
		t.Error(f.texCoordIds)
	}
	if f.normalIds != [3]int{3, 7, 13} {
		t.Error(f.normalIds)
	}
}

func BenchmarkParseLine(b *testing.B) {
	line := "v 3.4 1.2 4.1"
	for i := 0; i < b.N; i++ {
		o := ObjData{}
		o.insertLine(line)
	}
}

func BenchmarkTeapot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read("obj/teapot.obj")
	}
}