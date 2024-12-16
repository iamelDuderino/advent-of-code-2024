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
	x, y                                  int
	seed                                  string
	isRegioned                            bool
	regionId                              int
	upside, downside, leftside, rightside bool
}

type day12region struct {
	id                     int
	seed                   string
	plants                 []*day12plant
	edges, insides         []*day12plant
	area, perimeter, sides int
	cost                   int
}

// marks sides as traversible depending on whether or not a plant
// can be found on that side of the plant within the same region
func (x *day12region) markSides() {
	for _, i := range x.plants {
		// up
		if x.get(i.x, i.y-1) == nil {
			i.upside = true
		}
		// down
		if x.get(i.x, i.y+1) == nil {
			i.downside = true
		}
		// left
		if x.get(i.x-1, i.y) == nil {
			i.leftside = true
		}
		// right
		if x.get(i.x+1, i.y) == nil {
			i.rightside = true
		}
	}
}

func (x *day12region) checkRightAngle(p *day12plant, dir string) bool {
	var inRegion = func(p *day12plant) bool {
		for _, i := range x.plants {
			if i.x == p.x && i.y == p.y {
				return true
			}
		}
		return false
	}
	switch dir {
	case "up":
		if p.upleft() != nil && inRegion(p.upleft()) {
			return true
		}
	case "down":
		if p.downright() != nil && inRegion(p.downright()) {
			return true
		}
	case "right":
		if p.upright() != nil && inRegion(p.upright()) {
			return true
		}
	case "left":
		if p.downleft() != nil && inRegion(p.downleft()) {
			return true
		}
	}
	return false
}

func (x *day12region) isDeadEnd(p *day12plant, dir string) bool {
	// must have at least 3 nil sides containing no members
	var n int
	up := x.get(p.x, p.y-1)
	down := x.get(p.x, p.y+1)
	left := x.get(p.x-1, p.y)
	right := x.get(p.x+1, p.y)
	if up == nil {
		n += 1
	}
	if down == nil {
		n += 1
	}
	if left == nil {
		n += 1
	}
	if right == nil {
		n += 1
	}
	// must have 0 corner members in the intended direction
	var nn int
	upleft := x.get(p.x-1, p.y-1)
	upright := x.get(p.x+1, p.y-1)
	downleft := x.get(p.x-1, p.y+1)
	downright := x.get(p.x+1, p.y+1)
	switch dir {
	case "up":
		if upleft != nil {
			nn += 1
		}
		if upright != nil {
			nn += 1
		}
	case "down":
		if downleft != nil {
			nn += 1
		}
		if downright != nil {
			nn += 1
		}
	case "left":
		if upleft != nil {
			nn += 1
		}
		if downleft != nil {
			nn += 1
		}
	case "right":
		if upright != nil {
			nn += 1
		}
		if downright != nil {
			nn += 1
		}
	}
	// do a final check to determine if there are any plants to the sides of the intended
	// direction or now in front of the recently turned tracing angles intended direction
	var nnn int
	switch dir {
	case "up", "down":
		if dir == "up" {
			if x.get(p.x, p.y-1) != nil {
				nnn += 1
			}
		}
		if dir == "down" {
			if x.get(p.x, p.y+1) != nil {
				nnn += 1
			}
		}
		if x.get(p.x-1, p.y) != nil || x.get(p.x+1, p.y) != nil {
			nnn += 1
		}
	case "left", "right":
		if dir == "left" {
			if x.get(p.x-1, p.y) != nil {
				nnn += 1
			}
		}
		if dir == "right" {
			if x.get(p.x+1, p.y+1) != nil {
				nnn += 1
			}
		}
		if x.get(p.x, p.y-1) != nil || x.get(p.x, p.y+1) != nil {
			nnn += 1
		}
	}
	return (n == 3 && nn == 0 && nnn == 0)
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

// day12region.traceInsides() iterates through insides, checking
// for internal sub-regions that must be included
func (x *day12region) traceInsides() {
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
					r.trace()
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
					r.trace()
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
					r.trace()
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
					r.trace()
				}
				x.sides += r.sides
			}
		}
	}
}

func (x *day12region) trace() {
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
		tracing      bool = true
		iters        int
		start        *day12plant = x.plants[0]
		currentPlant *day12plant = x.plants[0]
		dir          string      = right
		surface      string      = up
		inDeadend    bool
	)
	x.markSides()
	for tracing {
		// fmt.Printf("%s (%d,%d), sides: %d\n", dir, currentPlant.x, currentPlant.y, x.sides)
		if iters >= 3 && currentPlant.x == start.x && currentPlant.y == start.y && dir == up {
			tracing = false
		}
		switch dir {
		case up: // ra = upleft
			if x.isDeadEnd(currentPlant, dir) && !inDeadend && iters > 1 && currentPlant.x != start.x && currentPlant.y != start.y {
				switch surface {
				case right:
					surface = left
				case left:
					surface = right
				}
				dir = down
				x.sides += 2
				inDeadend = true
			} else if currentPlant.up() != nil && currentPlant.up().seed == x.seed {
				inDeadend = false
				var traversible bool
				switch surface {
				case right:
					traversible = currentPlant.up().rightside
				case left:
					traversible = currentPlant.up().leftside
				}
				if traversible {
					currentPlant = currentPlant.up()
				} else {
					ra := x.checkRightAngle(currentPlant, dir)
					if ra {
						currentPlant = currentPlant.upleft()
						surface = down
						dir = left
					} else {
						surface = up
						dir = right
					}
					x.sides += 1
				}
			} else {
				inDeadend = false
				ra := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = currentPlant.upleft()
					surface = down
					dir = left
				} else {
					surface = up
					dir = right
				}
				x.sides += 1
			}
		case down: // ra = downright
			if x.isDeadEnd(currentPlant, dir) && !inDeadend && iters > 1 && currentPlant.x != start.x && currentPlant.y != start.y {
				switch surface {
				case right:
					surface = left
				case left:
					surface = right
				}
				dir = up
				x.sides += 2
				inDeadend = true
			} else if currentPlant.down() != nil && currentPlant.down().seed == x.seed {
				inDeadend = false
				var traversible bool
				switch surface {
				case right:
					traversible = currentPlant.down().rightside
				case left:
					traversible = currentPlant.down().leftside
				}
				if traversible {
					currentPlant = currentPlant.down()
				} else {
					ra := x.checkRightAngle(currentPlant, dir)
					if ra {
						currentPlant = currentPlant.downright()
						surface = up
						dir = right
					} else {
						surface = down
						dir = left
					}
					x.sides += 1
				}
			} else {
				inDeadend = false
				ra := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = currentPlant.downright()
					surface = up
					dir = right
				} else {
					surface = down
					dir = left
				}
				x.sides += 1
			}
		case left: // ra = downleft
			if x.isDeadEnd(currentPlant, dir) && !inDeadend && iters > 1 && currentPlant.x != start.x && currentPlant.y != start.y {
				switch surface {
				case up:
					surface = down
				case down:
					surface = up
				}
				dir = right
				x.sides += 2
				inDeadend = true
			} else if currentPlant.left() != nil && currentPlant.left().seed == x.seed {
				inDeadend = false
				var traversible bool
				switch surface {
				case up:
					traversible = currentPlant.left().upside
				case down:
					traversible = currentPlant.left().downside
				}
				if traversible {
					currentPlant = currentPlant.left()
				} else {
					ra := x.checkRightAngle(currentPlant, dir)
					if ra {
						currentPlant = currentPlant.downleft()
						surface = right
						dir = down
					} else {
						surface = left
						dir = up
					}
					x.sides += 1
				}
			} else {
				inDeadend = false
				ra := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = currentPlant.downleft()
					surface = left
					dir = down
				} else {
					surface = left
					dir = up
				}
				x.sides += 1
			}
		case right: // ra = upright
			if x.isDeadEnd(currentPlant, dir) && !inDeadend && iters > 1 && currentPlant.x != start.x && currentPlant.y != start.y {
				switch surface {
				case up:
					surface = down
				case down:
					surface = up
				}
				dir = left
				x.sides += 2
				inDeadend = true
			} else if currentPlant.right() != nil && currentPlant.right().seed == x.seed {
				inDeadend = false
				var traversible bool
				switch surface {
				case up:
					traversible = currentPlant.right().upside
				case down:
					traversible = currentPlant.right().downside
				}
				if traversible {
					currentPlant = currentPlant.right()
				} else {
					ra := x.checkRightAngle(currentPlant, dir)
					if ra {
						currentPlant = currentPlant.upright()
						surface = left
						dir = up
					} else {
						surface = right
						dir = down
					}
					x.sides += 1
				}
			} else {
				inDeadend = false
				ra := x.checkRightAngle(currentPlant, dir)
				if ra {
					currentPlant = currentPlant.upright()
					surface = left
					dir = up
				} else {
					surface = right
					dir = down
				}
				x.sides += 1
			}
		}
		iters += 1
		x.edges = append(x.edges, currentPlant)
		if iters > 200 {
			x.draw()
			msg := fmt.Sprintf("Region %d is looping @ (%d,%d)", x.id, currentPlant.x, currentPlant.y)
			panic(msg)
		}
	}
	x.fillInsides()
	x.traceInsides()
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

func (x *day12plant) upleft() *day12plant {
	if day12.get(x.x-1, x.y-1) == nil {
		return nil
	}
	return day12.get(x.x-1, x.y-1)
}

func (x *day12plant) upright() *day12plant {
	if day12.get(x.x+1, x.y-1) == nil {
		return nil
	}
	return day12.get(x.x+1, x.y-1)
}

func (x *day12plant) downleft() *day12plant {
	if day12.get(x.x-1, x.y+1) == nil {
		return nil
	}
	return day12.get(x.x-1, x.y+1)
}

func (x *day12plant) downright() *day12plant {
	if day12.get(x.x+1, x.y+1) == nil {
		return nil
	}
	return day12.get(x.x+1, x.y+1)
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
		i.trace()
		i.calculateCost("sides")
		//
		// i.draw()
		// fmt.Println("Area:", i.area)
		// fmt.Println("Sides:", i.sides)
		// fmt.Println("Cost:", i.cost)
	}
	fmt.Println("Part 2 Solution:", day12.garden.calculateCost())
}
