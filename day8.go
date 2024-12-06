package main

import "fmt"

var (
	day8 = &aocDay8{
		banner: `-------------
-   DAY 8   -
-------------`,
		data: dataFolder + "day8.txt",
	}
)

type aocDay8 struct {
	banner string
	data   string
}

func (x *aocDay8) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay8) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay8) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
