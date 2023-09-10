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
		for i := 3; i < len(txtArr); i += 1 {
			hasMatch := false
			for j := 0; j < 3; j += 1 {
				for k := 0; k < 3 - j; k += 1 {
					if (txtArr[i + j - 3] == txtArr[i + j + k - 2]) {
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
