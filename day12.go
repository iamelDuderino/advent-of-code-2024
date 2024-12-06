package main

import "fmt"

var (
	day12 = &aocDay12{
		banner: `-------------
-   DAY 12   -
-------------`,
		data: dataFolder + "day12.txt",
	}
)

type aocDay12 struct {
	banner string
	data   string
}

func (x *aocDay12) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay12) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay12) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
