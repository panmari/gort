package main

import (
	"scenes"
	"time"
	"fmt"
	"util"
	"runtime"
	"github.com/ungerik/go3d/vec3"
	"flag"
	"os"
	"log"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

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
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	scene := scenes.MakeSimpleScene()
	
	tasksize := 64
	start := time.Now()
	sampleChan := make(chan Sample, 10)
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y+= tasksize {
			x_border := util.Min(x + tasksize, scene.Film.GetWidth())
			y_border := util.Min(y + tasksize, scene.Film.GetHeight())
			go renderWindow(scene, x, int(x_border), y, int(y_border), sampleChan)
		}
	}
	for s := 0; s < scene.Film.GetHeight()*scene.Film.GetWidth()*scene.SPP; s++ {
		sample := <- sampleChan
		scene.Film.AddSample(sample.x, sample.y, sample.color)
	}
	close(sampleChan)
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

type Sample struct {
	x, y int
	color *vec3.T
}

func renderWindow(scene scenes.Scene, left, right, bottom, top int, sampleChan chan Sample) {
	sampler := scene.Sampler
	camera := scene.Camera
	integrator := scene.Integrator
	for x := left; x < right; x++ {
		for y := bottom; y < top; y++ {
			for s := 0; s < scene.SPP; s++ {
				sample := sampler.Get2DSample()
				ray := camera.MakeWorldSpaceRay(x, y, sample)
				color := integrator.Integrate(ray)
				sampleChan <- Sample{x,y,color}
			}
		}
	}
}
