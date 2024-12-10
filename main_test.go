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

func TestDay8CoordCalc(t *testing.T) {
	// ..........
	// ..........
	// ..B.......
	// A.........
	// ..........
	// .....C....
	var (
		ax = 0
		ay = 3
		bx = 2
		by = 2
		cx = 5
		cy = 5
	)
	d := calculateDistance(ax, ay, bx, by)
	fmt.Println("Distance between A and B:", d)
	d = calculateDistance(ax, ay, cx, cy)
	fmt.Println("Distance between A and C:", d)
}

func TestSameSlope(t *testing.T) {
	var (
		x1 = 1
		y1 = 1
		x2 = 5
		y2 = 7
		x3 = 9
		y3 = 9
		x4 = 9
		y4 = 13
	)
	t1 := sameSlope(x1, y1, x2, y2, x3, y3)
	t2 := sameSlope(x1, y1, x2, y2, x4, y4)
	if t1 {
		t.Fatal("t1 should be false")
	}
	if !t2 {
		t.Fatal("t2 should be true")
	}
}

// ensure lowercase and uppercase are considered != for day 8
func TestFrequencies(t *testing.T) {
	var (
		a = `n`
		b = `N`
	)
	if a == b {
		t.Fatal("should not be equal")
	}
}
