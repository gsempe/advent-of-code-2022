package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func mostCalories() {
	calories, err := computeCalories()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(maxCalories(calories))
}

func topThreeCalories() {
	calories, err := computeCalories()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(topXCalories(calories, 3))
}

func computeCalories() ([]int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	elfCalories := []int{0}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elfCalories = append(elfCalories, 0)
			continue
		}
		c, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		elfCalories[len(elfCalories)-1] += c
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return elfCalories, nil
}

func maxCalories(calories []int) int {
	return topXCalories(calories, 1)
}

func topXCalories(calories []int, x int) int {
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	v := 0
	for i := 0; i < x; i++ {
		v += calories[i]
	}
	return v
}
