package main

import "fmt"

func main() {
	var amount, horizontal, depth, aim int
	var command string
	_, err := fmt.Scan(&command, &amount)
	for ; err == nil; _, err = fmt.Scan(&command, &amount) {
		switch command {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	fmt.Println(horizontal * depth)
}
