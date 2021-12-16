package main

import (
	"container/heap"
	"fmt"
)

const (
	width  = 100
	height = 100
)

type point struct {
	x, y int
}

func (p point) neighbors() [4]point {
	return [4]point{
		point{p.x, p.y + 1},
		point{p.x + 1, p.y},
		point{p.x, p.y - 1},
		point{p.x - 1, p.y},
	}
}

type position struct {
	point
	risk int
	// So that we can fix the heap after updating total risk to this position.
	index int
}

type pQueue []*position

// In Go we have to implement heap.Interface to get a priority queue.
var _ heap.Interface = &pQueue{}

func (pq pQueue) Len() int {
	return len(pq)
}

func (pq pQueue) Less(i, j int) bool {
	iCost := pq[i].risk + (width - pq[i].x - 1) + (height - pq[i].y - 1)
	jCost := pq[j].risk + (width - pq[j].x - 1) + (height - pq[j].y - 1)
	return iCost < jCost
}

func (pq pQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *pQueue) Push(x interface{}) {
	pos := x.(*position)
	pos.index = len(*pq)
	*pq = append(*pq, pos)
}

func (pq *pQueue) Pop() interface{} {
	pos := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return pos
}

// Use A* search.
func main() {
	var risks [height][width]byte
	for y := 0; y < height; y++ {
		var line []byte
		fmt.Scan(&line)
		copy(risks[y][:], line)
	}

	var visited [height][width]bool
	var pq pQueue
	open := make(map[point]*position)

	end := point{width - 1, height - 1}
	current := &position{}
	visited[0][0] = true
	for current.point != end {
		for _, neighbor := range current.neighbors() {
			if neighbor.x < 0 || neighbor.x > end.x ||
				neighbor.y < 0 || neighbor.y > end.y ||
				visited[neighbor.y][neighbor.x] {
				continue
			}
			risk := int(risks[neighbor.y][neighbor.x]-'0') + current.risk
			if pos, ok := open[neighbor]; ok {
				if risk < pos.risk {
					pos.risk = risk
					heap.Fix(&pq, pos.index)
				}
			} else {
				pos := &position{point: neighbor, risk: risk}
				heap.Push(&pq, pos)
				open[neighbor] = pos
			}
		}
		current = heap.Pop(&pq).(*position)
		delete(open, current.point)
		visited[current.y][current.x] = true
	}
	fmt.Println(current.risk)
}
