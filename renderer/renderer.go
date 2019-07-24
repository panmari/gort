package renderer

import (
	"runtime"
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

const taskSize = 64

func StartRendering(scene *scenes.Scene, enableProgressbar bool, previewUpdateInterval time.Duration) {
	if scene.Film.GetWidth() == 0 || scene.Film.GetHeight() == 0 || scene.SPP == 0 {
		panic("Invalid settings detected in scene!")
	}

	var bar util.AbstractProgressBar
	if enableProgressbar {
		pb := pb.StartNew(scene.Film.GetWidth() * scene.Film.GetHeight())
		pb.ShowSpeed = true
		bar = pb
	} else {
		bar = &util.DummyProgressBar{}
	}

	if previewUpdateInterval > 0*time.Second {
		p := PreviewWindow{film: scene.Film}
		p.init()
		p.w.Show()
		go func() {
			for {
				time.Sleep(previewUpdateInterval)
				p.update()
			}
		}()
	}

	tasks := make([]task, 0, scene.Film.GetWidth()*scene.Film.GetHeight()/(taskSize*taskSize)+1)
	for x := 0; x < scene.Film.GetWidth(); x += taskSize {
		for y := 0; y < scene.Film.GetHeight(); y += taskSize {
			tasks = append(tasks, task{scene: scene, minX: x, minY: y})
		}
	}
	sort.Slice(tasks, func(i, j int) bool {
		center := &vec2.T{float32(scene.Film.GetWidth() / 2), float32(scene.Film.GetHeight() / 2)}
		a := &vec2.T{float32(tasks[i].minX), float32(tasks[i].minY)}
		b := &vec2.T{float32(tasks[j].minX), float32(tasks[j].minY)}
		return a.Sub(center).LengthSqr() < b.Sub(center).LengthSqr()
	})
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	taskChan := make(chan task, 128)

	for w := 0; w < runtime.NumCPU(); w++ {
		go worker(taskChan, &wg, bar)
	}

	for _, t := range tasks {
		taskChan <- t
	}
	// Either block on rendering finishing or the UI closing.
	if previewUpdateInterval > 0*time.Second {
		go blockOnRendering(&wg, bar)
		fyne.CurrentApp().Driver().Run()
	} else {
		blockOnRendering(&wg, bar)
	}
}

func blockOnRendering(wg *sync.WaitGroup, bar util.AbstractProgressBar) {
	wg.Wait()
	bar.Finish()
}

func worker(taskChan <-chan task, wg *sync.WaitGroup, bar util.AbstractProgressBar) {
	for t := range taskChan {
		maxX := util.Min(t.minX+taskSize, t.scene.Film.GetWidth())
		maxY := util.Min(t.minY+taskSize, t.scene.Film.GetHeight())
		renderWindow(t.scene, t.minX, int(maxX), t.minY, int(maxY), wg, bar)
	}
}

type task struct {
	scene      *scenes.Scene
	minX, minY int
}

type PreviewWindow struct {
	film films.Film
	w    fyne.Window
	img  *canvas.Raster
}

func (pw *PreviewWindow) init() {
	a := app.New()
	w := a.NewWindow("Rendering...")
	w.SetFixedSize(true)
	img := canvas.NewRasterFromImage(pw.film)
	// TODO(panmari): Window is not shown if img size is not set.
	img.SetMinSize(fyne.NewSize(pw.film.GetWidth(), pw.film.GetHeight()))
	w.SetContent(img)
	pw.w = w
	pw.img = img
}

func (pw *PreviewWindow) update() {
	canvas.Refresh(pw.img)
}

// renderWindow renders all pixels in the rectangle ((left, bottom), (top, right)) of the given scene.
func renderWindow(scene *scenes.Scene, left, right, bottom, top int, wg *sync.WaitGroup, bar util.AbstractProgressBar) {
	defer wg.Done()
	seed := int64(left*scene.Film.GetWidth() + top)
	// Sampler might have internal state, so make a copy here.
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

// RenderPixel renders only one single pixel at (x, y). Useful for debugging.
func RenderPixel(scene *scenes.Scene, x, y int) {
	var wg sync.WaitGroup
	wg.Add(1)
	bar := &util.DummyProgressBar{}
	renderWindow(scene, x, x+1, y, y+1, &wg, bar) //TODO fix for progress bar
	bar.Finish()
}
