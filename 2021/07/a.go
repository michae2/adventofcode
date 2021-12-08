package main

import (
	"fmt"
	"sort"
)

// The median minimizes the sum of deviations. (Arithmetic mean minimizes the
// sum of squared deviations.)
func main() {
	var positions []int

	// Read positions from stdin.
	var pos int
	_, err := fmt.Scanf("%v,", &pos)
	for err == nil {
		positions = append(positions, pos)
		_, err = fmt.Scanf("%v,", &pos)
	}
	fmt.Scan(&pos)
	positions = append(positions, pos)

	// There are cheaper ways to find a median, but this is easy to type. :)
	// We don't have to deal with the even-number-of-samples case, since any
	// number between the two middle numbers (inclusive) will have the same sum
	// of deviations.
	sort.Ints(positions)
	median := positions[len(positions)/2]

	var total int
	for _, pos := range positions {
		if pos < median {
			total += median - pos
		} else {
			total += pos - median
		}
	}
	fmt.Println(total)
}
