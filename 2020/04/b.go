package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const pattern = `byr:(\d{4})\s|` +
	`iyr:(\d{4})\s|` +
	`eyr:(\d{4})\s|` +
	`hgt:(\d+cm|\d+in)\s|` +
	`hcl:(#[0-9a-f]{6})\s|` +
	`ecl:(amb|blu|brn|gry|grn|hzl|oth)\s|` +
	`pid:(\d{9})\s|` +
	`cid:(\d+)\s`

var bounds = [][]int{{1920, 2002}, {2010, 2020}, {2020, 2030}}

func indexes(match []int, f int) (i, j int) {
	g := f + 1
	return match[2*g], match[2*g+1]
}

func validInt(s string, lo, hi int) bool {
	i, err := strconv.Atoi(s)
	return err == nil && lo <= i && i <= hi
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
		for _, match := range re.FindAllStringSubmatchIndex(line+"\n", -1) {
			for f, b := range bounds {
				if i, j := indexes(match, f); i >= 0 && j >= 0 &&
					validInt(line[i:j], b[0], b[1]) {
					fields |= 1 << f
				}
			}
			if i, j := indexes(match, 3); i >= 0 && j >= 0 &&
				((line[j-2:j] == "cm" && validInt(line[i:j-2], 150, 193)) ||
					(line[j-2:j] == "in" && validInt(line[i:j-2], 59, 76))) {
				fields |= 1 << 3
			}
			for f := 4; f < 8; f++ {
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
