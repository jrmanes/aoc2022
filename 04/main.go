package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://adventofcode.com/2022/day/4
func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	sc := bufio.NewScanner(readFile)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}

func main() {
	part1()
}
