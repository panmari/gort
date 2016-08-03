package renderer

import (
	"github.com/panmari/gort/scenes"
	"os"
	"path"
	"runtime"
	"testing"
)

func BenchmarkRenderingSimpleScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := scenes.NewSimpleScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !testing.Short() {
			s.SPP = 64
		}
		StartRendering(&s, false)
	}
}

func BenchmarkRenderingBoxScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := scenes.NewBoxScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//setting SPP doesn't change anything, has OneSampler
		StartRendering(&s, false)
	}
}

func BenchmarkRenderingDodecahedronScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := scenes.NewDodecahedronScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//setting SPP doesn't change anything, has OneSampler
		StartRendering(&s, false)
	}
}

func ChangeDirForReadingMeshes(b *testing.B) {
	if wd, _ := os.Getwd(); path.Base(wd) == "renderer" {
		if err := os.Chdir("../../"); err != nil {
			b.Error(err)
		}
	}

}

func BenchmarkRenderingTeapotInstancingScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ChangeDirForReadingMeshes(b)
	s := scenes.NewInstancingTeapotsScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartRendering(&s, false)
	}
}

func BenchmarkRenderingAcceleratorScene(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ChangeDirForReadingMeshes(b)
	s := scenes.NewAcceleratorTestScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartRendering(&s, false)
	}
}
