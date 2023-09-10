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

	prioSum := 0
	for scanner.Scan() {
		txt := scanner.Text();
		length := len(txt);
		comp1 := txt[:length/2]
		comp2 := txt[length/2:]
		letters := strings.Split(comp1, "")
		common := "a"
		for i := 0; i < length/2; i++ {
			if strings.Contains(comp2, letters[i]) {
				common = letters[i]
				break
			}
		}
		prioSum += strings.Index(prios, common) + 1
	}

	fmt.Printf("Sum of priorities is %d\n", prioSum)
}
