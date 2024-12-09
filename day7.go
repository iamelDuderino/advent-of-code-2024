package main

import (
	"bufio"
	"fmt"
	"math/big"
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

	day7rxp    = regexp.MustCompile(`^(\d+)(?:\: )([\d\s]+$)`)
	day7answer = func(p []*day7problem) int {
		var n int
		for _, i := range p {
			if i.solveable {
				n += i.answer
			}
		}
		return n
	}
)

type aocDay7 struct {
	banner string
	data   string
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

func bigIntRecursion(n []*big.Int, p []*big.Int) ([]*big.Int, []*big.Int) {
	if len(n) == 1 {
		return n, p
	}
	if len(p) == 0 {
		p = append(p, new(big.Int).Add(n[0], n[1]))
		p = append(p, new(big.Int).Mul(n[0], n[1]))
		p = append(p, merge(n[0], n[1]))
	} else {
		var np = []*big.Int{}
		for _, i := range p {
			np = append(np, new(big.Int).Add(i, n[1]))
			np = append(np, new(big.Int).Mul(i, n[1]))
			np = append(np, merge(i, n[1]))
		}
		p = np
	}
	return bigIntRecursion(n[1:], p)
}

func (x *aocDay7) readFile() []*day7problem {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s := bufio.NewScanner(f)
	var n int
	var p = []*day7problem{}
	for s.Scan() {
		line := s.Text()
		rxp := day7rxp.FindStringSubmatch(line)
		if len(rxp) != 3 {
			fmt.Printf("Line %d did not split into 3 capture groups!\n", n)
			return nil
		}
		answer, err := strconv.Atoi(rxp[1])
		if err != nil {
			fmt.Println(err)
			return nil
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
				return nil
			}
			problem.nums = append(problem.nums, n)
		}
		p = append(p, problem)
		n += 1
	}
	return p
}

func merge(a, b *big.Int) *big.Int {
	c, ok := new(big.Int).SetString(a.String()+b.String(), 10)
	if !ok {
		return nil
	}
	return c
}

func convertToBigInts(nums []int) []*big.Int {
	b := []*big.Int{}
	for _, i := range nums {
		b = append(b, big.NewInt(int64(i)))
	}
	return b
}

func (x *aocDay7) part1() {
	p := x.readFile()
	for _, i := range p {
		s, _ := recursion(i.nums, []int{})
		for _, sol := range s.solutions {
			if i.answer == sol {
				i.solveable = true
			}
		}
	}
	fmt.Println("Part 1 Solution:", day7answer(p))
}

func (x *aocDay7) part2() {
	p := x.readFile()
	for _, i := range p {
		if i.solveable {
			continue
		}
		_, r := bigIntRecursion(convertToBigInts(i.nums), []*big.Int{})
		for _, ii := range r {
			if int64(i.answer) == ii.Int64() && !i.solveable {
				i.solveable = true
			}
		}
	}
	fmt.Println("Part 2 Solution:", day7answer(p))
}
