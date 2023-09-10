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

	overlapsCount := 0
	for scanner.Scan() {
		txt := scanner.Text();
		intervals := strings.Split(txt, ",")
		interval1 := strings.Split(intervals[0], "-")
		interval2 := strings.Split(intervals[1], "-")
		i0, _ := strconv.Atoi(interval1[0])
		i1, _ := strconv.Atoi(interval1[1])
		i2, _ := strconv.Atoi(interval2[0])
		i3, _ := strconv.Atoi(interval2[1])
		if (i3 >= i0 && i1 >= i2) || (i0 >= i3 && i2 >= i1) {
			overlapsCount += 1
		}
	}

	fmt.Printf("No. of overlpas is %d\n", overlapsCount)
}
