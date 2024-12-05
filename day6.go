package main

import "fmt"

// -------------
// -   DAY 6   -
// -------------
//
//

var (
	day6 = &aocDay6{
		banner: `-------------
-   DAY 6   -
-------------`,
		data: dataFolder + "day6.txt",
	}
)

type aocDay6 struct {
	banner string
	data   string
}

func (x *aocDay6) printBanner() {
	fmt.Println(x.banner)
}

// PART 1
func (x *aocDay6) part1() {
	var (
		answer string
	)

	fmt.Println("Day 6, Part 1 Solution:", answer)
}

// PART 2
func (x *aocDay6) part2() {
	var (
		answer string
	)

	fmt.Println("Day 6, Part 2 Solution:", answer)
}
