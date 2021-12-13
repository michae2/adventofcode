package main

import "fmt"

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
	fmt.Scanf("fold along %c=%d", &axis, &z)

	foldedDots := make(map[dot]struct{}, len(dots))
	if axis == 'x' {
		for _, d := range dots {
			if d.x > z {
				d.x = 2*z - d.x
			}
			foldedDots[d] = struct{}{}
		}
	} else {
		for _, d := range dots {
			if d.y > z {
				d.y = 2*z - d.y
			}
			foldedDots[d] = struct{}{}
		}
	}
	fmt.Println(len(foldedDots))
}
