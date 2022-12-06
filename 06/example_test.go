package day06

import (
	"fmt"
)

func ExampleNumberOfCharacterToStartOfPacket_four() {
	n, err := NumberOfCharacterToStartOfPacket(input, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	// Output:
	// 1542
}

func ExampleNumberOfCharacterToStartOfPacket_fourteen() {
	n, err := NumberOfCharacterToStartOfPacket(input, 14)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	// Output:
	// 3153
}
