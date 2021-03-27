package main

import "fmt"

func main() {
	var valid int
	for {
		var a, b int
		var ch, pw string
		_, err := fmt.Scanf("%v-%v %1v: %v\n", &a, &b, &ch, &pw)
		if err != nil {
			break
		}
		if (a <= len(pw) && pw[a-1] == ch[0]) != (b <= len(pw) && pw[b-1] == ch[0]) {
			valid++
		}
	}
	fmt.Println(valid)
}
