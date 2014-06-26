package test

import (
	"testing"
	"renderer"
	"scenes"
)

func BenchmarkRenderingCurrentScene(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := scenes.MakeSimpleScene()
		renderer.StartRendering(s)
	}
}
