package day07

import (
	"fmt"
)

func ExampleSumOfDirectoryAbove100000() {
	sum, err := SumOfDirectoryBelow100000()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
	// Output:
	// 1743217
}

func ExampleSmallestDirectorySizeToDelete() {
	size, err := SmallestDirectorySizeToDelete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(size)
	// Output:
	// 8319096
}
