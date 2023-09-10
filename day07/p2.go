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

	var currentPathSlice []string
	dirSizes := make(map[string]int)
	currentPathStr := ""
	for scanner.Scan() {
		txt := scanner.Text();
		commandArr := strings.Split(txt, " ")
		if commandArr[0] == "$" {
			if commandArr[1] == "cd" {
				if (commandArr[2] == "..") {
					currentPathSlice = currentPathSlice[:len(currentPathSlice) - 1]
				} else {
					leaf := commandArr[2]
					if (leaf == "/") {
						leaf = "."
					}
					currentPathSlice = append(currentPathSlice, leaf)
					currentPathStr = strings.Join(currentPathSlice, "/")
					if (dirSizes[currentPathStr] == 0) {
						dirSizes[currentPathStr] = 0
					}
				}
			}
		} else if commandArr[0] != "dir" {
			fileSize, _ := strconv.Atoi(commandArr[0])
			currentPathSlice := strings.Split(currentPathStr, "/")
			for i := 0; i < len(currentPathSlice); i += 1 {
				pathToAdd := strings.Join(currentPathSlice[:len(currentPathSlice) - i], "/")
				dirSizes[pathToAdd] += fileSize
			}
		}
	}
	
	toRemoveSize := 30000000 - (70000000 - dirSizes["."])
	minDirSize := dirSizes["."]
	for _, dirSize := range dirSizes {
		if dirSize <= minDirSize && dirSize >= toRemoveSize {
			minDirSize = dirSize
		}
	}

	fmt.Println("Size: %d", minDirSize)
}
