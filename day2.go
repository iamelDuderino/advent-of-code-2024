package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var day2 = &aocDay2{
	banner: `-------------
-   DAY 2   -
-------------`,
	data: dataFolder + "day2.txt",
}

type aocDay2 struct {
	banner           string
	data             string
	day2ReportLevels []day2reportLevel
}

func (x *aocDay2) printBanner() {
	fmt.Println(x.banner)
}

type day2reportLevel struct {
	ints []int
}

func (x *aocDay2) readFile() error {
	f, err := os.Open(x.data)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		levels := strings.Split(s.Text(), " ")
		thisReport := day2reportLevel{}
		for _, level := range levels {
			thisLevelInt, err := strconv.Atoi(level)
			if err != nil {
				return err
			}
			thisReport.ints = append(thisReport.ints, thisLevelInt)
		}
		x.day2ReportLevels = append(x.day2ReportLevels, thisReport)
	}
	return nil
}

// returns the index of the conflicting level and where or not the report is safe as is
func (x *day2reportLevel) isSafeAndIncreasing() (int, bool) {
	for n, i := range x.ints {
		if n == 0 {
			continue
		}
		// The levels do not match
		if i == x.ints[n-1] {
			return n, false
		}
		// The levels are all increasing
		if i < x.ints[n-1] {
			return n, false
		}
		// And any two adjacent levels differ by at least one and at most three
		if !((i-x.ints[n-1]) > 0 && (i-x.ints[n-1]) < 4) {
			return n, false
		}
	}
	return 0, true
}

// returns the index of the conflicting level as applicable and whether or not the report is safe
func (x *day2reportLevel) isSafeAndDecreasing() (int, bool) {
	for n, i := range x.ints {
		if n == 0 {
			continue
		}
		// The levels do not match
		if i == x.ints[n-1] {
			return n, false
		}
		// The levels are all decreasing
		if i > x.ints[n-1] {
			return n, false
		}
		// And any two adjacent levels differ by at least one and at most three
		if !((x.ints[n-1]-i) > 0 && (x.ints[n-1]-i) < 4) {
			return n, false
		}
	}
	return 0, true
}

func (x *aocDay2) part1() {
	var (
		numSafeReports int
	)
	err := day2.readFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range day2.day2ReportLevels {
		_, safeIncr := i.isSafeAndIncreasing()
		_, safeDecr := i.isSafeAndDecreasing()
		if safeIncr || safeDecr {
			numSafeReports += 1
		}
	}
	fmt.Println("Day 2, Part 1 Solution:", numSafeReports)
}

func (x *aocDay2) part2() {
	var (
		trueNumSafeReports int
	)
	for _, i := range day2.day2ReportLevels {

		faultyIndexIncr, safeIncr := i.isSafeAndIncreasing()
		if safeIncr {
			trueNumSafeReports += 1
			continue
		}
		if x.tolerable(true, i, faultyIndexIncr) {
			trueNumSafeReports += 1
			continue
		} else if faultyIndexIncr > 0 && x.tolerable(true, i, faultyIndexIncr-1) {
			trueNumSafeReports += 1
			continue
		} else if faultyIndexIncr+1 <= len(i.ints) && x.tolerable(true, i, faultyIndexIncr+1) {
			trueNumSafeReports += 1
			continue
		}

		// --------------------------------------------------------

		faultyIndexDecr, safeDecr := i.isSafeAndDecreasing()
		if safeDecr {
			trueNumSafeReports += 1
			continue
		}
		if x.tolerable(false, i, faultyIndexDecr) {
			trueNumSafeReports += 1
			continue
		} else if faultyIndexDecr > 0 && x.tolerable(false, i, faultyIndexDecr-1) {
			trueNumSafeReports += 1
			continue
		} else if faultyIndexDecr+1 <= len(i.ints) && x.tolerable(false, i, faultyIndexDecr+1) {
			trueNumSafeReports += 1
			continue
		}

	}

	fmt.Println("Day 2, Part 2 Solution:", trueNumSafeReports)
}

func (x *aocDay2) tolerable(increasing bool, level day2reportLevel, faultyIndex int) bool {
	var updatedLevel day2reportLevel
	for n, i := range level.ints {
		if n == faultyIndex {
			continue
		}
		updatedLevel.ints = append(updatedLevel.ints, i)
	}
	var nowSafe bool
	switch increasing {
	case true:
		_, nowSafe = updatedLevel.isSafeAndIncreasing()
	case false:
		_, nowSafe = updatedLevel.isSafeAndDecreasing()
	}
	return nowSafe
}
