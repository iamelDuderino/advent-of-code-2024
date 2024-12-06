package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
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
		wg:         new(sync.WaitGroup),
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
	wg         *sync.WaitGroup
}

type day6guard struct {
	movingUp, movingDown, movingRight, movingLeft bool
	x, y                                          int
	patrolComplete                                bool
	inLoop                                        bool
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
	coords := onMap.getCoord(xx, yy)
	if coords != nil && onMap.getCoord(xx, yy).obstructed {
		switch {
		case x.movingUp:
			if coords.touchedB {
				x.inLoop = true
			}
			coords.touchedB = true
		case x.movingRight:
			if coords.touchedL {
				x.inLoop = true
			}
			coords.touchedL = true
		case x.movingLeft:
			if coords.touchedR {
				x.inLoop = true
			}
			coords.touchedR = true
		case x.movingDown:
			if coords.touchedU {
				x.inLoop = true
			}
			coords.touchedU = true
		}
		x.turn()
		return
	}
	if xx >= onMap.width || yy >= onMap.height || xx < 0 || yy < 0 {
		x.patrolComplete = true
		return
	}
	x.x = xx
	x.y = yy
	coords.patrolled = true
}

// wouldLoop takes the guards intended travel position and makes it an obstruction, if not already,
// in an attempt to lock the guard in an infinite loop
func (x *day6guard) wouldLoop(xx, yy int) {

	coords := day6.p2guardMap.getCoord(xx, yy)
	if coords == nil {
		return
	}
	if coords.testedForLoop {
		return
	}

	fakeGuard := &day6guard{
		movingUp:       x.movingUp,
		movingDown:     x.movingDown,
		movingRight:    x.movingRight,
		movingLeft:     x.movingLeft,
		x:              x.x,
		y:              x.y,
		patrolComplete: x.patrolComplete,
	}

	if fakeGuard.patrolComplete {
		return
	}

	fakeMap := day6.p2guardMap.copy()

	fakeCoords := fakeMap.getCoord(xx, yy)
	if fakeCoords == nil {
		return
	}
	if fakeCoords.obstructed {
		return
	}
	if !fakeCoords.obstructed {
		fakeCoords.obstructed = true
	}

	var (
		n   int = 0
		max int = 999999999
	)
	coords.testedForLoop = true
	for !fakeGuard.patrolComplete {
		if n == max {
			coords.causesLoop = true
			break
		}
		if fakeGuard.inLoop {
			coords.causesLoop = true
			break
		}
		fakeGuard.move(fakeMap)
		n += 1
	}

}

type day6guardMap struct {
	pos           []*day6guardMapPos
	width, height int
}

func (x *day6guardMap) copy() *day6guardMap {
	m := new(day6guardMap)
	m.width = x.width
	m.height = x.height
	for _, i := range x.pos {
		newPos := &day6guardMapPos{
			x:             i.x,
			y:             i.y,
			obstructed:    i.obstructed,
			testedForLoop: i.testedForLoop,
			causesLoop:    i.causesLoop,
		}
		m.pos = append(m.pos, newPos)
	}
	return m
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

func (x *day6guardMap) countLoopOpportunities() int {
	var n int
	for _, i := range x.pos {
		if i.causesLoop {
			n += 1
		}
	}
	return n
}

type day6guardMapPos struct {
	x, y                                   int
	obstructed                             bool
	patrolled                              bool
	testedForLoop                          bool
	causesLoop                             bool
	touchedL, touchedR, touchedU, touchedB bool
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

	for !x.p2guard.patrolComplete {

		var xx, yy int
		switch {
		case x.p2guard.movingUp:
			xx = x.p2guard.x
			yy = x.p2guard.y - 1
		case x.p2guard.movingRight:
			xx = x.p2guard.x + 1
			yy = x.p2guard.y
		case x.p2guard.movingDown:
			xx = x.p2guard.x
			yy = x.p2guard.y + 1
		case x.p2guard.movingLeft:
			xx = x.p2guard.x - 1
			yy = x.p2guard.y
		}

		x.p2guard.wouldLoop(xx, yy)
		x.p2guard.move(day6.p2guardMap)
	}

	x.wg.Wait()
	fmt.Println("Part 2 Solution:", x.p2guardMap.countLoopOpportunities())
}
