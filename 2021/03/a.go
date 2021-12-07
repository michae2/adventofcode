package main

import "fmt"

func main() {
	var binary string
	var counts [12]int
	for _, err := fmt.Scan(&binary); err == nil; _, err = fmt.Scan(&binary) {
		for i := 0; i < 12; i++ {
			if binary[i] == '1' {
				counts[i]++
			} else {
				counts[i]--
			}
		}
	}
	var gamma, epsilon int
	for i := 0; i < 12; i++ {
		if counts[i] > 0 {
			gamma += 1 << (11 - i)
		} else {
			epsilon += 1 << (11 - i)
		}
	}
	fmt.Println(gamma * epsilon)
}
