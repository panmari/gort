package films

import (
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
)

// Film represents a datastructure for capturing samples for each pixel, then allows displaying it or exporting to png.
type Film interface {
	// AddSample adds a sample at position x, y. Sample is interpereted as RGB.
	AddSample(x, y int, sample *vec3.T)
	// WriteToPng writes the image captured by this film to the given filename.
	WriteToPng(filename string)
	// GetWidth returns the width of this film.
	GetWidth() int
	// GetHeight returns the height of this film.
	GetHeight() int
	// GetToneMapper returns the function used for tone mapping each pixel.
	GetTonemapper() func(*vec4.T)
}
