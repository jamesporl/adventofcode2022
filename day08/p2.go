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
	maxScore := 0
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
			if !(rowIdx == 0 || rowIdx == len(row) - 1 || colIdx == 0 || colIdx == len(trees) - 1) {
				// check heights from the left
				leftScore := 0
				for i := colIdx - 1; i >= 0; i -= 1 {
					leftScore += 1
					if trees[rowIdx][i] >= tree {
						break
					} 
				}

				// check heights to the right
				rightScore := 0
				for i := 1; i < len(row) - colIdx; i += 1 {
					rightScore += 1
					if trees[rowIdx][colIdx + i] >= tree {
						break
					}
				}

				// check heights to from top
				topScore := 0
				for i := rowIdx - 1; i >=0; i -= 1 {
					topScore += 1
					if trees[i][colIdx] >= tree {
						break
					}
				}

				
				// check heights to from the bottom
				bottomScore := 0
				for i := 1; i < len(trees) - rowIdx; i += 1 {
					bottomScore += 1
					if trees[rowIdx + i][colIdx] >= tree {
						break
					}
				}

				currentScore := leftScore * rightScore * topScore * bottomScore
				if (currentScore >= maxScore) {
					maxScore = currentScore
				}
			}
		}
	}

	fmt.Printf("Max score is: %d\n", maxScore)
}
