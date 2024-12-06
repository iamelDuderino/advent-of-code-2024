package main

import "fmt"

var (
	day13 = &aocDay13{
		banner: `-------------
-   DAY 13   -
-------------`,
		data: dataFolder + "day13.txt",
	}
)

type aocDay13 struct {
	banner string
	data   string
}

func (x *aocDay13) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay13) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay13) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
