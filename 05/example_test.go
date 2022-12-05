package day05

import (
	"fmt"
)

func ExampleTopCrateListWithCrateMover_star1WithCrateMover9000() {
	topCrateList, err := TopCrateListWithCrateMover(doMoveWithCrateMover9000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topCrateList)
	// Output:
	// NTWZZWHFV
}

func ExampleTopCrateListWithCrateMover_star1WithCrateMover9001() {
	topCrateList, err := TopCrateListWithCrateMover(doMoveWithCrateMover9001)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topCrateList)
	// Output:
	// BRZGFVBTJ
}
