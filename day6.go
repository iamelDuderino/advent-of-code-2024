package main

import "fmt"

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

func (x *aocDay6) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay6) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
