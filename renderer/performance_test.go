package renderer

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/panmari/gort/scenes"
)

func BenchmarkRenderingSimpleScene(b *testing.B) {
	s := scenes.NewSimpleScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !testing.Short() {
			s.SPP = 64
		}
		StartRendering(&s, false, 0*time.Second)
	}
}

func BenchmarkRenderingBoxScene(b *testing.B) {
	s := scenes.NewBoxScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//setting SPP doesn't change anything, has OneSampler
		StartRendering(&s, false, 0*time.Second)
	}
}

func BenchmarkRenderingDodecahedronScene(b *testing.B) {
	s := scenes.NewDodecahedronScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//setting SPP doesn't change anything, has OneSampler
		StartRendering(&s, false, 0*time.Second)
	}
}

func ChangeDirForReadingMeshes(b *testing.B) {
	if wd, _ := os.Getwd(); path.Base(wd) == "renderer" {
		if err := os.Chdir("../"); err != nil {
			b.Error(err)
		}
	}

}

func BenchmarkRenderingTeapotInstancingScene(b *testing.B) {
	ChangeDirForReadingMeshes(b)
	s := scenes.NewInstancingTeapotsScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartRendering(&s, false, 0*time.Second)
	}
}

func BenchmarkRenderingAcceleratorScene(b *testing.B) {
	ChangeDirForReadingMeshes(b)
	s := scenes.NewAcceleratorTestScene()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartRendering(&s, false, 0*time.Second)
	}
}
