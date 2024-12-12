package main

import (
	"fmt"
	"math/big"
	"strconv"
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

func TestDay9Part1(t *testing.T) {

	// Stage Test Case
	var (
		hdd       = "2333133121414131402"
		allocated = "00...111...2...333.44.5555.6666.777.888899"
		defragged = "0099811188827773336446555566.............."
		checksum  = 1928
	)

	// Read "File"
	for _, i := range hdd {
		day9.mem = append(day9.mem, &day9mem{
			runes: []rune{i},
		})
	}

	// Allocate
	day9.allocate()
	if day9.str != allocated {
		t.Fatalf("%s != %s\n", day9.str, allocated)
	}
	t.Log("Allocated:", day9.str)
	day9.cache = day9.mem

	// Defrag
	day9.defrag()
	if day9.str != defragged {
		t.Fatalf("%s != %s\n", day9.str, defragged)
	}
	t.Log("Defragged:", day9.str)

	// Checksum
	c := day9.checksum()
	if c == 0 {
		t.Fatal("checksum == 0")
	}
	if c != checksum {
		t.Fatalf("%d != %d", c, checksum)
	}
	t.Log("Checksum:", c)

}

func TestDay9Part2(t *testing.T) {

	// Stage Test Case
	var (
		hdd       = "2333133121414131402"
		allocated = "00...111...2...333.44.5555.6666.777.888899"
		defragged = "00992111777.44.333....5555.6666.....8888.."
		checksum  = 2858
	)

	// Read "File"
	for _, i := range hdd {
		day9.mem = append(day9.mem, &day9mem{
			runes: []rune{i},
		})
	}

	// Allocate
	day9.allocate()
	if day9.str != allocated {
		t.Fatalf("%s != %s\n", day9.str, allocated)
	}
	t.Log("Allocated:", day9.str)
	day9.cache = day9.mem

	// Part 2
	day9.p2 = true
	day9.defrag()
	if day9.str != defragged {
		t.Fatalf("%s != %s", day9.str, defragged)
	}
	c := day9.checksum()
	if c != checksum {
		t.Fatalf("%d != %d", c, checksum)
	}
	t.Log(day9.str)
}

// day 9 curveball, a multi-digit number splits into different runes
// the test case did not cover this as the max number was 9!
func TestRunes(t *testing.T) {
	var (
		i = 10
		r = []rune{}
	)
	r = append(r, []rune(fmt.Sprint(i))...)
	for _, i := range r {
		fmt.Println(string(i))
	}
	n, err := strconv.Atoi(string(r))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n)
}
