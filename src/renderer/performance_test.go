package renderer

import (
	"testing"
	"scenes"
	"runtime"
)

func BenchmarkRenderingSimpleScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		s := scenes.MakeSimpleScene()
		if !testing.Short() {
			s.SPP = 64
		}
		StartRendering(s)
	}
}
