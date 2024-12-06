package main

import "fmt"

var (
	day14 = &aocDay14{
		banner: `-------------
-   DAY 14   -
-------------`,
		data: dataFolder + "day14.txt",
	}
)

type aocDay14 struct {
	banner string
	data   string
}

func (x *aocDay14) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay14) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay14) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
