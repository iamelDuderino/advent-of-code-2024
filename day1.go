package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var day1 = &aocDay1{
	banner: `-------------
-   DAY 1   -
-------------`,
	data: dataFolder + "day1.txt",
}

type aocDay1 struct {
	banner string
	data   string
}

func (x *aocDay1) printBanner() {
	fmt.Println(x.banner)
}

type day1LocationId struct {
	int             int
	checked         bool
	similarityScore int
}

type day1LocationMap struct {
	group1, group2 map[int]*day1LocationId
	distances      map[int]int
}

func (x *aocDay1) addToMap(mapper *day1LocationMap, s string) error {
	news := strings.Split(s, "   ")
	i1, err := strconv.Atoi(strings.TrimSpace(news[0]))
	if err != nil {
		return err
	}
	id1 := &day1LocationId{int: i1}
	i2, err := strconv.Atoi(strings.TrimSpace(news[1]))
	if err != nil {
		return err
	}
	id2 := &day1LocationId{int: i2}
	mapper.group1[len(mapper.group1)] = id1
	mapper.group2[len(mapper.group2)] = id2
	return nil
}

func (x *aocDay1) newLocationMap() *day1LocationMap {
	return &day1LocationMap{
		group1:    make(map[int]*day1LocationId),
		group2:    make(map[int]*day1LocationId),
		distances: make(map[int]int),
	}
}

func (x *aocDay1) part1() {

	var (
		mapper      = day1.newLocationMap()
		getDistance = func(i, ii int) int {
			switch {
			case i < ii:
				return ii - i
			case i > ii:
				return i - ii
			default:
				return 0 // i == ii
			}
		}

		solution = func() int {
			var n int
			for i, ii := range mapper.distances {
				n += getDistance(i, ii)
			}
			return n
		}
	)

	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		err := x.addToMap(mapper, s.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if len(mapper.group1) != len(mapper.group2) {
		fmt.Println("group lengths not equal")
		return
	}
	if len(mapper.group1) == 0 || len(mapper.group2) == 0 {
		fmt.Println("group length 0")
		return
	}

	length := len(mapper.group1)
	var iter int
	for iter != length {

		iter += 1

		// lowest in group 1
		var thisIterLowest1 int = 999999999999
		for _, i := range mapper.group1 {
			if !i.checked && (i.int < thisIterLowest1) {
				thisIterLowest1 = i.int
			}
		}

		// go back and mark i as checked
		for _, i := range mapper.group1 {
			if !i.checked && (i.int == thisIterLowest1) {
				i.checked = true
				break
			}
		}

		// lowest in group 2
		var thisIterLowest2 int = 999999999999
		for _, i := range mapper.group2 {
			if !i.checked && (i.int < thisIterLowest2) {
				thisIterLowest2 = i.int
			}
		}

		// go back and mark i as checked
		for _, i := range mapper.group2 {
			if !i.checked && (i.int == thisIterLowest2) {
				i.checked = true
				break
			}
		}

		// mark the distance for this iteration
		mapper.distances[thisIterLowest1] = thisIterLowest2

	}

	// solution() solves for total distance
	fmt.Println("Part 1 Solution:", solution())
}

func (x *aocDay1) part2() {

	var (
		mapper   = x.newLocationMap()
		solution = func() int {
			var s int
			for _, i := range mapper.group1 {
				s += i.similarityScore
			}
			return s
		}
	)

	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		err := x.addToMap(mapper, s.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if len(mapper.group1) != len(mapper.group2) {
		fmt.Println("group lengths not equal")
		return
	}
	if len(mapper.group1) == 0 || len(mapper.group2) == 0 {
		fmt.Println("group length 0")
		return
	}

	// solve for individual similarity score
	for _, i := range mapper.group1 {
		var simScore int
		for _, ii := range mapper.group2 {
			if i.int == ii.int {
				simScore += 1
			}
		}
		i.similarityScore = i.int * simScore
	}

	// solution() solves for total similarity score
	fmt.Println("Part 2 Solution:", solution())
}
