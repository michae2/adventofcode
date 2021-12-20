package main

import (
	"fmt"
	"math"
)

func triangularRoot(x float64) float64 {
	return (math.Sqrt(8*x+1) - 1) / 2
}

// When we only have to consider one direction of travel, the lowest possible
// velocity that could put the probe in the target range is either simply the
// position itself (if the lowest position is negative) or the minimum velocity
// that could eventually reach the position (if the lowest position is
// positive).
func minVelocity(z int) int {
	if z <= 0 {
		return z
	}
	return int(math.Ceil(triangularRoot(float64(z))))
}

// The highest possible y velocity will always involve an arc that "lands" in
// the target range. The arc is symmetrical, so it will always have a y position
// back at 0 after some n number of steps. The highest possible arc will hit the
// furthest y position in the target range from 0 either at step n+1 if this
// furthest position is negative or n-1 if this furthest position is positive.
func maxYVelocity(y1, y2 int) int {
	v1, v2 := y1, y2
	if y1 < 0 {
		v1 = -1 - y1
	}
	if y2 < 0 {
		v2 = -1 - y2
	}
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

// Because the differential equations for x and y are independent of each other,
// we can calculate all possible launch velocities for x and y separately and
// then join them together. I'm doing this in a brute force way, but fortunately
// the bounds are small enough that it doesn't take long.
func main() {
	var x1, x2, y1, y2 int
	fmt.Scanf("target area: x=%d..%d, y=%d..%d\n", &x1, &x2, &y1, &y2)

	if x1 > x2 || y1 > y2 {
		panic("bad args")
	}

	// All velocities that land in the target range, keyed by n number of steps.
	// There could be multiple entries for a single velocity if the probe stays
	// in the target range for multiple steps.
	yVelocities := make(map[int][]int)
	xVelocities := make(map[int][]int)

	// Some x velocities could stay in the target range for infinite steps, so
	// first use the y velocities to bound n.
	maxN := 0

	// Try all possible y launch velocities.
	minYV := minVelocity(y1)
	maxYV := maxYVelocity(y1, y2)

	for launchV := minYV; launchV <= maxYV; launchV++ {
		// Step through until we drop below y1.
		for n, y, v, above := 0, 0, launchV, false; !above || y >= y1; n++ {
			if y >= y1 && y <= y2 {
				yVelocities[n] = append(yVelocities[n], launchV)
				if n > maxN {
					maxN = n
				}
			}
			if y >= y1 {
				above = true
			}
			y += v
			v--
		}
	}

	// Try all possible x launch velocities.
	minXV := minVelocity(x1)
	maxXV := -minVelocity(-x2)

	// Negative x launch velocities.
	for launchV := minXV; launchV <= maxXV && launchV < 0; launchV++ {
		// Step until we pass below x1.
		for n, x, v := 0, 0, launchV; n <= maxN && x >= x1; n++ {
			if x <= x2 {
				xVelocities[n] = append(xVelocities[n], launchV)
			}
			x += v
			if v < 0 {
				v++
			}
		}
	}

	// Positive x launch velocities.
	if minXV < 0 {
		minXV = 0
	}
	for launchV := minXV; launchV <= maxXV; launchV++ {
		// Step until we pass above x2.
		for n, x, v := 0, 0, launchV; n <= maxN && x <= x2; n++ {
			if x >= x1 {
				xVelocities[n] = append(xVelocities[n], launchV)
			}
			x += v
			if v > 0 {
				v--
			}
		}
	}

	// Distinct velocity pairs.
	velocities := make(map[[2]int]struct{})

	// Join x and y velocities on number of steps.
	for n, ys := range yVelocities {
		if xs, ok := xVelocities[n]; ok {
			for _, y := range ys {
				for _, x := range xs {
					velocities[[2]int{x, y}] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(velocities))
}
