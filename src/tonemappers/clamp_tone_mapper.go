package tonemappers

import (
	"github.com/ungerik/go3d/vec4"
)

func ClampToneMap(s *vec4.T) {
	s.Clamp(&vec4.Zero, &vec4.T{255, 255, 255})
}
