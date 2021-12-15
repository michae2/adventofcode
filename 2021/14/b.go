package main

import (
	"fmt"
	"math"
)

type polymer struct {
	// The trick here is to make left and right subtrees overlap by one
	// element. Then (a) we don't need special logic for the boundary between
	// left and right, and (b) the input and output of reactions can both be
	// represented by the same type.
	left, right *polymer
	// Leftmost and rightmost elements in the tree. In the base case (nil left
	// and right subtrees) these are the only elements in the tree.
	lElem, rElem byte
}

func react(rules map[polymer]*polymer, tree polymer) *polymer {
	if p, ok := rules[tree]; ok {
		return p
	}
	p := &polymer{
		react(rules, *tree.left),
		react(rules, *tree.right),
		tree.lElem, tree.rElem,
	}
	rules[tree] = p
	return p
}

func count(counts map[polymer][26]int64, tree polymer) [26]int64 {
	if c, ok := counts[tree]; ok {
		return c
	}
	var c [26]int64
	if tree.left == nil {
		c[tree.lElem-'A']++
		c[tree.rElem-'A']++
	} else {
		lc, rc := count(counts, *tree.left), count(counts, *tree.right)
		for i := range c {
			c[i] = lc[i] + rc[i]
		}
		if tree.left.rElem != tree.right.lElem {
			panic("Left and right subtrees did not overlap??")
		}
		// Account for the overlap.
		c[tree.left.rElem-'A']--
	}
	counts[tree] = c
	return c
}

// For 40 generations, we have to use a smarter algorithm to avoid exponential
// time and space blowup. Dynamic programming works well here, because
// subsections of the polymer are often repeated. This algorithm resembles
// Hashlife.
func main() {
	var template []byte
	fmt.Scanf("%s\n\n", &template)

	// Turn the template into a balanced tree of polymer structs.
	trees := make([]*polymer, len(template)-1)
	for i := 0; i < len(template)-1; i++ {
		trees[i] = &polymer{nil, nil, template[i], template[i+1]}
	}
	for l := len(trees); l > 1; l = (l + 1) / 2 {
		for i := 0; i < l/2; i++ {
			l, r := trees[2*i], trees[2*i+1]
			trees[i] = &polymer{l, r, l.lElem, r.rElem}
		}
		if l%2 == 1 {
			trees[l/2] = trees[l-1]
		}
	}
	tree := trees[0]

	// Seed the memoization table with the base-case rules.
	rules := make(map[polymer]*polymer)
	var pair []byte
	var elem byte
	_, err := fmt.Scanf("%s -> %c\n", &pair, &elem)
	for err == nil {
		rules[polymer{nil, nil, pair[0], pair[1]}] = &polymer{
			&polymer{nil, nil, pair[0], elem},
			&polymer{nil, nil, elem, pair[1]},
			pair[0], pair[1],
		}
		_, err = fmt.Scanf("%s -> %c\n", &pair, &elem)
	}

	for step := 0; step < 40; step++ {
		tree = react(rules, *tree)
	}

	counts := count(make(map[polymer][26]int64), *tree)

	var min, max int64 = math.MaxInt64, math.MinInt64
	for _, count := range counts {
		if count == 0 {
			continue
		}
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	fmt.Println(max - min)
}
