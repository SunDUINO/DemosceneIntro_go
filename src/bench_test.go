package main

import "testing"

func BenchmarkCreateStars(b *testing.B) {
    for i := 0; i < b.N; i++ {
        createStars()
    }
}

func BenchmarkUpdateStars(b *testing.B) {
    a := newApp()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        a.update()
    }
}

func BenchmarkDrawCone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO: wywołaj drawCone(...)
		_ = i
	}
}

func BenchmarkDrawPyramid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO: wywołaj drawPyramid(...)
		_ = i
	}
}
