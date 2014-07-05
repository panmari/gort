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

func Read(fileName string) (*Data) {
	data := Data{}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data.InsertLine(scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return &data
}

type Face struct {
	VertexIds [3]int
	TexCoordIds [3]int
	NormalIds [3]int
}

type Data struct{
	Vertices []vec3.T
	TexCoords []vec2.T
	Normals []vec3.T
	Faces []Face
	HasTexCoords bool
}

func (o *Data) InsertLine(line string) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	switch scanner.Text() {
		case "#":
			//comment, do nothing			
		case "v":
			o.Vertices = append(o.Vertices, parseVec3(scanner))
		case "vn":
			o.Normals = append(o.Normals, parseVec3(scanner))
		case "vt":
			o.TexCoords = append(o.TexCoords, parseVec2(scanner))
			o.HasTexCoords = true
		case "f":
			o.Faces = append(o.Faces, parseFace(scanner))
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
	//TODO: convert quadrangle Faces into triangle Faces
	for scanner.Scan() {
		face.VertexIds[counter], face.TexCoordIds[counter], face.NormalIds[counter] = parseFacePoint(scanner.Bytes())
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
	// minus one bc one based counting system in obj
	return id -1
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