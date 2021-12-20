package main

import (
	"fmt"
	"math"
)

func triangular(n int) int {
	return n * (n + 1) / 2
}

func triangularRoot(x float64) float64 {
	return (math.Sqrt(8*x+1) - 1) / 2
}

// The highest possible y velocity will always involve an arc that "lands" in
// the target range. The arc is symmetrical, so it will always have a y position
// back at 0 after some n number of steps. The highest possible arc will hit the
// furthest y position in the target range from 0 either at step n+1 if this
// furthest position is negative or n-1 if this furthest position is positive.
func maxYVelocity(y1, y2 int) (v, n int) {
	v1, v2 := y1, y2
	n1, n2 := 2*v1, 2*v2
	if y1 < 0 {
		v1 = -1 - y1
		n1 = -2 * y1
	}
	if y2 < 0 {
		v2 = -1 - y2
		n2 = -2 * y2
	}
	if v1 > v2 {
		return v1, n1
	} else {
		return v2, n2
	}
}

// Making some assumptions about how the target range was picked turns this into
// a direct calculation rather than an iterative / numerical approach.
func main() {
	var x1, x2, y1, y2 int
	fmt.Scanf("target area: x=%d..%d, y=%d..%d\n", &x1, &x2, &y1, &y2)

	if x1 > x2 || y1 > y2 {
		panic("bad args")
	}

	// This is much simpler if we assume that (a) the x velocity of the probe
	// will hit 0 when inside the target range after k steps, thus allowing any
	// answer that takes >= k steps and (b) the highest y position will be
	// achieved with an arc that "lands" in the target range after >= k steps.
	maxYV, maxYN := maxYVelocity(y1, y2)
	maxY := triangular(maxYV)

	// Check that our assumptions hold.
	minK := 0
	if x1 > 0 {
		minK = int(math.Ceil(triangularRoot(float64(x1))))
	} else if x2 < 0 {
		minK = int(math.Floor(triangularRoot(float64(x2))))
	}

	minKX := triangular(minK)
	if minKX < x1 || minKX > x2 {
		panic("no x velocity reaches 0 inside target range")
	}

	if maxYN < minK {
		panic("max y velocity reaches target range before k steps")
	}

	fmt.Println(maxY)
}
