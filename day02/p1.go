package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func computeScore(them, me string) int {
	score := 1
	if me == "Y" {
		score = 2
	} else if me == "Z" {
		score = 3
	}
	if (me == "X" && them == "A") || (me == "Y" && them == "B") || (me == "Z" && them == "C") {
		score += 3
	} else if (me == "X" && them == "C") || (me == "Y" && them == "A" ) || (me == "Z" && them == "B") {
		score += 6
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
