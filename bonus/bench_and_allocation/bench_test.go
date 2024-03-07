package main

import (
	"bytes"
	"io"
	"testing"
)

func writeToBuffer(w io.Writer, msg []byte) {
	w.Write(msg)
}

func BenchmarkWriteToBuffer(b *testing.B) {
	msg := []byte("hello world")
	buf := new(bytes.Buffer)

	// this will generate 0 allocations
	for i := 0; i < b.N; i++ { // internal bench stuff handled by Golang
		for i := 0; i < 100; i++ { // our logic
			writeToBuffer(buf, msg)
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	n := 100

	b.Run("non alloc", func(b *testing.B) {
		// this will generate multiple heap allocations to resize the slice
		for i := 0; i < b.N; i++ {
			s := []int{}
			for j := 0; j < n; j++ {
				s = append(s, i)
			}
		}
	})

	b.Run("alloc", func(b *testing.B) {
		// initialization (avoiding multiple heap allocations)
		for i := 0; i < b.N; i++ {
			s := make([]int, n)
			for j := 0; j < n; j++ {
				s[j] = j
			}
		}
	})
}
