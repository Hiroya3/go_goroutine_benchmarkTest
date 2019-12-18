package main

import "testing"

func BenchmarkRunInTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runInTime()
	}
}

func BenchmarkRunConccurent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runConccurent()
	}
}
