package cufft

// This file provides CGO flags to find CUDA libraries and headers.

//#cgo LDFLAGS:-lcufft
//
////default location:
//#cgo LDFLAGS:-L/usr/local/cuda/lib64 -L/usr/local/cuda/lib
//#cgo CFLAGS: -I/usr/local/cuda/include/
//
////default location if not properly symlinked:
//#cgo LDFLAGS:-L/usr/local/cuda-6.0/lib64 -L/usr/local/cuda-6.0/lib
//#cgo LDFLAGS:-L/usr/local/cuda-5.5/lib64 -L/usr/local/cuda-5.5/lib
//#cgo LDFLAGS:-L/usr/local/cuda-5.0/lib64 -L/usr/local/cuda-5.0/lib
//#cgo CFLAGS: -I/usr/local/cuda-6.0/include/
//#cgo CFLAGS: -I/usr/local/cuda-5.5/include/
//#cgo CFLAGS: -I/usr/local/cuda-5.0/include/
//
////arch linux:
//#cgo LDFLAGS:-L/usr/local/cuda/lib64 -L/usr/local/cuda/lib
//#cgo CFLAGS: -I/usr/local/cuda/include
//
////WINDOWS:
//#cgo windows LDFLAGS:-LC:/cuda/v5.0/lib/x64 -LC:/cuda/v5.5/lib/x64 -LC:/cuda/v6.0/lib/x64
//#cgo windows CFLAGS: -IC:/cuda/v5.0/include -IC:/cuda/v5.5/include -IC:/cuda/v6.0/include -w
import "C"
