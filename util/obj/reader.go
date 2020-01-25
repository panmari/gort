package obj

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
)

// Read returns the given obj file as *obj.Data, centering it at origin and scaling it by the given amount.
func Read(fileName string, scale float32) *Data {
	if scale <= 0 {
		log.Fatalf("Invalid scale factor %f for %s: must be >= 0", scale, fileName)
	}
	data := Data{min: vec3.MaxVal, max: vec3.MinVal}
	file, err := os.Open(fileName)
	if err != nil {
		wd, _ := os.Getwd()
		log.Fatalf("Could not find obj file in %s/%s", wd, fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data.InsertLine(scanner.Bytes())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	trans := vec3.Add(&data.max, &data.min)
	trans.Scale(float32(-1) / 2)

	normalizeScale := float32(math.MaxFloat32)
	for i := 0; i < 3; i++ {
		s := 2 / (data.max[i] - data.min[i])
		if s < normalizeScale {
			normalizeScale = s
		}
	}
	usedScale := scale * normalizeScale
	for _, v := range data.Vertices {
		v.Add(&trans).Scale(usedScale)
	}

	// Manually compute normals with cross product if not available in obj
	if len(data.Normals) == 0 {
		data.interpolateNormals()
	}
	return &data
}

type Face struct {
	VertexIds   [3]int
	TexCoordIds [3]int
	NormalIds   [3]int
}

type Data struct {
	Vertices     []*vec3.T
	TexCoords    []*vec2.T
	Normals      []*vec3.T
	Faces        []Face
	HasTexCoords bool
	min          vec3.T
	max          vec3.T
}

func (o *Data) InsertLine(line []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	switch scanner.Text() {
	case "#":
		//comment, do nothing
	case "v":
		vertex, err := parseVec3(scanner)
		if err != nil {
			log.Fatalf("Could not parse vector %s\n error: %v", line, err)
		}
		o.min = vec3.Min(vertex, &o.min)
		o.max = vec3.Max(vertex, &o.max)
		o.Vertices = append(o.Vertices, vertex)
	case "vn":
		normal, err := parseVec3(scanner)
		if err != nil {
			log.Fatalf("Could not parse normal %s\n error: %v", line, err)
		}
		o.Normals = append(o.Normals, normal)
	case "vt":
		tex_coord, err := parseVec2(scanner)
		if err != nil {
			log.Fatalf("Could not parse text coordinate %s\n error: %v", line, err)
		}
		o.TexCoords = append(o.TexCoords, tex_coord)
		o.HasTexCoords = true
	case "f":
		o.Faces = append(o.Faces, parseFaces(scanner)...)
	default:
		if len(line) > 0 {
			//log.Printf("Unknown token (ignored): %s", line)
		}
	}
}

func (o *Data) interpolateNormals() {
	o.Normals = make([]*vec3.T, len(o.Vertices))
	for i := range o.Faces {
		f := &o.Faces[i]
		e1 := vec3.Sub(o.Vertices[f.VertexIds[1]], o.Vertices[f.VertexIds[0]])
		e2 := vec3.Sub(o.Vertices[f.VertexIds[2]], o.Vertices[f.VertexIds[0]])
		n := vec3.Cross(&e1, &e2)
		n.Normalize()
		for j, id := range f.VertexIds {
			f.NormalIds[j] = id
			o.Normals[id] = &n
		}
	}
}

func parseVec3(scanner *bufio.Scanner) (*vec3.T, error) {
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
	return &vector, nil
}

func parseVec2(scanner *bufio.Scanner) (*vec2.T, error) {
	var vector vec2.T
	counter := 0
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			return nil, err
		}
		if counter > 1 {
			//log.Print("Found w-coordinate for texture, ignoring..")
			break
		}
		vector[counter] = float32(f)
		counter++
	}
	return &vector, nil
}

// Takes a obj face line with arbitray many vertex information and divides it into triangle faces.
func parseFaces(scanner *bufio.Scanner) []Face {
	faces := make([]Face, 0, 1)
	counter := 0
	f := Face{}
	for scanner.Scan() {
		f.VertexIds[counter], f.TexCoordIds[counter], f.NormalIds[counter] = parseFacePoint(scanner.Bytes())
		counter++
		if counter >= 3 {
			faces = append(faces, f)
			first := f //make a copy of first face
			flast := f //save last face
			for scanner.Scan() {
				// take data for first vertex from first face
				f = first
				// take data for second vertex from last face
				f.NormalIds[1] = flast.NormalIds[2]
				f.TexCoordIds[1] = flast.TexCoordIds[2]
				f.VertexIds[1] = flast.VertexIds[2]
				// get new data for third vertex
				f.VertexIds[2], f.TexCoordIds[2], f.NormalIds[2] = parseFacePoint(scanner.Bytes())
				faces = append(faces, f)
				flast = f
			}
		}
	}
	return faces
}

// parse Face according to format: "vertex/texcoord/normal"
// also accepted is "vertex/texcoord" or "vertex//normal"
func parseFacePoint(data []byte) (int, int, int) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(scanSlashes)

	vertex_id := parseId(scanner)
	texCoord_id := parseId(scanner)
	normal_id := -1
	if strings.Count(string(data), "/") == 2 {
		normal_id = parseId(scanner)
	}

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
		log.Printf("Failed to parse %v", err)
	}
	// minus one bc one based counting system in obj
	return id - 1
}

func scanSlashes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '/'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
