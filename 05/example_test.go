package day05

import (
	"fmt"
)

func ExampleStar1TopCrateList() {
	topCrateList, err := Star1TopCrateList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topCrateList)
	// Output:
	// NTWZZWHFV
}
