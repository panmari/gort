package main

import (
	"scenes"
	"time"
	"fmt"
	"util"
	"runtime"
	"github.com/ungerik/go3d/vec3"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	scene := scenes.MakeSimpleScene()
	
	tasksize := 32
	start := time.Now()
	sampleChan := make(chan *Sample, 10)
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
	duration := time.Since(start)
	fmt.Println(duration.String())
	scene.Film.WriteToPng(scene.Filename)
	fmt.Printf("Printed to %s", scene.Filename)
}

type Sample struct {
	x, y int
	color *vec3.T
}

func renderWindow(scene scenes.Scene, left, right, bottom, top int, sampleChan chan *Sample) {
	for x := left; x < right; x++ {
		for y := bottom; y < top; y++ {
			for s := 0; s < scene.SPP; s++ {
				sample := scene.Sampler.Get2DSample()
				ray := scene.Camera.MakeWorldSpaceRay(x, y, sample)
				color := scene.Integrator.Integrate(ray)
				sampleChan <- &Sample{x,y,color}
			}
		}
	}
}
