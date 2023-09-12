package main

import (
	"fmt"
)

func main() {
	var monkeys [8][]int
	inspections := [8]int{}
	monkeys[0] = append(monkeys[0], 83, 97, 95, 67)
	monkeys[1] = append(monkeys[1], 71, 70, 79, 88, 56, 70)
	monkeys[2] = append(monkeys[2], 98, 51, 51, 63, 80, 85, 84, 95)
	monkeys[3] = append(monkeys[3], 77, 90, 82, 80, 79)
	monkeys[4] = append(monkeys[4], 68)
	monkeys[5] = append(monkeys[5], 60, 94 )
	monkeys[6] = append(monkeys[6], 81, 51, 85)
	monkeys[7] = append(monkeys[7], 98, 81, 63, 65, 84, 71, 84)
	
	for round := 0; round < 20; round += 1 {
		for monkey := 0; monkey < 8; monkey += 1 {
			for item := 0; item < len(monkeys[monkey]); item +=1  {
				new := 0
				if monkey == 0 {
					new = monkeys[monkey][item] * 19
				} else if monkey == 1 {
					new = monkeys[monkey][item] + 2
				} else if monkey == 2 {
					new = monkeys[monkey][item] + 7
				} else if monkey == 3 {
					new = monkeys[monkey][item] + 1
				} else if monkey == 4 {
					new = monkeys[monkey][item] * 5
				} else if monkey == 5 {
					new = monkeys[monkey][item] + 5
				} else if monkey == 6 {
					new = monkeys[monkey][item] * monkeys[monkey][item]
				} else if monkey == 7 {
					new = monkeys[monkey][item] + 3
				}

				inspections[monkey] += 1

				new = new / 3

				if (monkey == 0) {
					if (new % 17 == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[7] = append(monkeys[7], new)
					}
				} else if (monkey == 1) {
					if (new % 19 == 0) {
						monkeys[7] = append(monkeys[7], new)
					} else {
						monkeys[0] = append(monkeys[0], new)
					}
				} else if (monkey == 2) {
					if (new % 7 == 0) {
						monkeys[4] = append(monkeys[4], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				} else if (monkey == 3) {
					if (new % 11 == 0) {
						monkeys[6] = append(monkeys[6], new)
					} else {
						monkeys[4] = append(monkeys[4], new)
					}
				} else if (monkey == 4) {
					if (new % 13 == 0) {
						monkeys[6] = append(monkeys[6], new)
					} else {
						monkeys[5] = append(monkeys[5], new)
					}
				} else if (monkey == 5) {
					if (new % 3 == 0) {
						monkeys[1] = append(monkeys[1], new)
					} else {
						monkeys[0] = append(monkeys[0], new)
					}
				} else if (monkey == 6) {
					if (new % 5 == 0) {
						monkeys[5] = append(monkeys[5], new)
					} else {
						monkeys[1] = append(monkeys[1], new)
					}
				} else if (monkey == 7) {
					if (new % 2 == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				}
			}

			monkeys[monkey] = []int{}
		}
	}
	fmt.Print(inspections)
}
