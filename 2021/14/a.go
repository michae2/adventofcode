package main

import (
	"fmt"
	"math"
)

func react(rules map[[2]byte]byte, polymer []byte) []byte {
	p := make([]byte, 0, len(polymer)*2-1)
	p = append(p, polymer[0])
	for i := 1; i < len(polymer); i++ {
		elem, ok := rules[[2]byte{polymer[i-1], polymer[i]}]
		if !ok {
			panic("missing rule")
		}
		p = append(p, elem, polymer[i])
	}
	return p
}

func main() {
	var polymer []byte
	fmt.Scanf("%s\n\n", &polymer)

	rules := make(map[[2]byte]byte)
	var pair []byte
	var elem byte
	_, err := fmt.Scanf("%s -> %c\n", &pair, &elem)
	for err == nil {
		rules[[2]byte{pair[0], pair[1]}] = elem
		_, err = fmt.Scanf("%s -> %c\n", &pair, &elem)
	}

	for step := 0; step < 10; step++ {
		polymer = react(rules, polymer)
	}

	var counts [26]int
	for _, elem := range polymer {
		counts[elem-'A']++
	}

	min, max := math.MaxInt, math.MinInt
	for _, count := range counts {
		if count == 0 {
			continue
		}
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	fmt.Println(max - min)
}
