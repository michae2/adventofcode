package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const pattern = `(byr:\S*)|(iyr:\S*)|(eyr:\S*)|(hgt:\S*)|` +
	`(hcl:\S*)|(ecl:\S*)|(pid:\S*)|(cid:\S*)`

func indexes(match []int, f int) (i, j int) {
	g := f + 1
	return match[2*g], match[2*g+1]
}

func main() {
	var valid, fields int
	re := regexp.MustCompile(pattern)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if fields&0x7f == 0x7f {
				valid++
			}
			fields = 0
			continue
		}
		for _, match := range re.FindAllStringSubmatchIndex(line, -1) {
			for f := 0; f < 8; f++ {
				if i, j := indexes(match, f); i >= 0 && j >= 0 {
					fields |= 1 << f
				}
			}
		}
	}
	if fields&0x7f == 0x7f {
		valid++
	}
	fmt.Println(valid)
}
