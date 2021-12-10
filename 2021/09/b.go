package main

import "fmt"

type basin struct {
	size  int
	merge *basin
}

// As we scan the heightmap line-by-line we track the basin for each location in
// the previous line (north) and the current line. First we continue basins from
// the north, then we merge together basins that touch each other.
func main() {
	var basins []basin
	var line string
	_, err := fmt.Scan(&line)
	northBasins := make([]*basin, len(line))
	curBasins := make([]*basin, len(line))
	for err == nil {
		// First pass carries basins down from the north, or creates new basins.
		for i := 0; i < len(line); i++ {
			// The exact definition of a basin had something to do with flowing
			// down to a single low point, but I think this less precise
			// definition of "contiguous heights < 9" will probably also work
			// and is easier to implement.
			if line[i] < '9' {
				if northBasins[i] == nil {
					basins = append(basins, basin{})
					curBasins[i] = &basins[len(basins)-1]
				} else {
					curBasins[i] = northBasins[i]
				}
				curBasins[i].size++
			} else {
				curBasins[i] = nil
			}
		}
		// Second pass merges basins that touch each other.
		for i := 1; i < len(line); i++ {
			if curBasins[i] != nil {
				for curBasins[i].merge != nil {
					curBasins[i] = curBasins[i].merge
				}
				west := curBasins[i-1]
				if west != nil && west != curBasins[i] {
					curBasins[i].merge = west
					west.size += curBasins[i].size
					curBasins[i] = west
				}
			}
		}
		northBasins, curBasins = curBasins, northBasins
		_, err = fmt.Scan(&line)
	}

	var top [3]int
	for _, b := range basins {
		if b.merge != nil {
			continue
		}
		// Could do this with a heap if K were larger, but 3 is very small.
		if b.size > top[0] {
			top[2] = top[1]
			top[1] = top[0]
			top[0] = b.size
		} else if b.size > top[1] {
			top[2] = top[1]
			top[1] = b.size
		} else if b.size > top[2] {
			top[2] = b.size
		}
	}
	fmt.Println(top[0] * top[1] * top[2])
}
