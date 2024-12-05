package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	day4 = &aocDay4{
		banner: `-------------
-   DAY 4   -
-------------`,
		data: dataFolder + "day4.txt",
		xmas: []string{`X`, `M`, `A`, `S`},
		mas:  []string{`M`, `A`, `S`},
		grid: make(map[int]map[int]string),
	}
)

type aocDay4 struct {
	banner string
	data   string
	xmas   []string
	mas    []string
	grid   map[int]map[int]string
}

func (x *aocDay4) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay4) readFile() error {
	f, err := os.Open(x.data)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	var row int
	for s.Scan() {
		x.grid[row] = make(map[int]string)
		for col, i := range s.Text() {
			x.grid[row][col] = string(i)
		}
		row += 1
	}
	return nil
}

func (x *aocDay4) search(row, col int) int {
	var i int
	// check up
	if x.seek("up", row, col) {
		i += 1
	}
	// check down
	if x.seek("down", row, col) {
		i += 1
	}
	// check left
	if x.seek("left", row, col) {
		i += 1
	}
	// check right
	if x.seek("right", row, col) {
		i += 1
	}
	// check up-left
	if x.seek("up-left", row, col) {
		i += 1
	}
	// check up-right
	if x.seek("up-right", row, col) {
		i += 1
	}
	// check down-left
	if x.seek("down-left", row, col) {
		i += 1
	}
	// check down-right
	if x.seek("down-right", row, col) {
		i += 1
	}
	return i
}

// assumes X is the starting position
func (x *aocDay4) seek(dir string, row, col int) bool {
	switch dir {
	case "up":
		return x.grid[row-1][col] == x.xmas[1] &&
			x.grid[row-2][col] == x.xmas[2] &&
			x.grid[row-3][col] == x.xmas[3]
	case "down":
		return x.grid[row+1][col] == x.xmas[1] &&
			x.grid[row+2][col] == x.xmas[2] &&
			x.grid[row+3][col] == x.xmas[3]
	case "left":
		return x.grid[row][col-1] == x.xmas[1] &&
			x.grid[row][col-2] == x.xmas[2] &&
			x.grid[row][col-3] == x.xmas[3]
	case "right":
		return x.grid[row][col+1] == x.xmas[1] &&
			x.grid[row][col+2] == x.xmas[2] &&
			x.grid[row][col+3] == x.xmas[3]
	case "up-left":
		return x.grid[row-1][col-1] == x.xmas[1] &&
			x.grid[row-2][col-2] == x.xmas[2] &&
			x.grid[row-3][col-3] == x.xmas[3]
	case "up-right":
		return x.grid[row-1][col+1] == x.xmas[1] &&
			x.grid[row-2][col+2] == x.xmas[2] &&
			x.grid[row-3][col+3] == x.xmas[3]
	case "down-left":
		return x.grid[row+1][col-1] == x.xmas[1] &&
			x.grid[row+2][col-2] == x.xmas[2] &&
			x.grid[row+3][col-3] == x.xmas[3]
	case "down-right":
		return x.grid[row+1][col+1] == x.xmas[1] &&
			x.grid[row+2][col+2] == x.xmas[2] &&
			x.grid[row+3][col+3] == x.xmas[3]
	default:
		return false
	}
}

// anchor on A, must contain 2x MAS
func (x *aocDay4) searchXMAS(row, col int) int {
	// check \
	if x.grid[row-1][col-1] == x.mas[0] && x.grid[row+1][col+1] == x.mas[2] ||
		x.grid[row-1][col-1] == x.mas[2] && x.grid[row+1][col+1] == x.mas[0] {
	} else {
		return 0
	}
	// check /
	if x.grid[row-1][col+1] == x.mas[0] && x.grid[row+1][col-1] == x.mas[2] ||
		x.grid[row-1][col+1] == x.mas[2] && x.grid[row+1][col-1] == x.mas[0] {
	} else {
		return 0
	}
	return 1
}

func (x *aocDay4) part1() {
	var (
		answer int
	)
	err := x.readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	for n, row := range x.grid {
		for col, s := range row {
			if s == x.xmas[0] {
				answer += x.search(n, col)
			}
		}
	}

	fmt.Println("Day 4, Part 1 Solution:", answer)
}

func (x *aocDay4) part2() {
	var (
		answer int
	)
	for n, row := range x.grid {
		for col, s := range row {
			if s == x.mas[1] {
				answer += x.searchXMAS(n, col)
			}
		}
	}

	fmt.Println("Day 4, Part 2 Solution:", answer)
}
