# SquirrelNoise5-go - Squirrel's Raw Noise utilities (version 5)

`go get github.com/miltoncandelero/squirrelnoise5-go`

Original code written by Eiserloh Squirrel is made available under the Creative Commons attribution 3.0 license (CC-BY-3.0 US)

This GO port is licensed under the MIT Licence.

These noise functions were written by Squirrel Eiserloh as a cheap and simple substitute for the [sometimes awful]
bit-noise sample code functions commonly found on the web, many of which are hugely biased or terribly patterned, e.g.
having bits which are on (or off) 75% or even 100% of the time (or are excessively overkill/slow for our needs, such as
MD5 or SHA).

Peter Schmidt-Nielsen identified a weakness in the SquirrelNoise3 code originally used in the GDC 2017 talk, "Noise-based RNG". 
Version 5 avoids a noise repetition found in version 3 at extremely high position values caused by a lack of influence by some 
of the high input bits onto some of the low output bits.

The following functions are all based on a simple bit-noise hash function which returns an unsigned integer containing
32 reasonably-well-scrambled bits, based on a given (signed) integer input parameter (position/index) and [optional] seed. 
Kind of like looking up a value in an infinitely large [non-existent] table of previously rolled random numbers.

These functions are deterministic and random-access / order-independent (i.e. state-free), so they are particularly
well-suited for use in smoothed/fractal/simplex/Perlin noise functions and out-of-order (or or-demand) procedural
content generation (i.e. that mountain village is the same whether you generated it first or last, ahead of time or just
now).

The N-dimensional variations simply hash their multidimensional coordinates down to a single 32-bit index and then
proceed as usual, so while results are not unique they should
(hopefully) not seem locally predictable or repetitive.

## This code should NOT be considered cryptographically secure!

The idea behind this port is that I needed to have a stable and fast random algorithm that I could run in both my
games and the backend for them and have them generate the same numbers to prevent petty cheating.

This is not meant to be used as a safe random number for cryptography. This is meant for videogames and fun.

---

### Raw pseudorandom noise functions (random-access / deterministic). Basis of all other noise.

```go
func Get1dNoiseUint( index int32, seed uint32 ) uint32
func Get2dNoiseUint( indexX int32, indexY int32, seed uint32 ) uint32
func Get3dNoiseUint( indexX int32, indexY int32, indexZ int32, seed uint32 ) uint32
func Get4dNoiseUint( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 ) uint32
```

---

### Same functions, mapped to floats in [0,1] for convenience.

```go
func Get1dNoiseZeroToOne( index int32, seed uint32 ) float32
func Get2dNoiseZeroToOne( indexX int32, indexY int32, seed uint32 ) float32
func Get3dNoiseZeroToOne( indexX int32, indexY int32, indexZ int32, seed uint32 ) float32
func Get4dNoiseZeroToOne( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 ) float32
```

---

### Same functions, mapped to floats in [-1,1] for convenience.

```go
func Get1dNoiseNegOneToOne( index int32, seed uint32 ) float32
func Get2dNoiseNegOneToOne( indexX int32, indexY int32, seed uint32 ) float32
func Get3dNoiseNegOneToOne( indexX int32, indexY int32, indexZ int32, seed uint32 ) float32
func Get4dNoiseNegOneToOne( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 ) float32
```

---

## Benchmark and Testing
Since I wanted to be _very_ sure that the same numbers were being generated from this go port and the original C++ 
version, I added a `cgo` package that runs the original code and is used for testing.

Since I already had that, I could run benchmarks comparing them. Keep in mind that the poor performance in the CGO 
implementation is most likely due to the bridge between go and C.

Benchmark ran on my laptop with a AMD Ryzen 9 5900HS and 40GB of ram running windows using Zig as my C compiler.

| Method              | Go Native Implementation | CGO Implementation |
|---------------------|--------------------------|--------------------|
| SquirrelNoise5      | 0.2281 ns/op             | 38.50 ns/op        |
| Get1dNoiseUint      | 0.2304 ns/op             | 38.83 ns/op        |
| Get2dNoiseUint      | 0.2293 ns/op             | 39.62 ns/op        |
| Get3dNoiseUint      | 0.2281 ns/op             | 40.67 ns/op        |
| Get4dNoiseUint      | 0.2253 ns/op             | 40.91 ns/op        |
| Get1dNoiseZeroToOne | 0.2305 ns/op             | 39.64 ns/op        |
| Get2dNoiseZeroToOne | 0.2281 ns/op             | 40.48 ns/op        |
| Get3dNoiseZeroToOne | 0.2278 ns/op             | 40.64 ns/op        |
| Get4dNoiseZeroToOne | 2.601 ns/op              | 40.96 ns/op        |
| Get1dNoiseNegToOne  | 0.2283 ns/op             | 40.52 ns/op        |
| Get2dNoiseNegToOne  | 0.2299 ns/op             | 40.08 ns/op        |
| Get3dNoiseNegToOne  | 0.2277 ns/op             | 41.12 ns/op        |
| Get4dNoiseNegToOne  | 2.595 ns/op              | 40.79 ns/op        |

Not really sure why the 4D float version is an order of magnitude slower than its peers. 
I couldn't see anything obviously wrong with the implementation, and I barely use the 4D version so I don't care that much.