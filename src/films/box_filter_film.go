package films

import (
		"github.com/ungerik/go3d/vec4"
		"github.com/ungerik/go3d/vec3"
		"image"
		"image/color"
		"os"
		"image/png"
)

type BoxFilterFilm struct {
	Film []vec4.T
	Width, Height int
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
	//TODO: make more sense here
	s := i.Film[y*i.Width+x]
	s.Scale(255.0/s[3])
	s.Clamp(&vec4.Zero,&vec4.T{255,255,255})
	return color.RGBA{uint8(s[0]),uint8(s[1]),uint8(s[2]), 255}
}

func (i *BoxFilterFilm) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i *BoxFilterFilm) WriteToPng(filename string) {
	fo, err := os.Create(filename + ".png")
	if err != nil { panic(err) }
	png.Encode(fo, i)
}

func MakeBoxFilterFilm(w, h int) *BoxFilterFilm {
	return &BoxFilterFilm {
		Width:  w,
		Height: h,
		Film:    make([]vec4.T, w*h)}
}
