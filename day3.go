package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	day3 = &aocDay3{
		banner: `-------------
-   DAY 3   -
-------------`,
		data:    dataFolder + "day3.txt",
		rxp:     regexp.MustCompile(`mul\(\d*,\d*\)`),
		rxpExcl: regexp.MustCompile(`(?s)don't\(\).+?do\(\)`),
	}
)

type aocDay3 struct {
	banner  string
	data    string
	rxp     *regexp.Regexp
	rxpExcl *regexp.Regexp
}

func (x *aocDay3) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay3) parseMul(s string) (int, int, error) {
	ss := strings.Replace(s, "mul(", "", -1)
	ss = strings.Replace(ss, ")", "", -1)
	sss := strings.Split(ss, ",")
	xs := sss[0]
	tx, err := strconv.Atoi(xs)
	if err != nil {
		return 0, 0, err
	}
	ys := sss[1]
	ty, err := strconv.Atoi(ys)
	if err != nil {
		return 0, 0, nil
	}
	return tx, ty, nil
}

func (x *aocDay3) part1() {
	var (
		answer int
	)
	b, err := os.ReadFile(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	results := x.rxp.FindAllString(string(b), -1)
	for _, i := range results {
		x, y, err := x.parseMul(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		answer += x * y

	}

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay3) part2() {
	var (
		answer int
	)
	b, err := os.ReadFile(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := string(b)

	removals := x.rxpExcl.FindAllString(s, -1)
	for _, i := range removals {
		s = strings.Replace(s, i, "", -1)
	}
	results := x.rxp.FindAllString(s, -1)
	for _, i := range results {
		x, y, err := x.parseMul(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		answer += x * y
	}

	fmt.Println("Part 2 Solution:", answer)
}
