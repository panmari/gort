package renderer

import (
	"cameras"
	"films"
	"samplers"
	"integrators"
	"scenes"
)

type Worker struct {
	camera 		cameras.Camera
	sampler 	func(int64) samplers.Sampler
	film 		films.Film
	integrator 	integrators.Integrator
	spp 		int
}

type Task struct {
	left, right, bottom, top int
}

func (w *Worker) renderWindow(taskChan chan *Task, taskDone chan bool) {
	for t := range taskChan {
		seed := int64(t.top*w.film.GetWidth() + t.left)
		sampler := w.sampler(seed)
		for x := t.left; x < t.right; x++ {
			for y := t.bottom; y < t.top; y++ {
				for s := 0; s < w.spp; s++ {
					sample := sampler.Get2DSample()
					ray := w.camera.MakeWorldSpaceRay(x, y, sample)
					color := w.integrator.Integrate(ray)
					w.film.AddSample(x,y,color)
				}
			}
		}
		taskDone <- true
	}
}

func NewWorker(s scenes.Scene) *Worker {
	return &Worker{s.Camera, s.Sampler, s.Film, s.Integrator, s.SPP}
}
