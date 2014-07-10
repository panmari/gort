package renderer

import (
	"cameras"
	"samplers"
	"integrators"
	"github.com/ungerik/go3d/vec3"
	"scenes"
)

type Sample struct {
	x, y int
	color *vec3.T
}

type Worker struct {
	camera cameras.Camera
	sampler func(int64) samplers.Sampler
	integrator integrators.Integrator
	spp int
}

type Task struct {
	left, right, bottom, top int
}

func (w *Worker) renderWindow(taskChan chan *Task, sampleChan chan *Sample) {
	for t := range taskChan {
		//TODO: make seed correctly (left*width)
		seed := int64(t.left+ t.right)
		sampler := w.sampler(seed)
		for x := t.left; x < t.right; x++ {
			for y := t.bottom; y < t.top; y++ {
				for s := 0; s < w.spp; s++ {
					sample := sampler.Get2DSample()
					ray := w.camera.MakeWorldSpaceRay(x, y, sample)
					color := w.integrator.Integrate(ray)
					sampleChan <- &Sample{x,y,color}
				}
			}
		}
	}
}

func NewWorker(s scenes.Scene) *Worker {
	return &Worker{s.Camera, s.Sampler, s.Integrator, s.SPP}
}
