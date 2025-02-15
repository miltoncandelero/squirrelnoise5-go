//go:build cgo
// +build cgo

// The main purpose for these tests is to ensure that the GO implementation yields the exact same results as the C version.
// As a side note, it benchmarks the GO implementation against the C version.
// However, due to the cost of the CGO bridge, the GO implementation always beats the CGO one.

package squirrelnoise5_test

import (
	"math"
	"testing"
	"time"

	. "github.com/miltoncandelero/squirrelnoise5-go"
	cgo "github.com/miltoncandelero/squirrelnoise5-go/cgo"
)

// These tests compare the golang with the CGO implementation to try to ensure consistency.

// Shared, runtime defined seed to prevent go from optimizing away the math.
var seed = uint32(time.Now().UnixNano())

// Evenly distributed sample inside the int32 range
const loops = 65536
const step = int32(math.MaxInt32 / (loops / 2))
const from = int32(-loops / 2)
const to = loops / 2

func TestSquirrelNoise5(t *testing.T) {
	for i := from; i < to; i++ {
		position := i * step
		want := cgo.SquirrelNoise5(position, seed)
		got := SquirrelNoise5(position, seed)
		if got != want {
			t.Errorf("SquirrelNoise5(%d, %d) != cgo.SquirrelNoise5(%d, %d) => Got: %d | Want: %d", position, seed, position, seed, got, want)
		}
	}
}

func TestGet1dNoiseUint(t *testing.T) {
	for i := from; i < to; i++ {
		position := i * step
		want := cgo.Get1dNoiseUint(position, seed)
		got := Get1dNoiseUint(position, seed)
		if got != want {
			t.Errorf("Get1dNoiseUint(%d, %d) != cgo.Get1dNoiseUint(%d, %d) => Got: %d | Want: %d", position, seed, position, seed, got, want)
		}
	}
}
func TestGet2dNoiseUint(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		want := cgo.Get2dNoiseUint(positionX, positionY, seed)
		got := Get2dNoiseUint(positionX, positionY, seed)
		if got != want {
			t.Errorf("Get2dNoiseUint(%d, %d, %d) != cgo.Get2dNoiseUint(%d, %d, %d) => Got: %d | Want: %d", positionX, positionY, seed, positionX, positionY, seed, got, want)
		}
	}
}

func TestGet3dNoiseUint(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		positionZ := (i + 2) * step
		want := cgo.Get3dNoiseUint(positionX, positionY, positionZ, seed)
		got := Get3dNoiseUint(positionX, positionY, positionZ, seed)
		if got != want {
			t.Errorf("Get3dNoiseUint(%d, %d, %d, %d) != cgo.Get3dNoiseUint(%d, %d, %d, %d) => Got: %d | Want: %d", positionX, positionY, positionZ, seed, positionX, positionY, positionZ, seed, got, want)
		}
	}
}

func TestGet4dNoiseUint(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		positionZ := (i + 2) * step
		positionT := (i + 3) * step
		want := cgo.Get4dNoiseUint(positionX, positionY, positionZ, positionT, seed)
		got := Get4dNoiseUint(positionX, positionY, positionZ, positionT, seed)
		if got != want {
			t.Errorf("Get4dNoiseUint(%d, %d, %d, %d, %d) != cgo.Get4dNoiseUint(%d, %d, %d, %d, %d) => Got: %d | Want: %d", positionX, positionY, positionZ, positionT, seed, positionX, positionY, positionZ, positionT, seed, got, want)
		}
	}
}

func TestGet1dNoiseZeroToOne(t *testing.T) {
	for i := from; i < to; i++ {
		position := i * step
		want := cgo.Get1dNoiseZeroToOne(position, seed)
		got := Get1dNoiseZeroToOne(position, seed)
		if got != want {
			t.Errorf("Get1dNoiseZeroToOne(%d, %d) != cgo.Get1dNoiseZeroToOne(%d, %d) => Got: %f | Want: %f", position, seed, position, seed, got, want)
		}
	}
}

func TestGet2dNoiseZeroToOne(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		want := cgo.Get2dNoiseZeroToOne(positionX, positionY, seed)
		got := Get2dNoiseZeroToOne(positionX, positionY, seed)
		if got != want {
			t.Errorf("Get2dNoiseZeroToOne(%d, %d, %d) != cgo.Get2dNoiseZeroToOne(%d, %d, %d) => Got: %f | Want: %f", positionX, positionY, seed, positionX, positionY, seed, got, want)
		}
	}
}

func TestGet3dNoiseZeroToOne(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		positionZ := (i + 2) * step
		want := cgo.Get3dNoiseZeroToOne(positionX, positionY, positionZ, seed)
		got := Get3dNoiseZeroToOne(positionX, positionY, positionZ, seed)
		if got != want {
			t.Errorf("Get3dNoiseZeroToOne(%d, %d, %d, %d) != cgo.Get3dNoiseZeroToOne(%d, %d, %d, %d) => Got: %f | Want: %f", positionX, positionY, positionZ, seed, positionX, positionY, positionZ, seed, got, want)
		}
	}
}

func TestGet4dNoiseZeroToOne(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		positionZ := (i + 2) * step
		positionT := (i + 3) * step
		want := cgo.Get4dNoiseZeroToOne(positionX, positionY, positionZ, positionT, seed)
		got := Get4dNoiseZeroToOne(positionX, positionY, positionZ, positionT, seed)
		if got != want {
			t.Errorf("Get4dNoiseZeroToOne(%d, %d, %d, %d, %d) != cgo.Get4dNoiseZeroToOne(%d, %d, %d, %d, %d) => Got: %f | Want: %f", positionX, positionY, positionZ, positionT, seed, positionX, positionY, positionZ, positionT, seed, got, want)
		}
	}
}

func TestGet1dNoiseNegOneToOne(t *testing.T) {
	for i := from; i < to; i++ {
		position := i * step
		want := cgo.Get1dNoiseNegOneToOne(position, seed)
		got := Get1dNoiseNegOneToOne(position, seed)
		if got != want {
			t.Errorf("Get1dNoiseNegOneToOne(%d, %d) != cgo.Get1dNoiseNegOneToOne(%d, %d) => Got: %f | Want: %f", position, seed, position, seed, got, want)
		}
	}
}

func TestGet2dNoiseNegOneToOne(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		want := cgo.Get2dNoiseNegOneToOne(positionX, positionY, seed)
		got := Get2dNoiseNegOneToOne(positionX, positionY, seed)
		if got != want {
			t.Errorf("Get2dNoiseNegOneToOne(%d, %d, %d) != cgo.Get2dNoiseNegOneToOne(%d, %d, %d) => Got: %f | Want: %f", positionX, positionY, seed, positionX, positionY, seed, got, want)
		}
	}
}

func TestGet3dNoiseNegOneToOne(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		positionZ := (i + 2) * step
		want := cgo.Get3dNoiseNegOneToOne(positionX, positionY, positionZ, seed)
		got := Get3dNoiseNegOneToOne(positionX, positionY, positionZ, seed)
		if got != want {
			t.Errorf("Get3dNoiseNegOneToOne(%d, %d, %d, %d) != cgo.Get3dNoiseNegOneToOne(%d, %d, %d, %d) => Got: %f | Want: %f", positionX, positionY, positionZ, seed, positionX, positionY, positionZ, seed, got, want)
		}
	}
}

func TestGet4dNoiseNegOneToOne(t *testing.T) {
	for i := from; i < to; i++ {
		positionX := i * step
		positionY := (i + 1) * step
		positionZ := (i + 2) * step
		positionT := (i + 3) * step
		want := cgo.Get4dNoiseNegOneToOne(positionX, positionY, positionZ, positionT, seed)
		got := Get4dNoiseNegOneToOne(positionX, positionY, positionZ, positionT, seed)
		if got != want {
			t.Errorf("Get4dNoiseNegOneToOne(%d, %d, %d, %d, %d) != cgo.Get4dNoiseNegOneToOne(%d, %d, %d, %d, %d) => Got: %f | Want: %f", positionX, positionY, positionZ, positionT, seed, positionX, positionY, positionZ, positionT, seed, got, want)
		}
	}
}

// Benchmark the GO and the CGO versions for the Hash function
func BenchmarkSquirrelNoise5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquirrelNoise5(int32(i), seed)
	}
}
func BenchmarkGet1dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1dNoiseUint(int32(i), seed)
	}
}
func BenchmarkGet2dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get2dNoiseUint(int32(i), int32(i), seed)
	}
}
func BenchmarkGet3dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get3dNoiseUint(int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet4dNoiseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get4dNoiseUint(int32(i), int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet1dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1dNoiseZeroToOne(int32(i), seed)
	}
}
func BenchmarkGet2dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get2dNoiseZeroToOne(int32(i), int32(i), seed)
	}
}
func BenchmarkGet3dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get3dNoiseZeroToOne(int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet4dNoiseZeroToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get4dNoiseZeroToOne(int32(i), int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet1dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1dNoiseNegOneToOne(int32(i), seed)
	}
}
func BenchmarkGet2dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get2dNoiseNegOneToOne(int32(i), int32(i), seed)
	}
}
func BenchmarkGet3dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get3dNoiseNegOneToOne(int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet4dNoiseNegToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get4dNoiseNegOneToOne(int32(i), int32(i), int32(i), int32(i), seed)
	}
}

// CGO versions
func BenchmarkSquirrelNoise5CGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.SquirrelNoise5(int32(i), seed)
	}
}
func BenchmarkGet1dNoiseUintCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get1dNoiseUint(int32(i), seed)
	}
}
func BenchmarkGet2dNoiseUintCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get2dNoiseUint(int32(i), int32(i), seed)
	}
}
func BenchmarkGet3dNoiseUintCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get3dNoiseUint(int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet4dNoiseUintCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get4dNoiseUint(int32(i), int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet1dNoiseZeroToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get1dNoiseZeroToOne(int32(i), seed)
	}
}
func BenchmarkGet2dNoiseZeroToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get2dNoiseZeroToOne(int32(i), int32(i), seed)
	}
}
func BenchmarkGet3dNoiseZeroToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get3dNoiseZeroToOne(int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet4dNoiseZeroToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get4dNoiseZeroToOne(int32(i), int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet1dNoiseNegToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get1dNoiseNegOneToOne(int32(i), seed)
	}
}
func BenchmarkGet2dNoiseNegToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get2dNoiseNegOneToOne(int32(i), int32(i), seed)
	}
}
func BenchmarkGet3dNoiseNegToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get3dNoiseNegOneToOne(int32(i), int32(i), int32(i), seed)
	}
}
func BenchmarkGet4dNoiseNegToOneCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cgo.Get4dNoiseNegOneToOne(int32(i), int32(i), int32(i), int32(i), seed)
	}
}
