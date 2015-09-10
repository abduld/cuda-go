package cu

// This file provides CGO flags to find CUDA libraries and headers.

//#cgo LDFLAGS: -lcuda -lcudart -lcudadevrt
//
////default location:
//#cgo LDFLAGS:-L/usr/local/cuda/lib64 -L/usr/local/cuda/lib
//#cgo CFLAGS: -I/usr/local/cuda/include/
//
////arch linux:
//#cgo linux LDFLAGS:-L/usr/local/cuda/lib64 -L/usr/local/cuda/lib -lcuda -lcudart
//#cgo linux CFLAGS: -I/usr/local/cuda/include
////arch OSX:
//#cgo darwin LDFLAGS: -L/usr/local/cuda/lib
//#cgo darwin CFLAGS: -I/usr/local/cuda/include -framework CUDA
//
////WINDOWS:
//#cgo windows LDFLAGS:-LC:/cuda/v5.0/lib/x64 -LC:/cuda/v5.5/lib/x64 -LC:/cuda/v6.0/lib/x64
//#cgo windows CFLAGS: -IC:/cuda/v5.0/include -IC:/cuda/v5.5/include -IC:/cuda/v6.0/include
import "C"
