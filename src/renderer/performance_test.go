package renderer

import (
	"os"
	"path"
	"runtime"
	"scenes"
	"testing"
	//"fmt"
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

func BenchmarkRenderingDodecahedronScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := scenes.NewDodecahedronScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//setting SPP doesn't change anything, has OneSampler
		StartRendering(s)
	}
}

func BenchmarkRenderingTeapotInstancingScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if wd, _ := os.Getwd(); path.Base(wd) == "renderer" {
		if err := os.Chdir("../../"); err != nil {
			b.Error(err)
		}
	}
	s := scenes.NewInstancingTeapotsScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartRendering(s)
	}
}
