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

func computePixels(pixels [240]string, cycles int, x int) [240]string {
	row := cycles / 40
	if (40 * row + x - 1 == cycles) {
		pixels[40 * row + x - 1] = "#"
	} else if (40 * row + x == cycles) {
		pixels[40 * row + x] = "#"
	} else if (40 * row + x + 1 == cycles) {
		pixels[40 * row + x + 1] = "#"
	}
	return pixels
}

func printLine(line []string) {
	for  _, elt := range line {
		if (elt == "") {
			fmt.Print(".")
		} else {
			fmt.Print("#")
		}
	}
	fmt.Print("\n")
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
	pixels := [240]string{}
	x := 1
	for scanner.Scan() {
		txt := scanner.Text();
		
		if (txt == "noop") {
			pixels = computePixels(pixels, cycles, x)
			cycles += 1
			
		} else {
			pixels = computePixels(pixels, cycles, x)
			cycles += 1
			
			moveArr := strings.Split(txt, " ")
			toAdd, _ := strconv.Atoi(moveArr[1])
			pixels = computePixels(pixels, cycles, x)
			cycles += 1
			x += toAdd
		}
	}
	
	for i := 0; i < 6; i += 1 {
		printLine(pixels[40 * i:40 * (i + 1) - 1])
	}
}
