package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	day6 = &aocDay6{
		banner: `-------------
-   DAY 6   -
-------------`,
		data: dataFolder + "day6.txt",
		p1guard: &day6guard{
			movingUp: true,
		},
		p2guard: &day6guard{
			movingUp: true,
		},
		p1guardMap: new(day6guardMap),
		p2guardMap: new(day6guardMap),
	}
)

const (
	d6free       = `.`
	d6obstructed = `#`
	d6guard      = `^`
)

type aocDay6 struct {
	banner     string
	data       string
	p1guard    *day6guard
	p2guard    *day6guard
	p1guardMap *day6guardMap
	p2guardMap *day6guardMap
}

type day6guard struct {
	movingUp, movingDown, movingRight, movingLeft bool
	x, y                                          int
	patrolComplete                                bool
}

func (x *day6guard) turn() {
	switch {
	case x.movingUp:
		x.movingUp = false
		x.movingRight = true
	case x.movingDown:
		x.movingDown = false
		x.movingLeft = true
	case x.movingRight:
		x.movingRight = false
		x.movingDown = true
	case x.movingLeft:
		x.movingLeft = false
		x.movingUp = true
	}
}

func (x *day6guard) move(onMap *day6guardMap) {
	var (
		xx, yy int
	)
	switch {
	case x.movingUp:
		xx = x.x
		yy = x.y - 1
	case x.movingDown:
		xx = x.x
		yy = x.y + 1
	case x.movingRight:
		xx = x.x + 1
		yy = x.y
	case x.movingLeft:
		xx = x.x - 1
		yy = x.y
	}
	if onMap.getCoord(xx, yy) != nil && onMap.getCoord(xx, yy).obstructed {
		x.turn()
		return
	}
	if xx >= onMap.width || yy >= onMap.height {
		x.patrolComplete = true
		return
	}
	x.x = xx
	x.y = yy
	onMap.getCoord(xx, yy).patrolled = true
}

type day6guardMap struct {
	pos           []*day6guardMapPos
	width, height int
}

func (x *day6guardMap) getCoord(xx, yy int) *day6guardMapPos {
	for _, i := range x.pos {
		if i.x == xx && i.y == yy {
			return i
		}
	}
	return nil
}

func (x *day6guardMap) countPatrolled() int {
	var n int
	for _, i := range x.pos {
		if i.patrolled {
			n += 1
		}
	}
	return n
}

type day6guardMapPos struct {
	x, y          int
	obstructed    bool
	patrolled     bool
	testedForLoop bool
}

func (x *aocDay6) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay6) readFile() error {
	f, err := os.Open(x.data)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	var numRows int = 0
	var numCols int = 0
	var numColsCounted bool
	for s.Scan() {
		for col, ch := range s.Text() {
			pos := &day6guardMapPos{
				x: col,
				y: numRows,
			}
			switch string(ch) {
			case d6free:
				pos.obstructed = false
			case d6obstructed:
				pos.obstructed = true
			case d6guard:
				x.p1guard.x = col
				x.p1guard.y = numRows
				x.p2guard.x = col
				x.p2guard.y = numRows
				pos.patrolled = true
			}
			x.p1guardMap.pos = append(x.p1guardMap.pos, pos)
			x.p2guardMap.pos = append(x.p2guardMap.pos, pos)
			if !numColsCounted {
				numCols += 1
			}
		}
		numRows += 1
		numColsCounted = true
	}
	x.p1guardMap.height = numRows
	x.p1guardMap.width = numCols
	x.p2guardMap.height = numRows
	x.p2guardMap.width = numCols
	return nil
}

func (x *aocDay6) part1() {
	err := day6.readFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	for !day6.p1guard.patrolComplete {
		day6.p1guard.move(day6.p1guardMap)
	}

	fmt.Println("Part 1 Solution:", day6.p1guardMap.countPatrolled())
}

func (x *aocDay6) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
