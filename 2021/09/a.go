package main

import (
	"fmt"
	"strings"
)

// Adding sentinels on the borders simplifies the main loop.
func nextLine(line *string) bool {
	if len(*line) > 0 && (*line)[0] == 'Z' {
		return false
	}
	var heights string
	_, err := fmt.Scan(&heights)
	if err != nil {
		*line = "Z" + strings.Repeat("9", len(*line)-1)
		return true
	}
	*line = "9" + heights + "9"
	return true
}

func main() {
	var sum int
	var north, cur, south string
	nextLine(&cur)
	north = strings.Repeat("9", len(cur))
	for nextLine(&south) {
		for i := 1; i < len(cur)-1; i++ {
			n := north[i]
			w := cur[i-1]
			h := cur[i]
			e := cur[i+1]
			s := south[i]
			if h < n && h < w && h < e && h < s {
				sum += int(h-'0') + 1
			}
		}
		north = cur
		cur = south
	}
	fmt.Println(sum)
}
