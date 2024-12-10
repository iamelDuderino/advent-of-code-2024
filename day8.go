package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
)

const (
	antinode        = `#`
	antinodeAntenna = `@`
)

var (
	day8p2 bool
	day8   = &aocDay8{
		banner: `-------------
-   DAY 8   -
-------------`,
		data: dataFolder + "day8.txt",
	}
	answer = func() int {
		var n int
		for _, i := range day8.grid.coords {
			if i.hasAntinode {
				n += 1
			}
		}
		return n
	}
)

type aocDay8 struct {
	banner string
	data   string
	grid   *day8map
}

type day8map struct {
	width, height int
	coords        []*day8coord
	frequencies   []string
}

func (x *day8map) inFrequencies(s string) bool {
	for _, i := range x.frequencies {
		if i == s {
			return true
		}
	}
	return false
}

func (x *day8map) display() {
	fmt.Printf("\n\n---------------------------------\n\n")
	for _, i := range x.coords {
		switch {
		case i.hasAntenna && i.hasAntinode:
			fmt.Print(antinodeAntenna)
		case !i.hasAntenna && i.hasAntinode:
			fmt.Print(antinode)
		case i.hasAntenna && !i.hasAntinode:
			fmt.Print(i.antenna.frequency)
		default:
			fmt.Print(`.`)
		}
		if i.x == x.width {
			fmt.Println()
		}
	}
	fmt.Printf("\n\n---------------------------------\n\n")
}

type day8coord struct {
	x, y        int
	hasAntenna  bool
	hasAntinode bool
	antenna     *day8antenna
}

type day8antenna struct {
	x, y      int
	frequency string
}

// day8coord.ping surveys the map for similarly sloped antenna frequencies
func (x *day8coord) ping(m *day8map) bool {
	for _, i := range m.coords {
		if i.x == x.x && i.y == x.y {
			continue // self
		}
		if i.hasAntenna && i.antenna.ping(m, x.x, x.y, calculateDistance(x.x, x.y, i.x, i.y)) {
			return true
		}
		if day8p2 {
			if i.hasAntenna && i.antenna.ping(m, i.x, i.y, 0) {
				return true
			}
		}
	}
	return false
}

// day8antenna.ping surveys the map @ distance d and if (anx,any) is same slope to (z.x, z.y) with same frenquency, return true
func (x *day8antenna) ping(m *day8map, anx, any, d int) bool {
	for _, i := range m.coords {
		td := calculateDistance(x.x, x.y, i.x, i.y)
		if td != d && !day8p2 {
			continue
		}
		if i.x == anx && i.y == any && !day8p2 {
			continue // ignore source antinode, unless part 2 as sent antinode = antenna
		}
		if i.hasAntenna && i.antenna.frequency == x.frequency {
			if sameSlope(anx, any, x.x, x.y, i.x, i.y) {
				return true
			}
		}
	}
	return false
}

// calculates the distance from (X1,Y1) to (X2, Y2)
func calculateDistance(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

// x1,y1 should be the antinode
// x2,y2 should be the 1st antenna
// x3,y3 should be the 2nd antenna
// calculates that the slope between positions is the same
func sameSlope(x1, y1, x2, y2, x3, y3 int) bool {
	m1 := ((float64(y2) - float64(y1)) / (float64(x2) - float64(x1)))
	m2 := ((float64(y3) - float64(y2)) / (float64(x3) - float64(x2)))
	return m1 == m2
}

func (x *aocDay8) readFile() *day8map {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s := bufio.NewScanner(f)
	m := &day8map{
		height: -1,
	}
	for s.Scan() {
		line := s.Text()
		if m.width == 0 {
			m.width = len(line) - 1
		}
		m.height += 1
		for n, r := range line {
			ch := string(r)
			c := &day8coord{
				x: n,
				y: m.height,
			}
			if ch != `.` {
				a := &day8antenna{
					x:         n,
					y:         m.height,
					frequency: ch,
				}
				if !m.inFrequencies(ch) {
					m.frequencies = append(m.frequencies, ch)
				}
				c.hasAntenna = true
				c.antenna = a
			}
			m.coords = append(m.coords, c)
		}
	}
	return m
}

func (x *aocDay8) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay8) run() {
	var wg = new(sync.WaitGroup)
	for _, i := range x.grid.coords {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			i.hasAntinode = i.ping(x.grid)
		}(wg)
	}
	wg.Wait()
}

func (x *aocDay8) part1() {
	x.grid = x.readFile()
	x.run()
	x.grid.display()
	fmt.Println("Part 1 Solution:", answer())
}

func (x *aocDay8) part2() {
	day8p2 = true
	x.run()
	x.grid.display()
	fmt.Println("Part 2 Solution:", answer())
}
