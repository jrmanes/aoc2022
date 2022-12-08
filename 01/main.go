package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// part1 returns the biggest in the file input.txt
func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	//Search for the maximum sum of calories
	maxCalories := 0
	currentCalories := 0

	for fileScanner.Scan() {
		num := fileScanner.Text()
		t, err := strconv.Atoi(num)

		currentCalories += t

		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			//I start with a new count
			currentCalories = 0
		}
	}
	fmt.Println(maxCalories)
}

// part2
func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	maxCalories1 := 0
	maxCalories2 := 0
	maxCalories3 := 0
	currentCalories := 0

	for fileScanner.Scan() {
		num := fileScanner.Text()
		t, err := strconv.Atoi(num)

		currentCalories += t

		if err != nil {
			if currentCalories > maxCalories3 {
				maxCalories3 = currentCalories
			}
			if maxCalories3 > maxCalories2 {
				maxCalories3, maxCalories2 = maxCalories2, maxCalories3
			}
			if maxCalories2 > maxCalories1 {
				maxCalories2, maxCalories1 = maxCalories1, maxCalories2
			}

			//I start with a new count
			currentCalories = 0
		}
	}
	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)
}

// https://adventofcode.com/2022/day/1
func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
