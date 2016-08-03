package main

import (
	"flag"
	"fmt"
	"github.com/panmari/gort/renderer"
	"github.com/panmari/gort/scenes"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var maxProcs = flag.Int("procs", runtime.NumCPU(), "set the number of processors to use")
var startServer = flag.Bool("server", false, "Start in server mode")
var startClient = flag.Bool("client", false, "Start in client mode")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*maxProcs)
	if *startServer {
		renderer.StartServer()
		return
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	// define the scene to be rendered here
	//scene := scenes.NewSimpleScene()
	//scene := scenes.NewTriangleTestScene()
	//scene := scenes.NewDodecahedronScene()
	//scene := scenes.NewBoxScene()
	//scene := scenes.NewSimpleCSGScene()
	//scene := scenes.NewInstancingTeapotsScene()
	scene := scenes.NewObjReaderTestScene()
	//scene := scenes.NewAcceleratorTestScene()
	//scene := scenes.NewMassiveAcceleratorTestScene()

	if *startClient {
		renderer.RenderOnServer(&scene)
	} else {
		start := time.Now()
		renderer.StartRendering(&scene, true)
		//renderer.RenderPixel(scene, 300, 300)

		duration := time.Since(start)
		fmt.Printf("Render time: %s\n", duration.String())
		scene.Film.WriteToPng(scene.Filename)
		fmt.Printf("Wrote result to to %s\n", scene.Filename)
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
}
