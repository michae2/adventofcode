package main

import "fmt"

func main() {
	var amount, horizontal, depth int
	var command string
	_, err := fmt.Scan(&command, &amount)
	for ; err == nil; _, err = fmt.Scan(&command, &amount) {
		switch command {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}
	fmt.Println(horizontal * depth)
}
