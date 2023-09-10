package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const prios = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

	group := [3]string{}
	lineCounter := 0
	prioSum := 0
	for scanner.Scan() {
		txt := scanner.Text();
		group[lineCounter] = txt
		lineCounter += 1
		if lineCounter == 3 {
			lineCounter = 0
			letters := strings.Split(group[0], "")
			length := len(group[0])
			common := "a"
			for i := 0; i < length; i++ {
				if strings.Contains(group[1], letters[i]) &&  strings.Contains(group[2], letters[i]) {
					common = letters[i]
					break
				}
			}
			prioSum += strings.Index(prios, common) + 1
		}
	}
	fmt.Printf("Sum of priorities is %d\n", prioSum)
}
