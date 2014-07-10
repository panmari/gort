package renderer

import (
	"scenes"
	"sync"
	"util"
	"runtime"
	"github.com/cheggaaa/pb"
)

func StartRendering(scene scenes.Scene) {
	if scene.Film.GetWidth() == 0 || scene.Film.GetHeight() == 0 || scene.SPP == 0 {
		panic("Invalid settings for scene!")
	}
	tasksize := 64
	taskDone := make(chan bool)
	taskChan := make(chan *Task, 500)	
	for i := 0; i < runtime.NumCPU(); i++ {
		go NewWorker(scene).renderWindow(taskChan, taskDone)
	}
	nrTasks := 0
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y += tasksize {
			x_border := util.Min(x + tasksize, scene.Film.GetWidth())
			y_border := util.Min(y + tasksize, scene.Film.GetHeight())
			taskChan <- &Task{x, int(x_border), y, int(y_border)}
			nrTasks++
		}
	}
	close(taskChan)
	bar := pb.StartNew(nrTasks)
	for nrTasks > bar.Increment() {
		<- taskDone
	}
	bar.Finish()
}

// renders a window of the given scene
func renderWindow(scene scenes.Scene, left, right, bottom, top int, wg *sync.WaitGroup) {
	defer wg.Done()
	seed := int64(top*scene.Film.GetWidth() + left)
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