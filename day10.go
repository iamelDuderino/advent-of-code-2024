package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	zero int = iota
	one
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

var (
	day10 = &aocDay10{
		banner: `-------------
-   DAY 10   -
-------------`,
		data: dataFolder + "day10.txt",
	}
)

type aocDay10 struct {
	banner        string
	data          string
	trails        []*day10trail
	width, height int
	p2            bool
}

func (x *aocDay10) get(xx, yy int) *day10trail {
	for _, i := range x.trails {
		if i.x == xx && i.y == yy {
			return i
		}
	}
	return nil
}

type day10trail struct {
	x, y        int
	score       int
	height      int
	trailsFound []*day10trail
}

func (x *day10trail) hasSeen(t *day10trail) bool {
	for _, i := range x.trailsFound {
		if i.x == t.x && i.y == t.y {
			return true
		}
	}
	return false
}

// ping returns a trails total
func (x *day10trail) isTrailHead() bool {
	return x.height == zero
}

func (x *day10trail) ping(th *day10trail) {
	if x.height == nine {
		if !th.hasSeen(x) || day10.p2 {
			// fmt.Printf("(%d,%d) found a trail ending @ (%d, %d)\n", th.x, th.y, x.x, x.y)
			th.trailsFound = append(th.trailsFound, x)
			th.score += 1
		}
		return
	}
	if x.up() != nil && x.up().height == x.height+1 {
		x.up().ping(th)
	}
	if x.down() != nil && x.down().height == x.height+1 {
		x.down().ping(th)
	}
	if x.left() != nil && x.left().height == x.height+1 {
		x.left().ping(th)
	}
	if x.right() != nil && x.right().height == x.height+1 {
		x.right().ping(th)
	}
}

func (x *aocDay10) calculateScore() int {
	var n int
	for _, i := range x.trails {
		n += i.score
	}
	return n
}

func (x *day10trail) left() *day10trail {
	if x.x-1 < 0 {
		return nil
	}
	return day10.get(x.x-1, x.y)
}

func (x *day10trail) right() *day10trail {
	if x.x+1 > day10.width {
		return nil
	}
	return day10.get(x.x+1, x.y)
}

func (x *day10trail) up() *day10trail {
	if x.y-1 < 0 {
		return nil
	}
	return day10.get(x.x, x.y-1)
}

func (x *day10trail) down() *day10trail {
	if x.y+1 > day10.height {
		return nil
	}
	return day10.get(x.x, x.y+1)
}

func (x *aocDay10) readFile() {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	var xx, yy int
	for s.Scan() {
		line := s.Text()
		xx = len(line)
		for n, ch := range line {
			s := string(ch)
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				return
			}
			t := &day10trail{
				x:      n,
				y:      yy,
				height: i,
			}
			x.trails = append(x.trails, t)

		}
		yy += 1
	}
	day10.width = xx
	day10.height = yy
}

func (x *aocDay10) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay10) part1() {
	x.readFile()
	for _, i := range x.trails {
		if i.isTrailHead() {
			i.ping(i)
		}
	}
	fmt.Println("Part 1 Solution:", x.calculateScore())
}

func (x *aocDay10) part2() {
	x.readFile()
	x.p2 = true
	for _, i := range x.trails {
		if i.isTrailHead() {
			i.ping(i)
		}
	}
	fmt.Println("Part 2 Solution:", x.calculateScore())
}
