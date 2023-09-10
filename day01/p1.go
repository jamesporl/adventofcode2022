package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

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

	currentCals := 0
	maxCals := 0
	for scanner.Scan() {
		txt := scanner.Text();
		if (txt == "") {
			if (currentCals > maxCals) {
				maxCals = currentCals
			}
			currentCals = 0
		} else {
			cals, err := strconv.Atoi(txt)
			if err != nil {
				log.Fatal(err)
			}
			currentCals += cals
		}
	}
	
	fmt.Printf("Max cals is %d\n", maxCals)
}
