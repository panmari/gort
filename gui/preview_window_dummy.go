//go:build !cgo
// +build !cgo

package gui

import (
	"github.com/panmari/gort/films"
)

// Create returns a dummy preview window. It does nothing.
func Create(film films.Film) *previewWindow {
	pw := previewWindow{film: film}
	return &pw
}

func (pw *previewWindow) Update() {
	// Nothing
}
