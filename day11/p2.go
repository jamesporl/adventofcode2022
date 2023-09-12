package main

import (
	"fmt"
	"math/big"
)

func main() {
	var monkeys [8][]*big.Int
	inspections := [8]int{}
	monkeys[0] = append(monkeys[0], big.NewInt(83), big.NewInt(97), big.NewInt(95), big.NewInt(67))
	monkeys[1] = append(monkeys[1], big.NewInt(71), big.NewInt(70), big.NewInt(79), big.NewInt(88), big.NewInt(56), big.NewInt(70))
	monkeys[2] = append(monkeys[2], big.NewInt(98), big.NewInt(51), big.NewInt(51), big.NewInt(63), big.NewInt(80), big.NewInt(85), big.NewInt(84), big.NewInt(95))
	monkeys[3] = append(monkeys[3], big.NewInt(77), big.NewInt(90), big.NewInt(82), big.NewInt(80), big.NewInt(79))
	monkeys[4] = append(monkeys[4], big.NewInt(68))
	monkeys[5] = append(monkeys[5], big.NewInt(60), big.NewInt(94))
	monkeys[6] = append(monkeys[6], big.NewInt(81), big.NewInt(51), big.NewInt(85))
	monkeys[7] = append(monkeys[7], big.NewInt(98), big.NewInt(81), big.NewInt(63), big.NewInt(65), big.NewInt(84), big.NewInt(71), big.NewInt(84))
	
	for round := 0; round <= 10000; round += 1 {
		fmt.Print("Round ", round, "\n")
		for monkey := 0; monkey < 8; monkey += 1 {
			for item := 0; item < len(monkeys[monkey]); item +=1  {
				new := big.NewInt(0)
				if monkey == 0 {
					new.Mul(monkeys[monkey][item], big.NewInt(19))
				} else if monkey == 1 {
					new.Add(monkeys[monkey][item], big.NewInt(2))
				} else if monkey == 2 {
					new.Add(monkeys[monkey][item], big.NewInt(7))
				} else if monkey == 3 {
					new.Add(monkeys[monkey][item], big.NewInt(1))
				} else if monkey == 4 {
					new.Mul(monkeys[monkey][item], big.NewInt(5))
				} else if monkey == 5 {
					new.Add(monkeys[monkey][item], big.NewInt(5))
				} else if monkey == 6 {
					new.Mul(monkeys[monkey][item], monkeys[monkey][item])
				} else if monkey == 7 {
					new.Add(monkeys[monkey][item], big.NewInt(3))
				}

				inspections[monkey] += 1

				mod := big.NewInt(0)
				zero := big.NewInt(0)
				if (monkey == 0) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(17))) == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[7] = append(monkeys[7], new)
					}
				} else if (monkey == 1) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(19))) == 0) {
						monkeys[7] = append(monkeys[7], new)
					} else {
						monkeys[0] = append(monkeys[0], new)
					}
				} else if (monkey == 2) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(7))) == 0) {
						monkeys[4] = append(monkeys[4], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				} else if (monkey == 3) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(11))) == 0) {
						monkeys[6] = append(monkeys[6], new)
					} else {
						monkeys[4] = append(monkeys[4], new)
					}
				} else if (monkey == 4) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(13))) == 0) {
						monkeys[6] = append(monkeys[6], new)
					} else {
						monkeys[5] = append(monkeys[5], new)
					}
				} else if (monkey == 5) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(3))) == 0) {
						monkeys[1] = append(monkeys[1], new)
					} else {
						monkeys[0] = append(monkeys[0], new)
					}
				} else if (monkey == 6) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(5))) == 0) {
						monkeys[5] = append(monkeys[5], new)
					} else {
						monkeys[1] = append(monkeys[1], new)
					}
				} else if (monkey == 7) {
					if (zero.Cmp(mod.Mod(new, big.NewInt(2))) == 0) {
						monkeys[2] = append(monkeys[2], new)
					} else {
						monkeys[3] = append(monkeys[3], new)
					}
				}
			}

			monkeys[monkey] = []*big.Int{}
		}
	}
	fmt.Print(inspections)
}
