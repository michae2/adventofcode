package main

import (
	"fmt"
	"strconv"
)

// Use a Quicksort-like algorithm to recursively find the partition with the
// most (or least) numbers.
func quickFind(numbers []uint, most bool) uint {
	lo, hi := 0, len(numbers)
	b := 11
	var p uint = 1 << b
	for lo+1 < hi {
		i, j := lo, hi
		for i < j {
			// Numbers >= the pivot will have a 1 in the (12-b)th position and
			// must be swapped to the end of the current subset.
			if numbers[i] >= p {
				numbers[i], numbers[j-1] = numbers[j-1], numbers[i]
				j--
			} else {
				i++
			}
		}
		b--
		if (most && i-lo > hi-j) || (!most && i-lo <= hi-j) {
			// Keep the 0s.
			hi = i
			p -= 1 << b
		} else {
			// Keep the 1s.
			lo = j
			p += 1 << b
		}
	}
	return numbers[lo]
}

func main() {
	var binary string
	var numbers []uint
	for _, err := fmt.Scan(&binary); err == nil; _, err = fmt.Scan(&binary) {
		num, _ := strconv.ParseUint(binary, 2, 32)
		numbers = append(numbers, uint(num))
	}
	oxy := quickFind(numbers, true)
	co2 := quickFind(numbers, false)
	fmt.Println(oxy * co2)
}
