package main

import "fmt"

func main() {
	var depth [3]int
	var d, prev, increases int
	_, err := fmt.Scan(&depth[0], &depth[1], &depth[2])
	prev = depth[0] + depth[1] + depth[2]
	for ; err == nil; _, err = fmt.Scan(&depth[d%3]) {
		d++
		sum := depth[0] + depth[1] + depth[2]
		if sum > prev {
			increases++
		}
		prev = sum
	}
	fmt.Println(increases)
}
