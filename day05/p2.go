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

	
	var cratesInv [][]string // each array represents a row of crates from left to right
	var crates [][]string // each array represents a column of crates from bottom to top
	cratesStrDone := false
	for scanner.Scan() {
		txt := scanner.Text();
		if (!strings.HasPrefix(txt, "move") && strings.Contains(txt, "1")) {
			cratesStrDone = true
			for row := len(cratesInv) - 1; row >= 0; row -= 1 {
				if (row == len(cratesInv) - 1) {
					crates = make([][]string, len(cratesInv[row]), len(cratesInv[row]) * len(cratesInv[row]))
				}
				for col := 0; col < len(cratesInv[row]); col += 1 {
					crates[col] = append(crates[col], cratesInv[row][col])
				}
			}
		}
		if !cratesStrDone {
			var crateLine []string
			txtSplit := strings.Split(txt, "")
			for j := 1; j < len(txtSplit); j += 4 {
				crateLine = append(crateLine, txtSplit[j])
			}
			cratesInv = append(cratesInv, crateLine)
		}
		if strings.HasPrefix(txt, "move") {
			moveText := strings.Split(txt, " ")
			moveCount, _ := strconv.Atoi(moveText[1])
			moveFrom, _ := strconv.Atoi(moveText[3])
			moveTo, _ := strconv.Atoi(moveText[5])
			
			lastCrateFromIdx := 0
			for l := 0; l < len(crates[moveFrom - 1]); l += 1 {
				if (crates[moveFrom - 1][l] == " ") {
					break
				} else {
					lastCrateFromIdx = l
				}
			}

			for k := moveCount; k > 0; k -= 1 {
				letterToMove := crates[moveFrom - 1][lastCrateFromIdx - k + 1]
				// what is the last element of the source column that is not an empty string?
				crates[moveFrom - 1][lastCrateFromIdx - k + 1] = " "
				// what is the last element of the destination column that is an empty string?
				shouldAppend := true
				for m := 0; m < len(crates[moveTo - 1]); m += 1 {
					if (crates[moveTo - 1][m] == " ") {
						crates[moveTo - 1][m] = letterToMove
						shouldAppend = false
						break
					}
				}

				if (shouldAppend) {
					crates[moveTo - 1] = append(crates[moveTo - 1], letterToMove)
				}
			}
		}
	}
	for n := 0; n < len(crates); n += 1 {
		for p := 0; p < len(crates[n]); p += 1 {
			if (p == len(crates[n]) - 1) {
				fmt.Printf("%s", crates[n][p])
			}
			if (p > 0 && crates[n][p] == " ") {
				fmt.Printf("%s", crates[n][p - 1])
				break
			}
		}
	}
}
