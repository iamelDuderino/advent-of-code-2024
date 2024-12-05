package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	day5 = &aocDay5{
		banner: `-------------
-   DAY 5   -
-------------`,
		data:    dataFolder + "day5.txt",
		rules:   make(map[int][]int),
		updates: []*day5update{},
	}
)

type aocDay5 struct {
	banner  string
	data    string
	rules   map[int][]int
	updates []*day5update
}

type day5update struct {
	pages []int
}

func (x *aocDay5) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay5) readFile() error {
	f, err := os.Open(x.data)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	var iter int
	for s.Scan() {

		line := s.Text()

		switch {

		// sets a rule
		case strings.Contains(line, "|"):
			l := strings.Split(line, "|")
			beforeInt, err := strconv.Atoi(l[0])
			if err != nil {
				return err
			}
			afterInt, err := strconv.Atoi(l[1])
			if err != nil {
				return err
			}
			x.rules[beforeInt] = append(x.rules[beforeInt], afterInt)

		// sets an update
		case strings.Contains(line, ","):
			l := strings.Split(line, ",")
			u := new(day5update)
			for _, i := range l {
				ii, err := strconv.Atoi(i)
				if err != nil {
					return err
				}
				u.pages = append(u.pages, ii)
			}
			x.updates = append(x.updates, u)
		}

		iter += 1

	}
	return nil
}

func (x *day5update) isValid() bool {
	for idx, page := range x.pages {
		hasRules, rules := day5.hasRules(page)
		if hasRules {
			for _, rule := range rules {
				for _, p := range x.pages[:idx] {
					if rule == p {
						return false
					}
				}
			}
		}
	}
	return true
}

func (x *aocDay5) hasRules(page int) (bool, []int) {
	for n, i := range x.rules {
		if n == page {
			return true, i
		}
	}
	return false, []int{}
}

func (x *day5update) getMedian() int {
	if len(x.pages)%2 == 0 {
		return x.pages[len(x.pages)/2]
	}
	return x.pages[(len(x.pages)-1)/2]
}

func (x *day5update) makeValid() {
	for idx, page := range x.pages {
		hasRules, rules := day5.hasRules(page)
		if hasRules {
			for _, rule := range rules {
				for ridx, p := range x.pages[:idx] {
					if rule == p {
						x.reindex(page, ridx)
					}
				}
			}
		}
	}
}

func (x *day5update) reindex(i, idx int) {
	var reindexed []int
	for n, page := range x.pages {
		if n == idx {
			reindexed = append(reindexed, i)
			reindexed = append(reindexed, page)
			continue
		}
		if page == i {
			continue
		}
		reindexed = append(reindexed, page)
	}
	x.pages = reindexed
}

func (x *aocDay5) part1() {
	var (
		answer int
	)
	err := x.readFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, u := range x.updates {
		if u.isValid() {
			answer += u.getMedian()
		}
	}

	fmt.Println("Part 1 Solution:", answer)
}

func (x *aocDay5) part2() {
	var (
		answer int
	)
	for _, u := range x.updates {
		if u.isValid() {
			continue
		}
		for !u.isValid() {
			u.makeValid()
		}
		answer += u.getMedian()
	}

	fmt.Println("Part 2 Solution:", answer)
}
