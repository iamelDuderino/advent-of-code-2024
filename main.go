package main

import "flag"

// Advent of Code 2024
type adventOfCode struct {
	day1  *aocDay1
	day2  *aocDay2
	day3  *aocDay3
	day4  *aocDay4
	day5  *aocDay5
	day6  *aocDay6
	day7  *aocDay7
	day8  *aocDay8
	day9  *aocDay9
	day10 *aocDay10
	day11 *aocDay11
	day12 *aocDay12
	day13 *aocDay13
	day14 *aocDay14
}

var aoc = adventOfCode{
	day1:  day1,
	day2:  day2,
	day3:  day3,
	day4:  day4,
	day5:  day5,
	day6:  day6,
	day7:  day7,
	day8:  day8,
	day9:  day9,
	day10: day10,
	day11: day11,
	day12: day12,
	day13: day13,
	day14: day14,
}

const (
	dataFolder = "./data/"
)

func main() {

	day := flag.Int("day", 0, "-day calls the desired day")
	part := flag.Int("part", 1, "-part calls the desired part for the day")
	flag.Parse()

	switch *day {

	// Day 1
	case 1:
		aoc.day1.printBanner()
		switch *part {
		case 1:
			aoc.day1.part1()
		case 2:
			aoc.day1.part2()
		default:
			aoc.day1.part1()
			aoc.day1.part2()
		}

	// Day 2
	case 2:
		aoc.day2.printBanner()
		switch *part {
		case 1:
			aoc.day2.part1()
		case 2:
			aoc.day2.part2()
		default:
			aoc.day2.part1()
			aoc.day2.part2()
		}

	// Day 3
	case 3:
		aoc.day3.printBanner()
		switch *part {
		case 1:
			aoc.day3.part1()
		case 2:
			aoc.day3.part2()
		default:
			aoc.day3.part1()
			aoc.day3.part2()
		}

	// Day 4
	case 4:
		aoc.day4.printBanner()
		switch *part {
		case 1:
			aoc.day4.part1()
		case 2:
			aoc.day4.part2()
		default:
			aoc.day4.part1()
			aoc.day4.part2()
		}

	// Day 5
	case 5:
		aoc.day5.printBanner()
		switch *part {
		case 1:
			aoc.day5.part1()
		case 2:
			aoc.day5.part2()
		default:
			aoc.day5.part1()
			aoc.day5.part2()
		}

	// Day 6
	case 6:
		aoc.day6.printBanner()
		switch *part {
		case 1:
			aoc.day6.part1()
		case 2:
			aoc.day6.part2()
		default:
			aoc.day6.part1()
			aoc.day6.part2()
		}

	// Day 7
	case 7:
		aoc.day7.printBanner()
		switch *part {
		case 1:
			aoc.day7.part1()
		case 2:
			aoc.day7.part2()
		default:
			aoc.day7.part1()
			aoc.day7.part2()
		}

	// Day 8
	case 8:
		aoc.day8.printBanner()
		switch *part {
		case 1:
			aoc.day8.part1()
		case 2:
			aoc.day8.part2()
		default:
			aoc.day8.part1()
			aoc.day8.part2()
		}

	// Day 9
	case 9:
		aoc.day9.printBanner()
		switch *part {
		case 1:
			aoc.day9.part1()
		case 2:
			aoc.day9.part2()
		default:
			aoc.day9.part1()
			aoc.day9.part2()
		}

	// Day 10
	case 10:
		aoc.day10.printBanner()
		switch *part {
		case 1:
			aoc.day10.part1()
		case 2:
			aoc.day10.part2()
		default:
			aoc.day10.part1()
			aoc.day10.part2()
		}

	// Day 11
	case 11:
		aoc.day11.printBanner()
		switch *part {
		case 1:
			aoc.day11.part1()
		case 2:
			aoc.day11.part2()
		default:
			aoc.day11.part1()
			aoc.day11.part2()
		}

	// Day 12
	case 12:
		aoc.day12.printBanner()
		switch *part {
		case 1:
			aoc.day12.part1()
		case 2:
			aoc.day12.part2()
		default:
			aoc.day12.part1()
			aoc.day12.part2()
		}

	// Day 13
	case 13:
		aoc.day13.printBanner()
		switch *part {
		case 1:
			aoc.day13.part1()
		case 2:
			aoc.day13.part2()
		default:
			aoc.day13.part1()
			aoc.day13.part2()
		}

	// Day 14
	case 14:
		aoc.day14.printBanner()
		switch *part {
		case 1:
			aoc.day14.part1()
		case 2:
			aoc.day14.part2()
		default:
			aoc.day14.part1()
			aoc.day14.part2()
		}

	}

}
