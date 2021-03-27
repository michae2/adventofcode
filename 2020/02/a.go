package main

import (
	"fmt"
	"strings"
)

func main() {
	var valid int
	for {
		var lo, hi int
		var ch, pw string
		_, err := fmt.Scanf("%v-%v %1v: %v\n", &lo, &hi, &ch, &pw)
		if err != nil {
			break
		}
		n := strings.Count(pw, ch)
		if lo <= n && n <= hi {
			valid++
		}
	}
	fmt.Println(valid)
}
