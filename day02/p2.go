package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func computeScore(them, result string) int {
	score := 0
	if result == "Y" {
		score = 3
		if them == "A" {
			score += 1
		} else if them == "B" {
			score += 2
		} else {
			score += 3
		}
	} else if result == "Z" {
		score = 6
		if them == "A" {
			score += 2
		} else if them == "B" {
			score += 3
		} else {
			score += 1
		}
	} else {
		if them == "A" {
			score += 3
		} else if them == "B" {
			score += 1
		} else {
			score += 2
		}
	}

	return score
}

func main() {
	absPath, _ := filepath.Abs("input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
}

	score := 0
	for scanner.Scan() {
		txt := scanner.Text()
		shapes := strings.Split(txt, " ")
		score += computeScore(shapes[0], shapes[1])
	}	

	fmt.Printf("%d\n", score)
}
