package renderer

import (
	"scenes"
	"sync"
	"util"
)

func StartRendering(scene scenes.Scene) {
	tasksize := 64
	var wg sync.WaitGroup
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y += tasksize {
			x_border := util.Min(x+tasksize, scene.Film.GetWidth())
			y_border := util.Min(y+tasksize, scene.Film.GetHeight())
			wg.Add(1)
			go renderWindow(scene, x, int(x_border), y, int(y_border), &wg)
		}
	}
	wg.Wait()
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