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

	for scanner.Scan() {
		txt := scanner.Text();
		txtArr := strings.Split(txt, "")
		for i := 13; i < len(txtArr); i += 1 {
			hasMatch := false
			for j := 0; j < 13; j += 1 {
				for k := 0; k < 13 - j; k += 1 {
					if (txtArr[i + j - 13] == txtArr[i + j + k - 12]) {
						hasMatch = true
						break
					}
				}
			}
			if (!hasMatch) {
				fmt.Printf("%d\n", i + 1)
				break
			}
		}
	}
}
