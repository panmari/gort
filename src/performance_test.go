package test

import (
	"testing"
)

func BenchmarkRenderingCurrentScene(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
