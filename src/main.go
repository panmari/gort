package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime"
	"scenes"
	"time"
	"renderer"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var	maxProcs = flag.Int("procs", runtime.NumCPU(), "set the number of processors to use")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	runtime.GOMAXPROCS(*maxProcs)	
	start := time.Now()
	
	
	// define the scene to be rendered here
	scene := scenes.NewSimpleScene()
	//scene := scenes.NewTriangleTestScene()
	//scene := scenes.NewDodecahedronScene()
	//scene := scenes.NewBoxScene()
	//scene := scenes.NewSimpleCSGScene()
	renderer.StartRendering(scene)
	//renderer.RenderPixel(scene, 300, 300)
	
	duration := time.Since(start)
	fmt.Println(duration.String())
	scene.Film.WriteToPng(scene.Filename)
	fmt.Printf("Printed to %s\n", scene.Filename)
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
