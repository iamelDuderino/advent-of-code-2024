package main

import "fmt"

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

func (x *aocDay7) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay7) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
