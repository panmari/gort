package obj

import (
	"os"
	"bufio"
	"strconv"
	"log"
	"strings"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec2"
	"bytes"
)

func Read(fileName string) ([]float32, error){
	myObjData := ObjData{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		myObjData.insertLine(scanner.Text())
	}
	return nil, scanner.Err()
}

type Face struct {
	vertexIds [3]int
	texCoordIds [3]int
	normalIds [3]int
}

type ObjData struct{
	vertices []vec3.T
	texCoords []vec2.T
	normals []vec3.T
	faces []Face
}

func (o *ObjData) insertLine(line string) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	switch scanner.Text() {
		case "#":
			//comment, do nothing			
		case "v":
			o.vertices = append(o.vertices, parseVec3(scanner))
		case "vn":
			o.normals = append(o.normals, parseVec3(scanner))
		case "vt":
			o.texCoords = append(o.texCoords, parseVec2(scanner))
		case "f":
			o.faces = append(o.faces, parseFace(scanner))
		default:
			log.Printf("Can not parse %s", line)
		
	}
}


func parseVec3(scanner *bufio.Scanner) vec3.T {
	var vector vec3.T
	counter := 0
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			log.Printf("Could not turn into float: %v", scanner.Text())
		}
		vector[counter] = float32(f)
		counter++
	}
	return vector
}

func parseVec2(scanner *bufio.Scanner) vec2.T {
	var vector vec2.T
	counter := 0
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			log.Printf("Could not turn into float: %v", scanner.Text())
		}
		vector[counter] = float32(f)
		counter++
	}
	return vector
}

func parseFace(scanner *bufio.Scanner) Face {
	var face Face
	counter := 0
	//TODO: convert quadrangle faces into triangle faces
	for scanner.Scan() {
		face.vertexIds[counter], face.texCoordIds[counter], face.normalIds[counter] = parseFacePoint(scanner.Bytes())
		counter++
	}
	return face
}

//split according to format: "vertex/texcoord/normal"
func parseFacePoint(data []byte) (int, int, int) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(ScanSlashes)
	
	vertex_id := parseId(scanner)
	texCoord_id := parseId(scanner)
	normal_id := parseId(scanner)
	
	return vertex_id, texCoord_id, normal_id
}

func parseId(scanner *bufio.Scanner) int {
	scanner.Scan()
	t := scanner.Text()
	if t == "" {
		return -1
	}
	id, err := strconv.Atoi(t)
	if err != nil {
		log.Print("Failed to parse %v", err)
	}
	return id
}

func ScanSlashes(data []byte, atEOF bool) (advance int, token[]byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i:= bytes.IndexByte(data, '/'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}