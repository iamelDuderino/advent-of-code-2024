package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var (
	day9 = &aocDay9{
		banner: `-------------
-   DAY 9   -
-------------`,
		data: dataFolder + "day9.txt",
	}
)

type day9mem struct {
	runes    []rune
	id       int
	hasmoved bool
}

func (x *day9mem) display() string {
	var s string
	for _, i := range x.runes {
		s += string(i)
	}
	return s
}

type aocDay9 struct {
	banner string
	data   string
	str    string
	mem    []*day9mem
	cache  []*day9mem
	p2     bool
}

func (x *aocDay9) printBanner() {
	fmt.Println(x.banner)
}

func (x *aocDay9) display() {
	var s string
	for _, i := range x.mem {
		for _, r := range i.runes {
			s += string(r)
		}
	}
	fmt.Println(s)
}

func (x *aocDay9) readFile() {
	f, err := os.Open(x.data)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		for _, i := range s.Text() {
			x.mem = append(x.mem, &day9mem{
				runes: []rune{i},
			})
		}
	}
	x.restring()
}

func (x *aocDay9) allocate() {
	var (
		free   = []rune(`.`)[0]
		fileid int
	)
	x.cache = x.mem
	x.mem = []*day9mem{}
	for idx, i := range x.cache {
		n, err := strconv.Atoi(string(i.runes))
		if err != nil {
			fmt.Println(err)
			return
		}
		// even indexes are files
		// the id is the index/2 (# of files)
		// the string to add is N number of ID
		mem := []*day9mem{}
		if idx%2 == 0 {
			id := []rune(fmt.Sprint(fileid))
			for i := 0; i < n; i++ {
				mem = append(mem, &day9mem{
					runes: id,
					id:    fileid,
				})
			}
			fileid += 1
		} else {
			// odd files are free memory
			for i := 0; i < n; i++ {
				mem = append(mem, &day9mem{
					runes: []rune{free},
				})
			}
		}
		x.mem = append(x.mem, mem...)
	}
	x.cache = x.mem
	x.restring()
}

func (x *aocDay9) restring() {
	var s string
	for _, i := range x.mem {
		for _, r := range i.runes {
			s += string(r)
		}
	}
	x.str = s
}

func (x *aocDay9) defrag() {
	var (
		freeString = `.`
		freeRune   = []rune(freeString)[0]
		maxbits    = len(x.mem) / 4
		startidx   int
		lastidx    = len(x.mem) - 1
		nextidx    = lastidx - maxbits
		currentidx = lastidx
		currentid  int
		rxp        = regexp.MustCompile(`^([\d]+[\.]+$)`)
		p2done     bool
		filesmoved int
		decr       = func() {
			currentidx -= 1
			lastidx -= 1
			nextidx -= 1
		}
		reset = func() {
			lastidx = len(x.mem) - 1
			nextidx = lastidx - maxbits
			currentidx = lastidx
		}
		p1done = func() bool {
			x.restring()
			return len(rxp.FindAllStringSubmatch(x.str, 1)) == 1
		}
	)
	defer x.restring()
	if x.p2 {
		x.cache = []*day9mem{}
		for _, i := range x.mem {
			if i.id > currentid {
				currentid = i.id
			}
		}
		lastidx = len(x.mem) - 1
		for _, i := range x.mem {
			x.cache = append(x.cache, &day9mem{
				runes: i.runes,
				id:    i.id,
			})
		}
		slices.Reverse(x.cache)
	}
	for !p1done() && !x.p2 {
		// part 1
		for range x.mem[nextidx:lastidx] {
			m := x.mem[currentidx]
			if m.runes[0] == freeRune {
				decr()
				continue
			}
			if nextidx <= 0 {
				reset()
				break
			}
			for idx, ii := range x.mem[:currentidx] {
				if ii.runes[0] == freeRune {
					// fmt.Printf("Replacing idx %d %s with %s\n", idx, x.mem[idx].read(), m.read())
					// fmt.Printf("Replacing idx %d %s with %s\n", currentidx, x.mem[currentidx].read(), string(freeRune))
					x.mem[idx] = m
					x.mem[currentidx].runes = []rune{freeRune}
					break
				}
			}
			decr()
		}
	}

	// part2
	for x.p2 && !p2done {
		var (
			contig bool
			set    = []*day9mem{}
		)

		// Prep File(s)
		for idx, mem := range x.cache[startidx:] {
			fmt.Printf("IDX: %d | ID: %d | Files Moved: %d\n", startidx, currentid, filesmoved)
			if mem.id == currentid {
				set = append(set, mem)
				contig = true
			} else {
				contig = false
			}
			if !contig && len(set) > 0 {
				startidx = startidx + idx
				break
			}
		}

		// Move File(s)
		if len(set) > 0 {
			start := len(x.mem) - startidx
			end := start + len(set)
			ok, free := x.canMove(len(set), start)
			if ok {
				x.move(free, start, end, set...)
				filesmoved += len(set)
				// fmt.Printf("IDX: %d | ID: %d | %d | Moving to %d\n", currentidx, currentid, len(set), free)
			}
			currentid -= 1
			continue
		}
		if currentid <= 0 {
			p2done = true
		}
	}
}

// idxa = index of first available free memory
// idxb = index of mems... starting point
// idxc = index of mems... ending point
func (x *aocDay9) move(idxa, idxb, idxc int, mems ...*day9mem) {
	var dotm = &day9mem{
		runes: []rune(`.`),
	}
	for n, m := range mems {
		x.mem[idxa+n] = m
	}
	for n := idxb; n < idxc; n++ {
		x.mem[n] = dotm
	}
}

func (x *aocDay9) canMove(l, maxidx int) (bool, int) {
	var (
		n      int
		idxa   int
		contig bool
	)
	if l == 0 {
		return false, 0
	}
	for idxb, i := range x.mem {
		switch i.display() == `.` {
		case true:
			if n == 0 {
				idxa = idxb
				contig = true
			}
			n += 1
		case false:
			if n > 0 {
				n = 0
			}
		}
		if idxb == maxidx {
			return false, idxa
		}
		if n == l {
			break
		}
	}
	return (contig && n == l), idxa
}

func (x *aocDay9) checksum() int {
	var n int
	for idx, i := range x.mem {
		if string(i.runes[0]) == `.` {
			if x.p2 {
				continue
			} else {
				break
			}
		}
		n += (idx * i.id)
	}
	return n
}

func (x *aocDay9) part1() {
	x.readFile()
	x.allocate()
	x.defrag()
	x.display()
	fmt.Println("Part 1 Solution:", x.checksum())
}

func (x *aocDay9) part2() {
	x.p2 = true // set flag for defrag
	x.readFile()
	x.allocate()
	x.defrag()
	x.display()
	fmt.Println("Part 2 Solution:", x.checksum())
}
