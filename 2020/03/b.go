package main

import "fmt"

func treeCheck(grid []string, x, y int) (trees int) {
	for n, line := range grid {
		if n%y == 0 && line[n/y*x%len(line)] == '#' {
			trees++
		}
	}
	return
}

func main() {
	var grid []string
	var line string
	for _, err := fmt.Scanln(&line); err == nil; _, err = fmt.Scanln(&line) {
		grid = append(grid, line)
	}
	fmt.Println(treeCheck(grid, 1, 1) *
		treeCheck(grid, 3, 1) *
		treeCheck(grid, 5, 1) *
		treeCheck(grid, 7, 1) *
		treeCheck(grid, 1, 2))
}
