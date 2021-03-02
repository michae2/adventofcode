package main

import "fmt"

func main() {
	var e int
	es := make(map[int]bool)
	for _, err := fmt.Scan(&e); err == nil; _, err = fmt.Scan(&e) {
		d := 2020 - e
		if es[d] {
			fmt.Println(d * e)
			break
		}
		es[e] = true
	}
}
