package renderer

import (
	"sync"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"github.com/cheggaaa/pb"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/scenes"
	"github.com/panmari/gort/util"
)

func StartRendering(scene *scenes.Scene, enableProgressbar bool, enablePreviewWindow bool) {
	if scene.Film.GetWidth() == 0 || scene.Film.GetHeight() == 0 || scene.SPP == 0 {
		panic("Invalid settings detected in scene!")
	}
	const taskSize = 64

	var bar util.AbstractProgressBar
	if enableProgressbar {
		pb := pb.StartNew(scene.Film.GetWidth() * scene.Film.GetHeight())
		pb.ShowSpeed = true
		bar = pb
	} else {
		bar = &util.DummyProgressBar{}
	}
	if enablePreviewWindow {
		p := PreviewWindow{film: scene.Film}
		go p.init()
		go func() {
			time.Sleep(1 * time.Second)
			p.update()
		}()
	}

	var wg sync.WaitGroup
	for x := 0; x < scene.Film.GetWidth(); x += taskSize {
		for y := 0; y < scene.Film.GetHeight(); y += taskSize {
			x_border := util.Min(x+taskSize, scene.Film.GetWidth())
			y_border := util.Min(y+taskSize, scene.Film.GetHeight())
			wg.Add(1)
			go renderWindow(*scene, x, int(x_border), y, int(y_border), &wg, bar)
		}
	}
	wg.Wait()
	bar.Finish()
}

type PreviewWindow struct {
	film films.Film
	w    fyne.Window
	c    fyne.Canvas
}

func (pw *PreviewWindow) init() {
	a := app.New()
	w := a.NewWindow("Rendering...")
	c := canvas.NewImageFromImage(pw.film)
	c.SetMinSize(fyne.NewSize(pw.film.GetWidth(), pw.film.GetHeight()))

	w.SetContent(c)
	w.SetFixedSize(true)
	go w.ShowAndRun()
	pw.w = w
}

func (pw *PreviewWindow) update() {
	c := canvas.NewImageFromImage(pw.film)
	c.SetMinSize(fyne.NewSize(pw.film.GetWidth(), pw.film.GetHeight()))
	pw.w.SetContent(c)
}

// renders a window of the given scene
func renderWindow(scene scenes.Scene, left, right, bottom, top int, wg *sync.WaitGroup, bar util.AbstractProgressBar) {
	defer wg.Done()
	seed := int64(left*scene.Film.GetWidth() + top)
	// Makes a copy of the sampler
	sampler := scene.Sampler.DuplicateAndSeed(seed, scene.SPP)
	camera := scene.Camera
	integrator := scene.Integrator
	film := scene.Film
	for x := left; x < right; x++ {
		for y := bottom; y < top; y++ {
			samples := sampler.Get2DSamples(scene.SPP)
			for s := range samples {
				ray := camera.MakeWorldSpaceRay(x, y, samples[s])
				color := integrator.Integrate(ray, 0)
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
	bar := &util.DummyProgressBar{}
	renderWindow(scene, x, x+1, y, y+1, &wg, bar) //TODO fix for progress bar
	bar.Finish()
}
