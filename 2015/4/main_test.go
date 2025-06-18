package main

import "testing"

func BenchmarkMine(b *testing.B) {
	for b.Loop() {
		mine([]byte("abcdef"))
	}	
}

func BenchmarkMineWithPrefix(b *testing.B) {
	for b.Loop() {
		mineWithPrefix([]byte("abcdef"), "000000")
	}
}
