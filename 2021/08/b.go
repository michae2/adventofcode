package main

import "fmt"

// It's a little easier to perform set operations on bits, so we parse a signal
// into a bitset when necessary.
func parse(pattern string) (code int) {
	for _, wire := range pattern {
		code |= 1 << (wire - 'a')
	}
	return
}

// As stated in the problem, counting the number of lit wires will distinguish
// digits 1, 4, 7, 8 without further logic. But digits 2, 3, 5 all have five lit
// wires; and digits 0, 6, 9 all have six lit wires. For these we need
// additional logic. Here's what I've devised:
//
// Five lit wires (2/3/5):
// If the intersection of the signal and 1 (cf) == 1 (cf), it is 3,
// else if the intersection of the signal and 4 - 1 (bd) == 4 - 1 (bd), it is 5,
// else it is 2.
//
// Six lit wires (0/6/9):
// If the intersection of the signal and 4 (bcdf) == 4 (bcdf), it is 9,
// else if the intersection of the signal and 1 (cf) == 1 (cf), it is 0,
// else it is 6.
func main() {
	var sum int
	var patterns [10]string
	var values [4]string
	var err error
	for err == nil {
		// First time through loop, patterns and values are all zero-length.

		// Determine a few bit patterns we need to use to distinguish between
		// digits with the same length.
		var bd, cf, bcdf int
		for _, pattern := range patterns {
			switch len(pattern) {
			case 2:
				cf = parse(pattern)
			case 4:
				bcdf = parse(pattern)
			}
		}
		bd = bcdf - cf

		// Now decode each scrambled output value.
		factor := 1000
		for _, value := range values {
			var digit int
			switch len(value) {
			case 2:
				digit = 1
			case 3:
				digit = 7
			case 4:
				digit = 4
			case 5:
				code := parse(value)
				if code&cf == cf {
					digit = 3
				} else if code&bd == bd {
					digit = 5
				} else {
					digit = 2
				}
			case 6:
				code := parse(value)
				if code&bcdf == bcdf {
					digit = 9
				} else if code&cf == cf {
					digit = 0
				} else {
					digit = 6
				}
			case 7:
				digit = 8
			}
			sum += factor * digit
			factor /= 10
		}
		_, err = fmt.Scanf(
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n",
			&patterns[0], &patterns[1], &patterns[2], &patterns[3],
			&patterns[4], &patterns[5], &patterns[6], &patterns[7],
			&patterns[8], &patterns[9],
			&values[0], &values[1], &values[2], &values[3],
		)
	}
	fmt.Println(sum)
}
