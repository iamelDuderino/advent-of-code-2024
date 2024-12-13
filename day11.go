package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// big ups to https://github.com/hsabbas
// for the epic recursive map memoization method!

// Below blink() function worked for part 1 but got sluggish around 30th iteration, crashing out around 40

// func (x *aocDay11) blink() {
// 	var insertions = make(map[int]int)
// 	for idx, i := range x.plutonianPebbles {
// 		switch x.consider(i) {
// 		case d11opt1:
// 			x.plutonianPebbles[idx] = x.opt1()
// 		case d11opt2:
// 			i1, i2 := x.opt2(i)
// 			x.plutonianPebbles[idx] = i1
// 			if len(x.plutonianPebbles)-1 == idx {
// 				x.plutonianPebbles = append(x.plutonianPebbles, i2)
// 			} else {
// 				insertions[idx+1] = i2
// 			}
// 		case d11opt3:
// 			x.plutonianPebbles[idx] = x.opt3(i)
// 		}
// 	}
// 	var iter int
// 	for idx, n := range insertions {
// 		idx = idx + iter
// 		x.plutonianPebbles = append(x.plutonianPebbles, 0)
// 		copy(x.plutonianPebbles[idx+1:], x.plutonianPebbles[idx:])
// 		x.plutonianPebbles[idx] = n
// 		iter += 1
// 	}
// }

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
	}
)

type d11opt int

type aocDay11 struct {
	banner           string
	data             string
	plutonianPebbles []int
}

type plutonianPebble struct {
	num    int // original number
	blinks int // remaining number of blinks
} // mapped to current total

func (x *aocDay11) blink(pebbles []int, blinks int) int {
	m := make(map[plutonianPebble]int)
	t := 0
	for _, i := range pebbles {
		t += x.blinkPebble(i, blinks, m)
	}
	return t
}

func (x *aocDay11) blinkPebble(pebble int, blinks int, pebbleMap map[plutonianPebble]int) int {
	if blinks == 0 {
		return 1
	}
	prev := pebbleMap[plutonianPebble{num: pebble, blinks: blinks}]
	if prev != 0 {
		return prev
	}
	t := 0
	switch x.consider(pebble) {
	case d11opt1:
		t = x.blinkPebble(x.opt1(), blinks-1, pebbleMap)
	case d11opt2:
		l, r := x.opt2(pebble)
		t = x.blinkPebble(l, blinks-1, pebbleMap) + x.blinkPebble(r, blinks-1, pebbleMap)
	case d11opt3:
		t = x.blinkPebble(x.opt3(pebble), blinks-1, pebbleMap)
	}
	pebbleMap[plutonianPebble{num: pebble, blinks: blinks}] = t
	return t
}

func (x *aocDay11) consider(n int) d11opt {
	if n == 0 {
		return d11opt1
	}
	s := fmt.Sprintf("%d", n)
	if len(s)%2 == 0 {
		return d11opt2
	}
	return d11opt3
}

func (x *aocDay11) opt1() int {
	return 1
}

func (x *aocDay11) opt2(p int) (int, int) {
	s := fmt.Sprintf("%d", p)
	len := len(s) / 2
	l := s[:len]
	r := s[len:]
	left, _ := strconv.Atoi(l)
	right, _ := strconv.Atoi(r)
	return left, right
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
		x.plutonianPebbles = make([]int, len(s))
		for idx, i := range s {
			n, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			x.plutonianPebbles[idx] = n
		}
	}
}

func (x *aocDay11) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay11) part1() {
	x.readFile()
	numTimesToBlink := 25
	total := x.blink(x.plutonianPebbles, numTimesToBlink)
	fmt.Println("Part 1 Solution:", total)
}

func (x *aocDay11) part2() {
	x.readFile()
	numTimesToBlink := 75
	total := x.blink(x.plutonianPebbles, numTimesToBlink)
	fmt.Println("Part 2 Solution:", total)
}
