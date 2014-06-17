package films

import (
	"github.com/ungerik/go3d/vec3"
)

type Film interface {
	AddSample(x, y int, sample *vec3.T)
	WriteToPng(filename string)
	GetWidth() int
	GetHeight() int
}