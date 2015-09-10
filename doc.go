/*
	Go bindings for nVIDIA CUDA 5.
	This package compiles with both gc and gccgo.
*/
package cuda5

// Dummy imports so that
// 	go get github.com/abduld/cuda5
// will install everything.
import (
	_ "github.com/abduld/cuda5/cu"
	_ "github.com/abduld/cuda5/cufft"
	_ "github.com/abduld/cuda5/safe"
)
