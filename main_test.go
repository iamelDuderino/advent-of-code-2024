package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMerge(t *testing.T) {
	a := big.NewInt(1011500535579593609)
	b := big.NewInt(613)
	c := merge(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, c)
	fmt.Println(c.Int64())
}

func TestD7P2(t *testing.T) {

	// Problem
	var p = &day7problem{
		answer: 605, // 1 || 2 + 3 * 4 || 5
		nums:   []int{1, 2, 3, 4, 5},
	}

	// Answers
	_, a := bigIntRecursion(convertToBigInts(p.nums), []*big.Int{})

	// Can It Be Solved!?
	fmt.Println("-------Solveable?!--------------------------------")
	fmt.Println("Number of Possible Answers:", len(a))
	for _, i := range a {
		fmt.Println(i)
		if int64(p.answer) == i.Int64() && !p.solveable {
			p.solveable = true
		}
	}
	fmt.Println(p.solveable)
}
