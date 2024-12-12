package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	d11opt1 d11opt = iota
	d11opt2
	d11opt3
)

var (
	day11 = &aocDay11{
		banner: `-------------
-   DAY 11   -
-------------`,
		data: dataFolder + "day11.txt",
		// plutonianPebbles: make([]int, 1000000),
	}
)

type d11opt int

type aocDay11 struct {
	banner           string
	data             string
	plutonianPebbles []int
	cache            []int
}

func (x *aocDay11) blink() {
	type insertion struct {
		idx int
		n   int
	}
	var (
		insertions = []*insertion{}
	)
	for idx, i := range x.plutonianPebbles {
		switch x.consider(i) {
		case d11opt1:
			x.plutonianPebbles[idx] = x.opt1()
		case d11opt2:
			i1, i2 := x.opt2(i)
			x.plutonianPebbles[idx] = i1
			if len(x.plutonianPebbles)-1 == idx {
				x.plutonianPebbles = append(x.plutonianPebbles, i2)
			} else {
				insertions = append(insertions, &insertion{
					idx: idx + 1,
					n:   i2,
				})
			}
		case d11opt3:
			x.plutonianPebbles[idx] = x.opt3(i)
		}
	}
	for _, i := range insertions {
		x.cache = make([]int, len(x.plutonianPebbles[i.idx:]))
		copy(x.cache, x.plutonianPebbles[i.idx:])
		x.plutonianPebbles = x.plutonianPebbles[:i.idx]
		x.plutonianPebbles = append(x.plutonianPebbles, i.n)
		x.plutonianPebbles = append(x.plutonianPebbles, x.cache...)
		for _, ii := range insertions {
			if ii.idx == i.idx {
				continue
			}
			ii.idx += 1
		}
	}
}

func (x *aocDay11) consider(n int) d11opt {
	if n == 0 {
		return d11opt1
	}
	r := []rune(fmt.Sprint(n))
	if len(r)%2 == 0 {
		return d11opt2
	}
	return d11opt3
}

func (x *aocDay11) opt1() int {
	return 1
}

func (x *aocDay11) opt2(p int) (int, int) {
	r := []rune(fmt.Sprint(p))
	r1 := r[:len(r)/2]
	r2 := r[len(r)/2:]
	do := func(r []rune) int {
		var s string
		for _, i := range r {
			s += string(i)
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return n
	}
	return do(r1), do(r2)
}

func (x *aocDay11) opt3(p int) int {
	return p * 2024
}

func (x *aocDay11) readFile() {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		s := strings.Split(s.Text(), ` `)
		for _, i := range s {
			n, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			x.plutonianPebbles = append(x.plutonianPebbles, n)
		}
	}
}

func (x *aocDay11) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay11) part1() {
	x.readFile()
	numTimesToBlink := 25
	for n := 1; n <= numTimesToBlink; n++ {
		fmt.Printf(" ~ BLINK %d ~  %d\n", n, len(x.plutonianPebbles))
		// fmt.Println(x.plutonianPebbles)
		x.blink()
		// fmt.Println(x.plutonianPebbles)
	}
	// fmt.Println(x.plutonianPebbles)
	fmt.Println("Part 1 Solution:", len(x.plutonianPebbles))
}

func (x *aocDay11) part2() {
	x.readFile()
	numTimesToBlink := 75
	for n := 1; n <= numTimesToBlink; n++ {
		fmt.Printf(" ~ BLINK %d | %d ~\n", n, len(x.plutonianPebbles))
		x.blink()
	}
	fmt.Println("Part 2 Solution:", len(x.plutonianPebbles))
}
