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

	var trees [][]int
	visibleCount := 0
	for scanner.Scan() {
		txt := scanner.Text();
		treeRowStrs := strings.Split(txt, "")
		var treeRow []int
		for _, heightStr := range treeRowStrs {
			height, _ := strconv.Atoi(heightStr)
			treeRow = append(treeRow, height)
		}
		trees = append(trees, treeRow)
	}

	for rowIdx, row := range trees {
		for colIdx, tree := range row {
			if (rowIdx == 0 || rowIdx == len(row) - 1 || colIdx == 0 || colIdx == len(trees) - 1) {
				visibleCount += 1
			} else {
				// check heights from the left
				isVisible := false
				for i := 0; i < colIdx; i += 1 {
					if trees[rowIdx][i] >= tree {
						break
					}
					if (i == colIdx - 1) {
						isVisible = true
						if (isVisible) {
						}
					}
				}
				if (!isVisible) {
					// check heights to the right
					for i := 1; i < len(row) - colIdx; i += 1 {
						if trees[rowIdx][colIdx + i] >= tree {
							break
						}
						if (i == len(row) - colIdx - 1) {
							isVisible = true
						}
					}
				}
				if (!isVisible) {
					// check heights to from top
					for i := 0; i < rowIdx; i += 1 {
						if trees[i][colIdx] >= tree {
							break
						}
						if (i == rowIdx - 1) {
							isVisible = true
						}
					}
				}
				if (!isVisible) {
					// check heights to from the bottom
					for i := 1; i < len(trees) - rowIdx; i += 1 {
						if trees[rowIdx + i][colIdx] >= tree {
							break
						}
						if (i == len(trees) - rowIdx - 1) {
							isVisible = true
						}
					}
				}
				if (isVisible) {
					visibleCount += 1
				}
			}
		}
	}
	fmt.Printf("Visible trees count: %d\n", visibleCount)
}
