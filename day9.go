package main

import "fmt"

var (
	day9 = &aocDay9{
		banner: `-------------
-   DAY 9   -
-------------`,
		data: dataFolder + "day9.txt",
	}
)

type aocDay9 struct {
	banner string
	data   string
}

func (x *aocDay9) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay9) part1() {
	var (
		answer string
	)

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay9) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
