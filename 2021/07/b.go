package main

import (
	"fmt"
	"math"
)

func sumDeviations(positions []float64, mean float64) (sum float64) {
	for _, position := range positions {
		deviation := math.Abs(position - mean)
		// The nth triangular number can be calculated with n(n+1)/2.
		cost := deviation * (deviation + 1) / 2
		sum += cost
	}
	return
}

// The arithmetic mean minimizes the sum of squared deviations. In this case
// we're trying to minimize the sum of triangular deviations, not squared
// deviations, but triangular numbers are roughly proportional to square numbers
// (roughly square numbers / 2), so the arithmetic mean will also minimize the
// sum of triangular deviations.
func main() {
	var positions []float64

	// Read positions from stdin. Using floats for convenience.
	var pos float64
	_, err := fmt.Scanf("%v,", &pos)
	for err == nil {
		positions = append(positions, pos)
		_, err = fmt.Scanf("%v,", &pos)
	}
	fmt.Scan(&pos)
	positions = append(positions, pos)

	// Calculate the mean.
	var sum float64
	for _, pos := range positions {
		sum += pos
	}
	mean := sum / float64(len(positions))

	// The optimal position will be one of the positions on either side of the
	// mean, so sum the deviations for both and pick the smaller.
	totalF := sumDeviations(positions, math.Floor(mean))
	totalC := sumDeviations(positions, math.Ceil(mean))
	fmt.Printf("%.f\n", math.Min(totalF, totalC))
}
