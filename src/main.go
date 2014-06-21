package main

import (
	"scenes"
	"time"
	"fmt"
	"util"
	"sync"
)

func main() {
	scene := scenes.MakeSimpleScene()
	
	tasksize := 4
	start := time.Now()
	var wg sync.WaitGroup
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y+= tasksize {
			wg.Add(1)
			x_border := util.Min(x + tasksize, scene.Film.GetWidth())
			y_border := util.Min(y + tasksize, scene.Film.GetHeight())
			go renderWindow(scene, x, int(x_border), y, int(y_border), &wg)
		}
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println(duration.String())
	scene.Film.WriteToPng(scene.Filename)
	fmt.Printf("Printed to %s", scene.Filename)
}

func renderWindow(scene *scenes.Scene, left, right, bottom, top int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := left; x < right; x++ {
		for y := bottom; y < top; y++ {
			for s := 0; s < scene.SPP; s++ {
				sample := scene.Sampler.Get2DSample()
				ray := scene.Camera.MakeWorldSpaceRay(x, y, sample)
				color := scene.Integrator.Integrate(ray)
				scene.Film.AddSample(x,y,color)
			}
		}
	}
}
