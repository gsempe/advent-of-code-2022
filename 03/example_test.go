package day03

import (
	"fmt"
)

func ExampleSumOfPriorities() {
	total, err := SumOfPriorities()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)
	// output:
	// 8072
}
