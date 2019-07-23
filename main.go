package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/panmari/gort/renderer"
	"github.com/panmari/gort/scenes"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var startServer = flag.Bool("server", false, "Start in server mode")
var startClient = flag.Bool("client", false, "Start in client mode")
var progressBar = flag.Bool("progress_bar", true, "If true, shows a progress bar")
var previewUpdateInterval = flag.Duration("preview_update_interval", 0*time.Second, "If non 0, shows a preview window that updates in the given interval. E.g. 2s")

func main() {
	flag.Parse()
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
	scene := scenes.NewSimpleCSGScene()
	//scene := scenes.NewInstancingTeapotsScene()
	//scene := scenes.NewObjReaderTestScene()
	// scene := scenes.NewAcceleratorTestScene()
	//scene := scenes.NewMassiveAcceleratorTestScene()

	if *startClient {
		renderer.RenderOnServer(&scene)
	} else {
		start := time.Now()
		renderer.StartRendering(&scene, *progressBar, *previewUpdateInterval)
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
