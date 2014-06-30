package films

import (
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
	"image"
	"image/color"
	"image/png"
	"os"
)

type BoxFilterFilm struct {
	Film          []vec4.T
	width, height int
}

func (f *BoxFilterFilm) AddSample(x, y int, sample *vec3.T) {
	t := f.Film[y*f.width+x]
	t[0] += sample[0]
	t[1] += sample[1]
	t[2] += sample[2]
	t[3] += 1
}

func (i *BoxFilterFilm) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *BoxFilterFilm) At(x, y int) color.Color {
	// invert y axis here
	y = i.height - y - 1
	s := i.Film[y*i.width+x]
	s.Scale(255.0 / s[3])
	s.Clamp(&vec4.Zero, &vec4.T{255, 255, 255})
	return color.RGBA{uint8(s[0]), uint8(s[1]), uint8(s[2]), 255}
}

func (i *BoxFilterFilm) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i *BoxFilterFilm) WriteToPng(filename string) {
	fo, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}
	png.Encode(fo, i)
}

func MakeBoxFilterFilm(w, h int) *BoxFilterFilm {
	return &BoxFilterFilm{
		width:  w,
		height: h,
		Film:   make([]vec4.T, w*h)}
}

func (i *BoxFilterFilm) GetWidth() int {
	return i.width
}

func (i *BoxFilterFilm) GetHeight() int {
	return i.height
}
