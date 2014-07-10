package renderer

import (
	"scenes"
	"sync"
	"util"
	"runtime"
	"fmt"
)

func StartRendering(scene scenes.Scene) {
	if scene.Film.GetWidth() == 0 || scene.Film.GetHeight() == 0 || scene.SPP == 0 {
		panic("Invalid settings for scene!")
	}
	tasksize := 64
	sampleChan := make(chan *Sample, 100)
	taskChan := make(chan *Task, 100)	
	for i := 0; i < runtime.NumCPU(); i++ {
		go NewWorker(scene).renderWindow(taskChan, sampleChan)
	}
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y+= tasksize {
			x_border := util.Min(x + tasksize, scene.Film.GetWidth())
			y_border := util.Min(y + tasksize, scene.Film.GetHeight())
			taskChan <- &Task{x, int(x_border), y, int(y_border)}
		}
	}
	close(taskChan)
	
	nSamples := scene.Film.GetHeight()*scene.Film.GetWidth()*scene.SPP
	for s := 0; s < nSamples; s++ {
		sample := <- sampleChan
		scene.Film.AddSample(sample.x, sample.y, sample.color)
		if s % 100 == 0 {
			fmt.Print("*")
		}
	}
	close(sampleChan)
}

// renders a window of the given scene
func renderWindow(scene scenes.Scene, left, right, bottom, top int, wg *sync.WaitGroup) {
	defer wg.Done()
	seed := int64(left*scene.Film.GetWidth() + top)
	sampler := scene.Sampler(seed)
	camera := scene.Camera
	integrator := scene.Integrator
	film := scene.Film
	for x := left; x < right; x++ {
		for y := bottom; y < top; y++ {
			for s := 0; s < scene.SPP; s++ {
				sample := sampler.Get2DSample()
				ray := camera.MakeWorldSpaceRay(x, y, sample)
				color := integrator.Integrate(ray)
				film.AddSample(x, y, color)
			}
		}
	}
}

// mainly used for debugging
func RenderPixel(scene scenes.Scene, x, y int) {
	var wg sync.WaitGroup
	wg.Add(1)
	renderWindow(scene, x, x+1, y, y+1, &wg)
}