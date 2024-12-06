package main

// Advent of Code 2024
type adventOfCode struct {
	day1 *aocDay1
	day2 *aocDay2
	day3 *aocDay3
	day4 *aocDay4
	day5 *aocDay5
	day6 *aocDay6
	day7 *aocDay7
}

var aoc = adventOfCode{
	day1: day1,
	day2: day2,
	day3: day3,
	day4: day4,
	day5: day5,
	day6: day6,
	day7: day7,
}

const (
	dataFolder = "./data/"
)

func main() {

	// Day 1
	aoc.day1.printBanner()
	aoc.day1.part1()
	aoc.day1.part2()

	// Day 2
	aoc.day2.printBanner()
	aoc.day2.part1()
	aoc.day2.part2()

	// Day 3
	aoc.day3.printBanner()
	aoc.day3.part1()
	aoc.day3.part2()

	// Day 4
	aoc.day4.printBanner()
	aoc.day4.part1()
	aoc.day4.part2()

	// Day 5
	aoc.day5.printBanner()
	aoc.day5.part1()
	aoc.day5.part2()

	// // Day 6
	aoc.day6.printBanner()
	aoc.day6.part1()
	aoc.day6.part2()

	// // Day7
	// aoc.day7.printBanner()
	// aoc.day7.part1()
	// aoc.day7.part2()

}
