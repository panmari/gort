package obj

import (
	"os"
	"bufio"
	"strconv"
	"log"
	"strings"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec2"
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
		myObjData.parseLine(scanner.Text())
	}
	return nil, scanner.Err()
}

type Face [3]int

type ObjData struct{
	vertices []vec3.T
	texCoords []vec2.T
	normals []vec3.T
	faces []Face
}



func (o *ObjData) parseLine(line string) {
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
		//TODO, split according to format: "vertex/texcoord/normal"
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Could not turn into float: %v", scanner.Text())
		}
		face[counter] = i
		counter++
	}
	return face
}