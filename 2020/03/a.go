package main

import "fmt"

func main() {
	var n, trees int
	var line string
	for _, err := fmt.Scanln(&line); err == nil; _, err = fmt.Scanln(&line) {
		if line[n*3%len(line)] == '#' {
			trees++
		}
		n++
	}
	fmt.Println(trees)
}
