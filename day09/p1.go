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

	var tailPositions [][2]int
	tail := [2]int{0, 0}
	head := [2]int{0, 0}
	tailPositions = append(tailPositions, tail)
	for scanner.Scan() {
		txt := scanner.Text();
		moveArr := strings.Split(txt, " ")
		dir := moveArr[0]
		moveCount, _ := strconv.Atoi(moveArr[1])
		for i := 1; i <= moveCount; i += 1 {
			// prevHead := head
			if dir == "L" {
				head[0] = head[0] - 1
			} else if dir == "R" {
				head[0] = head[0] + 1
			} else if dir == "U" {
				head[1] = head[1] + 1
			} else {
				head[1] = head[1] - 1
			}
			diffX := head[0] -  tail[0]
			diffY := head[1] -  tail[1]
			
			if (diffX > 1 || diffX < -1) {
				if diffX > 1 {
					tail[0] = tail[0] + 1
				} else {
					tail[0] = tail[0] - 1
				}
				if (diffY >= 1) {
					tail[1] = tail[1] + 1
				} else if (diffY <= -1) {
					tail[1] = tail[1] - 1
				}
			} else if (diffY > 1 || diffY < -1) {
				if diffY > 1 {
					tail[1] = tail[1] + 1
				} else {
					tail[1] = tail[1] - 1
				}
				if (diffX >= 1) {
					tail[0] = tail[0] + 1
				} else if (diffX <= -1) {
					tail[0] = tail[0] - 1
				}
			}

			for i := 0; i < len(tailPositions); i += 1 {
				if tailPositions[i][0] == tail[0] && tailPositions[i][1] == tail[1] {
					break
				}
				if (i == len(tailPositions) - 1) {
					tailPositions = append(tailPositions, tail)
				}
			}
		}
	}
	fmt.Printf("Tail unique positions %d\n", len(tailPositions))
}
