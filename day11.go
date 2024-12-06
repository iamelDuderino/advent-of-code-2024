package main

import "fmt"

var (
	day11 = &aocDay11{
		banner: `-------------
-   DAY 11   -
-------------`,
		data: dataFolder + "day11.txt",
	}
)

type aocDay11 struct {
	banner string
	data   string
}

func (x *aocDay11) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay11) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay11) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
