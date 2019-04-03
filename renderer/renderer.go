package renderer

import (
	"sort"
	"sync"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"github.com/cheggaaa/pb"
	"github.com/panmari/gort/films"
	"github.com/panmari/gort/scenes"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec2"
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
		p.init()
		go p.w.ShowAndRun()
		go func() {
			for {
				time.Sleep(2 * time.Second)
				p.update()
			}
		}()
	}

	tasks := make([]task, 0) // TODO(panmari): Pre-allocate.
	for x := 0; x < scene.Film.GetWidth(); x += tasksize {
		for y := 0; y < scene.Film.GetHeight(); y += tasksize {
			tasks = append(tasks, task{minX: x, minY: y})
		}
	}
	sort.Slice(tasks, func(i, j int) bool {
		center := &vec2.T{float32(scene.Film.GetWidth() / 2), float32(scene.Film.GetHeight() / 2)}
		a := &vec2.T{float32(tasks[i].minX), float32(tasks[i].minY)}
		b := &vec2.T{float32(tasks[j].minX), float32(tasks[j].minY)}
		return a.Sub(center).LengthSqr() < b.Sub(center).LengthSqr()
	})
	var wg sync.WaitGroup
	for _, t := range tasks {
		maxX := util.Min(t.minX+tasksize, scene.Film.GetWidth())
		maxY := util.Min(t.minY+tasksize, scene.Film.GetHeight())
		wg.Add(1)
		go renderWindow(*scene, t.minX, int(maxX), t.minY, int(maxY), &wg, bar)
		// TODO(panmari): Instead of throttling submission of tasks here, make
		// render windows block when preview window wants to update.
		time.Sleep(100 * time.Millisecond)
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
	w.SetFixedSize(true)
	c := canvas.NewImageFromImage(pw.film)
	c.SetMinSize(fyne.NewSize(pw.film.GetWidth(), pw.film.GetHeight()))

	w.SetContent(c)
	pw.w = w
}

func (pw *PreviewWindow) update() {
	pw.w.Canvas().Refresh(pw.w.Content())
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
	// TODO(panmari): Change iteration order (every pixel 1 sample first) for preview to work better.
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
