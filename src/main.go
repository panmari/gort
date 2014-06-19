package main

import (
	"scenes"
)

func main() {
	scene := scenes.MakeSimpleScene()
	for x := 0; x < scene.Film.GetWidth(); x++ {
		for y := 0; y < scene.Film.GetHeight(); y++ {
			sample := scene.Sampler.Get2DSample()
			ray := scene.Camera.MakeWorldSpaceRay(x, y, sample)
			color := scene.Integrator.Integrate(ray, scene.Root)
			scene.Film.AddSample(x,y,color)
		}
	}
	scene.Film.WriteToPng("test")
}
