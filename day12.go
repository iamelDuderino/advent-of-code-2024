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
	regionId   int
}

type day12region struct {
	id                     int
	seed                   string
	plants                 []*day12plant
	edges, insides         []*day12plant
	area, perimeter, sides int
	cost                   int
}

// a right angle is checked counter clockwise of the intended direction
func (x *day12region) checkRightAngle(p *day12plant, dir string) (bool, *day12plant) {
	switch dir {
	case "up":
		r := x.get(p.x-1, p.y-1)
		if r != nil {
			return true, r
		}
	case "down":
		r := x.get(p.x+1, p.y+1)
		if r != nil {
			return true, r
		}
	case "right":
		r := x.get(p.x+1, p.y-1)
		if r != nil {
			return true, r
		}
	case "left":
		r := x.get(p.x-1, p.y+1)
		if r != nil {
			return true, r
		}
	}
	return false, nil
}

func (x *day12region) get(xx, yy int) *day12plant {
	for _, i := range x.plants {
		if i.x == xx && i.y == yy {
			return i
		}
	}
	return nil
}

func (x *day12region) draw() {
	var inRegion = func(p *day12plant) bool {
		for _, i := range x.plants {
			if i.x == p.x && i.y == p.y {
				return true
			}
		}
		return false
	}
	for _, i := range day12.garden.plants {
		if inRegion(i) {
			fmt.Print(i.seed)
		} else {
			fmt.Print(`.`)
		}
		if i.x == day12.garden.width {
			fmt.Println()
		}
	}
}

func (x *day12region) inEdges(p *day12plant) bool {
	for _, i := range x.edges {
		if i.x == p.x && i.y == p.y {
			return true
		}
	}
	return false
}

func (x *day12region) fillInsides() {
	for _, i := range x.plants {
		if !x.inEdges(i) {
			x.insides = append(x.insides, i)
		}
	}
}

func (x *day12garden) getRegion(id int) *day12region {
	for _, i := range x.regions {
		if i.id == id {
			return i
		}
	}
	return nil
}

// day12region.calculateInsides() iterates through insides, checking
// for internal sub-regions that must be included in the total sides count
func (x *day12region) calculateInsides() {
	var (
		regionsFound   = []int{}
		inRegionsFound = func(n int) bool {
			for _, i := range regionsFound {
				if i == n {
					return true
				}
			}
			return false
		}
	)
	for _, i := range x.insides {
		if i.right() != nil {
			p := i.right()
			if p.regionId != x.id && !inRegionsFound(p.regionId) {
				regionsFound = append(regionsFound, p.regionId)
				r := day12.garden.getRegion(p.regionId)
				if r.sides == 0 {
					r.calculateSides()
				}
				x.sides += r.sides
			}
		}
		if i.left() != nil {
			p := i.left()
			if p.regionId != x.id && !inRegionsFound(p.regionId) {
				regionsFound = append(regionsFound, p.regionId)
				r := day12.garden.getRegion(p.regionId)
				if r.sides == 0 {
					r.calculateSides()
				}
				x.sides += r.sides
			}
		}
		if i.up() != nil {
			p := i.up()
			if p.regionId != x.id && !inRegionsFound(p.regionId) {
				regionsFound = append(regionsFound, p.regionId)
				r := day12.garden.getRegion(p.regionId)
				if r.sides == 0 {
					r.calculateSides()
				}
				x.sides += r.sides
			}
		}
		if i.down() != nil {
			p := i.down()
			if p.regionId != x.id && !inRegionsFound(p.regionId) {
				regionsFound = append(regionsFound, p.regionId)
				r := day12.garden.getRegion(p.regionId)
				if r.sides == 0 {
					r.calculateSides()
				}
				x.sides += r.sides
			}
		}
	}
}

func (x *day12region) calculateSides() {
	if x.sides > 0 {
		return
	}
	const (
		up    = "up"
		right = "right"
		down  = "down"
		left  = "left"
	)
	var (
		iters        int
		start        *day12plant = x.plants[0]
		currentPlant *day12plant = x.plants[0]
		nextPlant    *day12plant
		dir          string = right
		// in order to move "forward", must not have a counter clockwise neighbor
		hasNeighbor = func(p *day12plant, dir string) bool {
			switch dir {
			case right:
				if p.y == 0 {
					return false
				}
				return x.get(p.x, p.y-1) != nil
			case down:
				if p.x == day12.garden.width {
					return false
				}
				return x.get(p.x+1, p.y) != nil
			case left:
				if p.y == day12.garden.length {
					return false
				}
				return x.get(p.x, p.y+1) != nil
			case up:
				return x.get(p.x-1, p.y) != nil
			}
			return false
		}
	)
	for {
		// fmt.Printf("(%d,%d) %s\n", currentPlant.x, currentPlant.y, dir)
		if currentPlant.x == start.x && currentPlant.y == start.y && dir == up {
			x.sides += 1
			break
		}
		if hasNeighbor(currentPlant, dir) {
			x.sides += 1
			switch dir {
			case right:
				dir = up
			case down:
				dir = right
			case left:
				dir = down
			case up:
				dir = left
			}
		}
		switch dir {
		case right:
			nextPlant = x.get(currentPlant.x+1, currentPlant.y)
			if nextPlant != nil {
				currentPlant = nextPlant
			} else {
				ra, p := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = p
					dir = up
					x.sides += 1
				} else {
					dir = down
					x.sides += 1
				}
			}
		case down:
			nextPlant = x.get(currentPlant.x, currentPlant.y+1)
			if nextPlant != nil {
				currentPlant = nextPlant
			} else {
				ra, p := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = p
					dir = right
					x.sides += 1
				} else {
					dir = left
					x.sides += 1
				}
			}
		case left:
			nextPlant = x.get(currentPlant.x-1, currentPlant.y)
			if nextPlant != nil {
				currentPlant = nextPlant
			} else {
				ra, p := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = p
					dir = down
					x.sides += 1
				} else {
					dir = up
					x.sides += 1
				}
			}
		case up:
			nextPlant = x.get(currentPlant.x, currentPlant.y-1)
			if nextPlant != nil {
				currentPlant = nextPlant
			} else {
				ra, p := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = p
					dir = left
					x.sides += 1
				} else {
					dir = right
					x.sides += 1
				}
			}

		}
		x.edges = append(x.edges, currentPlant)
		iters += 1
		if iters > 500 {
			x.draw()
			msg := fmt.Sprintf("Region %d is looping @ (%d,%d)", x.id, currentPlant.x, currentPlant.y)
			panic(msg)
		}
	}
	x.fillInsides()
	x.calculateInsides()
}

func (x *day12region) calculateArea() {
	if x.area > 0 {
		return
	}
	x.area = len(x.plants)
}

func (x *day12region) calculatePerimeter() {
	if x.perimeter > 0 {
		return
	}
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

func (x *day12region) calculateCost(by string) {
	switch by {
	case "perimeter":
		x.cost = x.area * x.perimeter
	case "sides":
		x.cost = x.area * x.sides
	}
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

// fenceCost = regionCostA + regionCostB + etc
//
//	part1: regionCost = perimeter * area
//	part2: regionCost = # of sides * area
func (x *day12garden) calculateCost() int {
	var n int
	for _, i := range x.regions {
		n += i.cost
	}
	return n
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
			day12.garden.regions = append(day12.garden.regions, r)
		}
	}
	for _, i := range day12.garden.regions {
		i.calculateArea()
		i.calculatePerimeter()
		i.calculateCost("perimeter")
	}
	fmt.Println("Part 1 Solution:", day12.garden.calculateCost())
}

func (x *aocDay12) part2() {
	x.readFile()
	for n, i := range day12.garden.plants {
		if !i.isRegioned {
			p := i.ping(i.seed, []*day12plant{})
			r := &day12region{
				id:     n,
				seed:   i.seed,
				plants: p,
			}
			for _, i := range p {
				i.regionId = n
			}
			day12.garden.regions = append(day12.garden.regions, r)
		}
	}
	for n, i := range day12.garden.regions {
		fmt.Printf("Calculating Region %d/%d\n", n+1, len(day12.garden.regions))
		i.calculateArea()
		i.calculateSides()
		i.calculateCost("sides")
		//
		i.draw()
		fmt.Println("Area:", i.area)
		fmt.Println("Sides:", i.sides)
		fmt.Println("Cost:", i.cost)
		// break
	}
	fmt.Println("Part 2 Solution:", day12.garden.calculateCost())
}
