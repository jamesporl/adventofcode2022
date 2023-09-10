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
	maxCals := [3]int{0, 0, 0}
	minInMaxCals := 0
	minInMaxCalsIdx := 0
	for scanner.Scan() {
		txt := scanner.Text();
		if (txt == "") {
			fmt.Printf("%d\n", currentCals)
			if (currentCals >= minInMaxCals) {
				maxCals[minInMaxCalsIdx] = currentCals;

				minInMaxCals = maxCals[0]
				minInMaxCalsIdx = 0
				for i := 0; i < 3; i++ {
					if (maxCals[i] < minInMaxCals) {
						minInMaxCalsIdx = i
						minInMaxCals = maxCals[i]
					}
				}
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

	totalMaxCals := 0;
	for i := 0; i < 3; i++ {
		totalMaxCals += maxCals[i]
	}
	
	fmt.Printf("Sum of max cals is %d\n", totalMaxCals)
}
