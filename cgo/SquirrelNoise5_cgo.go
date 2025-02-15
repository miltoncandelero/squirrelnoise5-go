//go:build cgo
// +build cgo

package cgo

// #cgo CFLAGS: -Ofast -march=native -mtune=native -flto -ffast-math
// #include "SquirrelNoise5.h"
import "C"

// This is only intended for testing purposes to compare the Go implementation with the C++ implementation.
// You probably shouldn't be using this package in your project since the GO implementation is faster.
func SquirrelNoise5(positionX int32, seed uint32) uint32 {
	return uint32(C.SquirrelNoise5(C.int(positionX), C.uint(seed)))
}

func Get1dNoiseUint(positionX int32, seed uint32) uint32 {
	return uint32(C.Get1dNoiseUint(C.int(positionX), C.uint(seed)))
}

func Get2dNoiseUint(indexX int32, indexY int32, seed uint32) uint32 {
	return uint32(C.Get2dNoiseUint(C.int(indexX), C.int(indexY), C.uint(seed)))

}

func Get3dNoiseUint(indexX int32, indexY int32, indexZ int32, seed uint32) uint32 {
	return uint32(C.Get3dNoiseUint(C.int(indexX), C.int(indexY), C.int(indexZ), C.uint(seed)))
}

func Get4dNoiseUint(indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32) uint32 {
	return uint32(C.Get4dNoiseUint(C.int(indexX), C.int(indexY), C.int(indexZ), C.int(indexT), C.uint(seed)))
}

func Get1dNoiseZeroToOne(index int32, seed uint32) float32 {
	return float32(C.Get1dNoiseZeroToOne(C.int(index), C.uint(seed)))
}

func Get2dNoiseZeroToOne(indexX int32, indexY int32, seed uint32) float32 {
	return float32(C.Get2dNoiseZeroToOne(C.int(indexX), C.int(indexY), C.uint(seed)))
}

func Get3dNoiseZeroToOne(indexX int32, indexY int32, indexZ int32, seed uint32) float32 {
	return float32(C.Get3dNoiseZeroToOne(C.int(indexX), C.int(indexY), C.int(indexZ), C.uint(seed)))
}

func Get4dNoiseZeroToOne(indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32) float32 {
	return float32(C.Get4dNoiseZeroToOne(C.int(indexX), C.int(indexY), C.int(indexZ), C.int(indexT), C.uint(seed)))
}

func Get1dNoiseNegOneToOne(index int32, seed uint32) float32 {
	return float32(C.Get1dNoiseNegOneToOne(C.int(index), C.uint(seed)))
}

func Get2dNoiseNegOneToOne(indexX int32, indexY int32, seed uint32) float32 {
	return float32(C.Get2dNoiseNegOneToOne(C.int(indexX), C.int(indexY), C.uint(seed)))
}

func Get3dNoiseNegOneToOne(indexX int32, indexY int32, indexZ int32, seed uint32) float32 {
	return float32(C.Get3dNoiseNegOneToOne(C.int(indexX), C.int(indexY), C.int(indexZ), C.uint(seed)))
}

func Get4dNoiseNegOneToOne(indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32) float32 {
	return float32(C.Get4dNoiseNegOneToOne(C.int(indexX), C.int(indexY), C.int(indexZ), C.int(indexT), C.uint(seed)))
}
