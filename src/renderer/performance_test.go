package renderer

import (
	"testing"
	"scenes"
	"runtime"
)

func BenchmarkRenderingSimpleScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := scenes.NewSimpleScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !testing.Short() {
			s.SPP = 64
		}
		StartRendering(s)
	}
}

func BenchmarkRenderingBoxScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := scenes.NewBoxScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//setting SPP doesn't change anything, has OneSampler
		StartRendering(s)
	}
}
