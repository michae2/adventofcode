package main

import "fmt"

func main() {
	var score int
	var line string
	_, err := fmt.Scan(&line)
	for err == nil {
		var stack []byte
	Line:
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '(', '[', '{', '<':
				stack = append(stack, line[i])
			case ')':
				if stack[len(stack)-1] != '(' {
					score += 3
					break Line
				}
				stack = stack[:len(stack)-1]
			case ']':
				if stack[len(stack)-1] != '[' {
					score += 57
					break Line
				}
				stack = stack[:len(stack)-1]
			case '}':
				if stack[len(stack)-1] != '{' {
					score += 1197
					break Line
				}
				stack = stack[:len(stack)-1]
			case '>':
				if stack[len(stack)-1] != '<' {
					score += 25137
					break Line
				}
				stack = stack[:len(stack)-1]
			}
		}
		_, err = fmt.Scan(&line)
	}
	fmt.Println(score)
}
