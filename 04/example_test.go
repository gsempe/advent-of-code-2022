package day04

import (
	"fmt"
)

func ExampleStar1FullyContain() {
	c, err := Star1FullyContain()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	// output:
	// 496
}

func ExampleStar2Overlapping() {
	c, err := Star2Overlapping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	// output:
	// 496
}
