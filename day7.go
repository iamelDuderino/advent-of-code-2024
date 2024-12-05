package main

import "fmt"

// -------------
// -   DAY 7   -
// -------------
//
//

var (
	day7 = &aocDay7{
		banner: `-------------
-   DAY 7   -
-------------`,
		data: dataFolder + "day7.txt",
	}
)

type aocDay7 struct {
	banner string
	data   string
}

func (x *aocDay7) printBanner() {
	fmt.Println(x.banner)
}

// PART 1
func (x *aocDay7) part1() {
	var (
		answer string
	)

	fmt.Println("Day 7, Part 1 Solution:", answer)
}

// PART 2
func (x *aocDay7) part2() {
	var (
		answer string
	)

	fmt.Println("Day 7, Part 2 Solution:", answer)
}
