package main

import (
	"fmt"
	"math/big"
)

func main() {
	// new := big.NewInt(0)
	test := big.NewInt(1)
	// new.Mod(big.NewInt(25),  big.NewInt(4))
	// new.Mod(big.NewInt(25),  big.NewInt(4))
	fmt.Print(test.Cmp(big.NewInt(1)) == 0)
	// fmt.Print(new.Mod(big.NewInt(25),  big.NewInt(4)) == big.NewInt(1))
}
