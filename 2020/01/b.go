package main

import "fmt"

func main() {
	var e int
	var es []int
	tab := make(map[int]int)
	for _, err := fmt.Scan(&e); err == nil; _, err = fmt.Scan(&e) {
		d := 2020 - e
		f, ok := tab[d]
		if ok {
			fmt.Println(f * e)
			break
		}
		for _, e0 := range es {
			tab[e0+e] = e0 * e
		}
		es = append(es, e)
	}
}
