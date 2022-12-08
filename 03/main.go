package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// https://adventofcode.com/2022/day/3
func part1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	sc := bufio.NewScanner(readFile)

	var sumOfPriorities int

	for sc.Scan() {
		items := make(map[rune]bool)
		// Elements of the first compartment
		for _, itemLeftPart := range sc.Text()[:len(sc.Text())/2] {
			items[itemLeftPart] = true
		}

		// Elements of the second compartment
		for _, itemRightPart := range sc.Text()[len(sc.Text())/2:] {
			// If an item is in the first set it's in both compartments
			if items[itemRightPart] {
				sumOfPriorities += int(unicode.ToLower(itemRightPart) - 96)
				if unicode.IsUpper(itemRightPart) {
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	sc := bufio.NewScanner(readFile)

	var sumOfPriorities int

	for sc.Scan() {
		// Create three sets with the elements of each elf
		itemsFirst := createSetOfItems(sc.Text())
		sc.Scan()
		itemsSecond := createSetOfItems(sc.Text())
		sc.Scan()
		itemsThird := createSetOfItems(sc.Text())

		// For each item of the first (or second or third, it's the same) elf we check if the other two elves have that item too
		for itemFirstElf := range itemsFirst {
			if itemsSecond[itemFirstElf] && itemsThird[itemFirstElf] {
				sumOfPriorities += int(unicode.ToLower(itemFirstElf) - 96)
				if unicode.IsUpper(itemFirstElf) {
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}

func createSetOfItems(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}

func main() {
	part1()
	part2()
}
