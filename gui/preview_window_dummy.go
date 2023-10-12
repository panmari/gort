//go:build !cgo
// +build !cgo

package gui

import (
	"github.com/panmari/gort/films"
)

// Create returns a new preview window. Clients need to call
// fyne.CurrentApp().Driver().Run()
// to enter the application loop.
func Create(film films.Film) *previewWindow {
	pw := previewWindow{film: film}
	return &pw
}

func (pw *previewWindow) Update() {
	// Nothing
}
