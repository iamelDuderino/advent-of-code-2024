package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	day12 = &aocDay12{
		banner: `-------------
-   DAY 12   -
-------------`,
		data:   dataFolder + "day12.txt",
		garden: new(day12garden),
	}
)

type aocDay12 struct {
	banner string
	data   string
	garden *day12garden
}

type day12garden struct {
	plants        []*day12plant
	regions       []*day12region
	length, width int
}

func (x *aocDay12) get(xx, yy int) *day12plant {
	for _, i := range x.garden.plants {
		if i.x == xx && i.y == yy {
			return i
		}
	}
	return nil
}

type day12plant struct {
	x, y       int
	seed       string
	isRegioned bool
}

type day12region struct {
	seed                   string
	plants                 []*day12plant
	area, perimeter, sides int
	cost                   int
}

func (x *day12plant) ping(seed string, plants []*day12plant) []*day12plant {
	if x.seed != seed || x.isRegioned {
		return nil
	}
	if !x.isRegioned {
		plants = append(plants, x)
		x.isRegioned = true
	}
	if x.up() != nil && x.up().seed == seed {
		plants = append(plants, x.up().ping(seed, plants)...)
	}
	if x.down() != nil && x.down().seed == seed {
		plants = append(plants, x.down().ping(seed, plants)...)
	}
	if x.left() != nil && x.left().seed == seed {
		plants = append(plants, x.left().ping(seed, plants)...)
	}
	if x.right() != nil && x.right().seed == seed {
		plants = append(plants, x.right().ping(seed, plants)...)
	}
	return day12.garden.removeDuplicates(plants)
}

func (x *day12garden) removeDuplicates(p []*day12plant) []*day12plant {
	var np = []*day12plant{}
	var inNp = func(p *day12plant) bool {
		for _, i := range np {
			if i.x == p.x && i.y == p.y {
				return true
			}
		}
		return false
	}
	for _, i := range p {
		if !inNp(i) {
			np = append(np, i)
		}
	}
	return np
}

func (x *day12plant) up() *day12plant {
	if x.y-1 < 0 {
		return nil
	}
	return day12.get(x.x, x.y-1)
}

func (x *day12plant) down() *day12plant {
	if x.y+1 > day12.garden.length {
		return nil
	}
	return day12.get(x.x, x.y+1)
}

func (x *day12plant) left() *day12plant {
	if x.x-1 < 0 {
		return nil
	}
	return day12.get(x.x-1, x.y)
}

func (x *day12plant) right() *day12plant {
	if x.x+1 > day12.garden.width {
		return nil
	}
	return day12.get(x.x+1, x.y)
}

func (x *aocDay12) readFile() {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	var (
		l, w int
	)
	for s.Scan() {
		line := s.Text()
		if w == 0 {
			w = len(line) - 1
		}
		for idx, ch := range line {
			s := string(ch)
			p := &day12plant{
				x:    idx,
				y:    l,
				seed: s,
			}
			x.garden.plants = append(x.garden.plants, p)
		}
		l += 1
	}
	day12.garden.length = l - 1
	day12.garden.width = w
}

func (x *day12region) calculateArea() {
	x.area = len(x.plants)
}

func (x *day12region) calculatePerimeter() {
	for _, i := range x.plants {
		if i.up() != nil && i.up().seed != x.seed || i.up() == nil {
			x.perimeter += 1
		}
		if i.down() != nil && i.down().seed != x.seed || i.down() == nil {
			x.perimeter += 1
		}
		if i.left() != nil && i.left().seed != x.seed || i.left() == nil {
			x.perimeter += 1
		}
		if i.right() != nil && i.right().seed != x.seed || i.right() == nil {
			x.perimeter += 1
		}
	}
}

func (x *day12region) calculateSides() {}

// fenceCost = regionCostA + regionCostB + etc
// regionCost = perimeter * area
func (x *day12garden) calculateCost(by string) int {
	var n int
	switch by {
	case "area":
		for _, i := range x.regions {
			n += i.cost
		}
	case "sides":

	}
	return n
}

func (x *day12region) calculateCost() {
	x.cost = x.perimeter * x.area
}

func (x *aocDay12) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay12) part1() {
	x.readFile()
	for _, i := range day12.garden.plants {
		if !i.isRegioned {
			p := i.ping(i.seed, []*day12plant{})
			r := &day12region{
				seed:   i.seed,
				plants: p,
			}
			r.calculateArea()
			r.calculatePerimeter()
			r.calculateCost()
			day12.garden.regions = append(day12.garden.regions, r)
		}
	}
	for _, i := range day12.garden.regions {
		i.calculateArea()
		i.calculatePerimeter()
	}
	fmt.Println("Part 1 Solution:", day12.garden.calculateCost("area"))
}

func (x *aocDay12) part2() {
	x.readFile()
	for _, i := range day12.garden.plants {
		if !i.isRegioned {
			p := i.ping(i.seed, []*day12plant{})
			r := &day12region{
				seed:   i.seed,
				plants: p,
			}
			r.calculateArea()
			r.calculatePerimeter()
			r.calculateCost()
			day12.garden.regions = append(day12.garden.regions, r)
		}
	}
	for _, i := range day12.garden.regions {
		i.calculateArea()
		i.calculateSides()
	}
	fmt.Println("Part 2 Solution:", day12.garden.calculateCost("sides"))
}
