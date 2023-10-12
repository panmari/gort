// Package gui provides a graphical user interface for showing render progress.
// Not available for platforms that don't support cgo, hence split into separate compilation units.
package gui

import (
	"github.com/panmari/gort/films"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type previewWindow struct {
	film films.Film
	// Window where progress is shown.
	w fyne.Window
	// The raster containing the image render so far. Needs to be updated for showing changes on screen.
	img *canvas.Raster
}
