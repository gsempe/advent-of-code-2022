package day02

import (
	"fmt"
)

func ExampleTotalScore() {
	score, err := TotalScore()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(score)
	// output:
	// 10310
}
