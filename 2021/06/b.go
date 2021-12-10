package main

import "fmt"

func main() {
	// Bucket by timer instead of tracking each individual lanternfish.
	var fish [9]int64
	var timer int64
	_, err := fmt.Scanf("%v,", &timer)
	for err == nil {
		fish[timer]++
		_, err = fmt.Scanf("%v,", &timer)
	}
	fmt.Scan(&timer)
	fish[timer]++

	for i := 0; i < 256; i++ {
		births := fish[0]
		for t := 1; t < 9; t++ {
			fish[t-1] = fish[t]
		}
		fish[6] += births
		fish[8] = births
	}

	var count int64
	for _, n := range fish {
		count += n
	}
	fmt.Println(count)
}
