package main

import (
	"fmt"
	"sort"
)

func main() {
	var scores []int
	var line string
NextLine:
	for _, err := fmt.Scan(&line); err == nil; _, err = fmt.Scan(&line) {
		var stack []byte
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '(', '[', '{', '<':
				stack = append(stack, line[i])
			case ')':
				if stack[len(stack)-1] != '(' {
					continue NextLine
				}
				stack = stack[:len(stack)-1]
			case ']':
				if stack[len(stack)-1] != '[' {
					continue NextLine
				}
				stack = stack[:len(stack)-1]
			case '}':
				if stack[len(stack)-1] != '{' {
					continue NextLine
				}
				stack = stack[:len(stack)-1]
			case '>':
				if stack[len(stack)-1] != '<' {
					continue NextLine
				}
				stack = stack[:len(stack)-1]
			}
		}
		var score int
		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			switch stack[i] {
			case '(':
				score += 1
			case '[':
				score += 2
			case '{':
				score += 3
			case '<':
				score += 4
			}
		}
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
