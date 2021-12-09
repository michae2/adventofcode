package main

import "fmt"

func main() {
	var count int
	var patterns [10]string
	var values [4]string
	var err error
	for err == nil {
		// First time through loop, patterns and values are all zero-length.

		for _, value := range values {
			switch len(value) {
			case 2, 3, 4, 7:
				count++
			}
		}
		_, err = fmt.Scanf(
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n",
			&patterns[0], &patterns[1], &patterns[2], &patterns[3],
			&patterns[4], &patterns[5], &patterns[6], &patterns[7],
			&patterns[8], &patterns[9],
			&values[0], &values[1], &values[2], &values[3],
		)
	}
	fmt.Println(count)
}
