package main

import "fmt"

func main() {
	var depth, prev, increases int
	_, err := fmt.Scan(&depth)
	prev = depth
	for ; err == nil; _, err = fmt.Scan(&depth) {
		if depth > prev {
			increases++
		}
		prev = depth
	}
	fmt.Println(increases)
}
