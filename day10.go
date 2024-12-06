package main

import "fmt"

var (
	day10 = &aocDay10{
		banner: `-------------
-   DAY 10   -
-------------`,
		data: dataFolder + "day10.txt",
	}
)

type aocDay10 struct {
	banner string
	data   string
}

func (x *aocDay10) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay10) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay10) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
