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
	Width, Height int
	Tonemapper    func(*vec4.T)
}

func (f *BoxFilterFilm) AddSample(x, y int, sample *vec3.T) {
	t := &f.Film[y*f.Width+x]
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
	y = i.Height - y - 1
	s := i.Film[y*i.Width+x]
	s.Scale(255.0 / s[3])
	i.Tonemapper(&s)
	return color.RGBA{uint8(s[0]), uint8(s[1]), uint8(s[2]), 255}
}

func (i *BoxFilterFilm) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i *BoxFilterFilm) WriteToPng(filename string) {
	if err := os.Mkdir("output", os.ModePerm); err != nil && !os.IsExist(err) {
		panic(err)
	}
	fo, err := os.Create("output/" + filename + ".png")
	if err != nil {
		panic(err)
	}
	if err = png.Encode(fo, i); err != nil {
		panic(err)
	}
}

func MakeBoxFilterFilm(w, h int, tonemapper func(*vec4.T)) *BoxFilterFilm {
	return &BoxFilterFilm{
		Width:      w,
		Height:     h,
		Tonemapper: tonemapper,
		Film:       make([]vec4.T, w*h)}
}

func (i *BoxFilterFilm) GetWidth() int {
	return i.Width
}

func (i *BoxFilterFilm) GetHeight() int {
	return i.Height
}
