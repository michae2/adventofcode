package main

import (
	"fmt"
	"math"
)

type dot struct {
	x, y int
}

func main() {
	var dots []dot
	var x, y int
	_, err := fmt.Scanf("%d,%d\n", &x, &y)
	for err == nil {
		dots = append(dots, dot{x, y})
		_, err = fmt.Scanf("%d,%d\n", &x, &y)
	}

	var axis rune
	var z int
	_, err = fmt.Scanf("fold along %c=%d", &axis, &z)
	for err == nil {
		if axis == 'x' {
			for i := range dots {
				if dots[i].x > z {
					dots[i].x = 2*z - dots[i].x
				}
			}
		} else {
			for i := range dots {
				if dots[i].y > z {
					dots[i].y = 2*z - dots[i].y
				}
			}
		}
		_, err = fmt.Scanf("fold along %c=%d", &axis, &z)
	}

	finalDots := make(map[dot]struct{}, len(dots))
	min := dot{math.MaxInt, math.MaxInt}
	max := dot{math.MinInt, math.MinInt}
	for _, d := range dots {
		finalDots[d] = struct{}{}
		if d.x < min.x {
			min.x = d.x
		}
		if d.y < min.y {
			min.y = d.y
		}
		if d.x > max.x {
			max.x = d.x
		}
		if d.y > max.y {
			max.y = d.y
		}
	}

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if _, ok := finalDots[dot{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
