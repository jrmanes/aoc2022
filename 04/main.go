package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Filename = "example"
	s.Mode ^= scanner.SkipComments  // don't skip comments
	s.Whitespace |= 1<<'-' | 1<<',' // skip minus and comma

	res := 0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		a, _ := strconv.Atoi(s.TokenText())
		//s.Scan() // skip -
		s.Scan()
		b, _ := strconv.Atoi(s.TokenText())
		//s.Scan() // skip ,
		s.Scan()
		c, _ := strconv.Atoi(s.TokenText())
		//s.Scan() // skip -
		s.Scan()
		d, _ := strconv.Atoi(s.TokenText())
		// check inclusion
		if (a <= c && d <= b) || (c <= a && b <= d) {
			res++
		}
	}
	return res
}

func Part2(input string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Whitespace |= 1<<'-' | 1<<',' // skip minus and comma
	res := 0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		a, _ := strconv.Atoi(s.TokenText())
		s.Scan()
		b, _ := strconv.Atoi(s.TokenText())
		s.Scan()
		c, _ := strconv.Atoi(s.TokenText())
		s.Scan()
		d, _ := strconv.Atoi(s.TokenText())
		// check overlap
		if !(b < c || d < a) {
			res++
		}
	}
	return res
}

// https://adventofcode.com/2022/day/4
func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start2 := time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start2))
}
