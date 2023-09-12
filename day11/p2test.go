package main

import (
	"fmt"
	"math/big"
)

func main() {
	var monkeys [4][]*big.Int
	inspections := [4]int{}
	monkeys[0] = append(monkeys[0], big.NewInt(79), big.NewInt(98))
	monkeys[1] = append(monkeys[1], big.NewInt(54), big.NewInt(65), big.NewInt(75), big.NewInt(74))
	monkeys[2] = append(monkeys[2], big.NewInt(79), big.NewInt(60), big.NewInt(97))
	monkeys[3] = append(monkeys[3], big.NewInt(74))
	
	for round := 0; round < 1000; round += 1 {
		fmt.Print(round, "\n")
		for monkey := 0; monkey < 4; monkey += 1 {
			for item := 0; item < len(monkeys[monkey]); item +=1  {
				new := big.NewInt(0)
				if monkey == 0 {
					new.Mul(monkeys[monkey][item], big.NewInt(19))
				} else if monkey == 1 {
					new.Add(monkeys[monkey][item], big.NewInt(6))
				} else if monkey == 2 {
					new.Mul(monkeys[monkey][item], monkeys[monkey][item])
				} else if monkey == 3 {
					new.Add(monkeys[monkey][item], big.NewInt(3))
				}

				inspections[monkey] += 1

				mod := big.NewInt(0)
				zero := big.NewInt(0)
				if (monkey == 0) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(23))) == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				} else if (monkey == 1) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(19))) == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[0] = append(monkeys[0], new)
					}
				} else if (monkey == 2) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(13))) == 0) {
						monkeys[1] = append(monkeys[1], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				} else if (monkey == 3) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(17))) == 0) {
						monkeys[0] = append(monkeys[0], new)
					} else {
						monkeys[1] = append(monkeys[1], new)
					}
				}
			}

			monkeys[monkey] = []*big.Int{}
		}
	}
	fmt.Print(inspections)
}
