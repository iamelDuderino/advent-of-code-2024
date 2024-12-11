package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	runes []rune
	id    int
}

func (x day9mem) display() string {
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
	mem    []day9mem
	cache  []day9mem
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
			x.mem = append(x.mem, day9mem{
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
	defer func() { x.cache = nil }()
	x.mem = []day9mem{}
	for idx, i := range x.cache {
		n, err := strconv.Atoi(string(i.runes))
		if err != nil {
			fmt.Println(err)
			return
		}
		// even indexes are files
		// the id is the index/2 (# of files)
		// the string to add is N number of ID
		mem := []day9mem{}
		if idx%2 == 0 {
			id := []rune(fmt.Sprint(fileid))
			for i := 0; i < n; i++ {
				mem = append(mem, day9mem{
					runes: id,
					id:    fileid,
				})
			}
			fileid += 1
		} else {
			// odd files are free memory
			for i := 0; i < n; i++ {
				mem = append(mem, day9mem{
					runes: []rune{free},
				})
			}
		}
		x.mem = append(x.mem, mem...)
	}
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
		lastbit    = len(x.mem) - 1
		nextbit    = lastbit - maxbits
		currentIdx = lastbit
		rxp        = regexp.MustCompile(`^([\d]+[\.]+$)`)
		decr       = func() {
			currentIdx -= 1
			lastbit -= 1
			nextbit -= 1
		}
		reset = func() {
			lastbit = len(x.mem) - 1
			nextbit = lastbit - maxbits
			currentIdx = lastbit
		}

		done = func() bool {
			x.restring()
			return len(rxp.FindAllStringSubmatch(x.str, 1)) == 1
		}
	)
	for !done() {
		for range x.mem[nextbit:lastbit] {
			m := x.mem[currentIdx]
			if m.runes[0] == freeRune {
				decr()
				continue
			}
			if nextbit <= 0 {
				reset()
				break
			}
			for idx, ii := range x.mem[:currentIdx] {
				if ii.runes[0] == freeRune {
					// fmt.Printf("Replacing idx %d %s with %s\n", idx, x.mem[idx].read(), m.read())
					// fmt.Printf("Replacing idx %d %s with %s\n", currentIdx, x.mem[currentIdx].read(), string(freeRune))
					x.mem[idx] = m
					x.mem[currentIdx].runes = []rune{freeRune}
					break
				}
			}
			decr()
		}
	}
}

func (x *aocDay9) checksum() int {
	var n int
	for idx, i := range x.mem {
		if string(i.runes[0]) == `.` {
			break
		}
		n += (idx * i.id)
	}
	return n
}

func (x *aocDay9) part1() {
	x.readFile()
	x.allocate()
	x.defrag()
	fmt.Println("Part 1 Solution:", x.checksum())
}

func (x *aocDay9) part2() {
	var (
		answer string
	)

	fmt.Println("Part 2 Solution:", answer)
}
