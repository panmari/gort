package main

import (
	"scenes"
	"time"
	"fmt"
)

func main() {
	
	scene := scenes.MakeSimpleScene()
	
	start := time.Now()
	for x := 0; x < scene.Film.GetWidth(); x++ {
		for y := 0; y < scene.Film.GetHeight(); y++ {
			for s := 0; s < scene.SPP; s++ {
				sample := scene.Sampler.Get2DSample()
				ray := scene.Camera.MakeWorldSpaceRay(x, y, sample)
				color := scene.Integrator.Integrate(ray)
				scene.Film.AddSample(x,y,color)
			}
		}
	}
	duration := time.Since(start)
	fmt.Println(duration.String())
	scene.Film.WriteToPng(scene.Filename)
	fmt.Printf("Printed to %s", scene.Filename)
}
