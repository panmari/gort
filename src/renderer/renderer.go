package renderer

import (
	"scenes"
	"sync"
	"util"
	"github.com/cheggaaa/pb"
)

func StartRendering(scene scenes.Scene) {
	if scene.Film.GetWidth() == 0 || scene.Film.GetHeight() == 0 || scene.SPP == 0 {
		panic("Invalid settings for scene!")
	}
	tasksize := 64
	var wg sync.WaitGroup
	bar := pb.StartNew(scene.Film.GetWidth()*scene.Film.GetHeight())
	bar.ShowSpeed = true
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y += tasksize {
			x_border := util.Min(x+tasksize, scene.Film.GetWidth())
			y_border := util.Min(y+tasksize, scene.Film.GetHeight())
			wg.Add(1)
			go renderWindow(scene, x, int(x_border), y, int(y_border), &wg, bar)
		}
	}
	wg.Wait()
	bar.Finish()
}

// renders a window of the given scene
func renderWindow(scene scenes.Scene, left, right, bottom, top int, wg *sync.WaitGroup, bar *pb.ProgressBar) {
	defer wg.Done()
	seed := int64(left*scene.Film.GetWidth() + top)
	sampler := scene.Sampler(seed)
	camera := scene.Camera
	integrator := scene.Integrator
	film := scene.Film
	for x := left; x < right; x++ {
		for y := bottom; y < top; y++ {
			samples := sampler.Get2DSamples(scene.SPP)
			for s := range samples {
				ray := camera.MakeWorldSpaceRay(x, y, samples[s])
				color := integrator.Integrate(ray)
				film.AddSample(x, y, color)
			}
			bar.Increment()
		}
	}
}

// mainly used for debugging
func RenderPixel(scene scenes.Scene, x, y int) {
	var wg sync.WaitGroup
	wg.Add(1)
	bar := pb.StartNew(1)
	renderWindow(scene, x, x+1, y, y+1, &wg, bar) //TODO fix for progress bar
	bar.Finish()
}