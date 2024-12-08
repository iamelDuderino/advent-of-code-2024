package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	day7 = &aocDay7{
		banner: `-------------
-   DAY 7   -
-------------`,
		data: dataFolder + "day7.txt",
	}

	day7rxp = regexp.MustCompile(`^(\d+)(?:\: )([\d\s]+$)`)
)

type aocDay7 struct {
	banner   string
	data     string
	problems []*day7problem
}

type day7problem struct {
	answer    int
	nums      []int
	solveable bool
}

func (x *aocDay7) printBanner() {
	fmt.Println(x.banner)
}

type solution struct {
	solutions []int
}

func recursion(n []int, p []int) (solution, []int) {
	if len(n) == 1 {
		s := solution{
			solutions: []int{n[0]},
		}
		s.solutions = append(s.solutions, p...)
		return s, p
	}

	if len(p) == 0 {
		p = append(p, (n[0] + n[1]))
		p = append(p, (n[0] * n[1]))
	} else {
		for idx, i := range p {
			p[idx] = p[idx] * n[1]
			p = append(p, (i + n[1]))
		}
	}

	n[1] = n[0] + n[1]

	return recursion(n[1:], p)
}

func (x *aocDay7) readFile() {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	var n int
	for s.Scan() {
		line := s.Text()
		rxp := day7rxp.FindStringSubmatch(line)
		if len(rxp) != 3 {
			fmt.Printf("Line %d did not split into 3 capture groups!\n", n)
			return
		}
		answer, err := strconv.Atoi(rxp[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		nums := rxp[2]
		problem := &day7problem{
			answer: answer,
		}
		ss := strings.Split(nums, " ")
		for _, i := range ss {
			n, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			problem.nums = append(problem.nums, n)
		}
		day7.problems = append(day7.problems, problem)
		n += 1
	}
}

func (x *aocDay7) part1() {
	var (
		answer = func() int {
			var n int
			for _, i := range day7.problems {
				if i.solveable {
					n += i.answer
				}
			}
			return n
		}
	)
	x.readFile()

	for _, i := range day7.problems {
		s, _ := recursion(i.nums, []int{})
		for _, sol := range s.solutions {
			if i.answer == sol {

				i.solveable = true
			}
		}
	}

	fmt.Println("Part 1 Solution:", answer())
}

func (x *aocDay7) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
