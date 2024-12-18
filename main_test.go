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
		i = 253000
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

	sr := []rune(fmt.Sprint(i))
	if len(r)%2 == 0 {
		fmt.Println("sr is even")
	} else {
		fmt.Println(sr)
	}

}

func TestDay10(t *testing.T) {
	var (
		numTrailHeads = 9
		score         = 36
		p2score       = 81
		trailheads    int
	)
	day10.readFile()

	day10.p2 = true

	for _, i := range day10.trails {
		if i.isTrailHead() {
			trailheads += 1
			i.ping(i)
		}
	}
	if trailheads != numTrailHeads {
		t.Fatalf("%d != %d", trailheads, numTrailHeads)
	}
	cscore := day10.calculateScore()
	if !day10.p2 && cscore != score {
		t.Fatalf("%d != %d", cscore, score)
	}
	if day10.p2 && cscore != p2score {
		t.Fatalf("%d != %d", cscore, score)
	}
	fmt.Println("Score:", cscore)
}

func TestDay12(t *testing.T) {
	var (
		totalPriceA = 1930
		totalPriceB = 1206
	)
	day12.readFile()
	for _, i := range day12.garden.plants {
		if !i.isRegioned {
			p := i.ping(i.seed, []*day12plant{})
			r := &day12region{
				seed:   i.seed,
				plants: p,
			}
			r.calculateArea()
			r.calculatePerimeter()
			r.calculateCost("perimeter")
			day12.garden.regions = append(day12.garden.regions, r)
		}
	}
	c := day12.garden.calculateCost()
	if c != totalPriceA {
		t.Fatalf("%d != %d", c, totalPriceA)
	}
	fmt.Printf("Total Cost A %d == %d\n", c, totalPriceA)
	for n, i := range day12.garden.regions {
		fmt.Printf("Calculating Region %d/%d\n", n+1, len(day12.garden.regions))
		i.draw()
		i.calculateArea()
		i.calculateSides()
		i.calculateCost("sides")
		fmt.Println("Area:", i.area)
		fmt.Println("Sides:", i.sides)
		fmt.Println("Cost:", i.cost)
	}
	c = day12.garden.calculateCost()
	if c != totalPriceB {
		t.Fatalf("%d != %d", c, totalPriceB)
	}
	fmt.Printf("Total Cost B %d == %d\n", c, totalPriceB)
}

func TestDay12Region(t *testing.T) {
	day12.readFile()
	for _, i := range day12.garden.plants {
		if !i.isRegioned {
			p := i.ping(i.seed, []*day12plant{})
			r := &day12region{
				seed:   i.seed,
				plants: p,
			}
			r.calculateArea()
			r.calculatePerimeter()
			r.calculateCost("perimeter")
			day12.garden.regions = append(day12.garden.regions, r)
		}
	}
	r := day12.garden.regions[283]
	// r.draw()
	r.calculateSides()
	fmt.Println("Area:", r.area)
	fmt.Println("Sides:", r.sides)
	fmt.Println("Cost:", r.cost)
}
