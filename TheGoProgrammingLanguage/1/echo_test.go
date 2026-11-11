package main

import (
	"testing"
)

var result string

var args []string

func init() {
	for i := 0; i < 10000; i++ {
		args = append(args, "golang")
	}
}

func BenchmarkManual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = manual(args)
	}
}

func BenchmarkJoinfunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = joinfunc(args)
	}
}
