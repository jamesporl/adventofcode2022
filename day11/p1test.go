package main

import (
	"fmt"
)

func main() {
	var monkeys [4][]int
	inspections := [4]int{}
	monkeys[0] = append(monkeys[0], 79, 98)
	monkeys[1] = append(monkeys[1], 54, 65, 75, 74)
	monkeys[2] = append(monkeys[2], 79, 60, 97)
	monkeys[3] = append(monkeys[3], 74)
	
	for round := 0; round < 20; round += 1 {
		for monkey := 0; monkey < 4; monkey += 1 {
			for item := 0; item < len(monkeys[monkey]); item +=1  {
				new := 0
				if monkey == 0 {
					new = monkeys[monkey][item] * 19
				} else if monkey == 1 {
					new = monkeys[monkey][item] + 6
				} else if monkey == 2 {
					new = monkeys[monkey][item] * monkeys[monkey][item]
				} else if monkey == 3 {
					new = monkeys[monkey][item] + 3
				}

				inspections[monkey] += 1

				new = new / 3

				if (monkey == 0) {
					if (new % 23 == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				} else if (monkey == 1) {
					if (new % 19 == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[0] = append(monkeys[0], new)
					}
				} else if (monkey == 2) {
					if (new % 13 == 0) {
						monkeys[1] = append(monkeys[1], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				} else if (monkey == 3) {
					if (new % 17== 0) {
						monkeys[0] = append(monkeys[0], new)
					} else {
						monkeys[1] = append(monkeys[1], new)
					}
				}
			}

			monkeys[monkey] = []int{}
		}
	}
	fmt.Print(inspections)
}
