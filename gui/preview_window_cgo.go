//go:build cgo
// +build cgo

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/panmari/gort/films"
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
	img.SetMinSize(fyne.NewSize(float32(pw.film.GetWidth()), float32(pw.film.GetHeight())))
	w.SetContent(img)
	w.Show()
	pw.w = w
	pw.img = img
	return &pw
}

func (pw *previewWindow) Update() {
	canvas.Refresh(pw.img)
}
