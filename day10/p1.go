package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

func computeTotalStrength(cycles int, x int) int {
	if cycles % 40 == 20 {
		return cycles * x
	}
	return 0
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

	cycles := 0
	totalStrength := 0
	x := 1
	for scanner.Scan() {
		txt := scanner.Text();
		
		if (txt == "noop") {
			cycles += 1
			totalStrength += computeTotalStrength(cycles, x)
		} else {
			cycles += 1
			totalStrength += computeTotalStrength(cycles, x)
			moveArr := strings.Split(txt, " ")
			toAdd, _ := strconv.Atoi(moveArr[1])
			cycles += 1
			totalStrength += computeTotalStrength(cycles, x)
			x += toAdd
		}
	}
	fmt.Printf("Total strength %d\n", totalStrength)
}
