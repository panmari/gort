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

