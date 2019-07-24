package gui

import (
	"github.com/panmari/gort/films"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

// Create returns a new preview window. Clients need to call
// fyne.CurrentApp().Driver().Run()
// to enter the application loop.
func Create(film films.Film) *previewWindow {
	pw := previewWindow{film: film}
	a := app.New()
	w := a.NewWindow("Rendering...")
	w.SetFixedSize(true)
	img := canvas.NewRasterFromImage(pw.film)
	// TODO(panmari): Window is not shown if img size is not set.
	// TODO(panmari): GUI-scaling in gnome/wayland isn't taken properly into account here.
	img.SetMinSize(fyne.NewSize(pw.film.GetWidth(), pw.film.GetHeight()))
	w.SetContent(img)
	w.Show()
	pw.w = w
	pw.img = img
	return &pw
}

type previewWindow struct {
	film films.Film
	w    fyne.Window
	img  *canvas.Raster
}

func (pw *previewWindow) Update() {
	canvas.Refresh(pw.img)
}
