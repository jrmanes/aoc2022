package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://adventofcode.com/2022/day/2
func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	//A : X Rock
	//B : Y Paper
	//C : Z Scissors

	var score1 int
	var score2 int
	scores1 := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}
	scores2 := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4, "B Y": 5, "C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}

	for fileScanner.Scan() {
		score1 += scores1[fileScanner.Text()]
		score2 += scores2[fileScanner.Text()]
	}

	fmt.Println(score1)
	fmt.Println(score2)
}
