package renderer

import (
	"testing"
	"scenes"
)

func BenchmarkRenderingSimpleScene(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := scenes.MakeSimpleScene()
		StartRendering(s)
	}
}
