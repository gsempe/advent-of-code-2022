package day03

import (
	"bufio"
	"fmt"
	"os"
)

func SumOfPriorities() (int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		if len(rucksack) == 0 {
			return 0, fmt.Errorf("empty line")
		}
		firstCompartment, secondCompartment := rucksackToCompartment(rucksack)
		common, err := commonItem(firstCompartment, secondCompartment)
		if err != nil {
			return 0, err
		}
		priority, err := itemTypeToPriority(common)
		if err != nil {
			return 0, err
		}
		total += priority
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return total, nil
}

func rucksackToCompartment(rucksack string) (string, string) {
	r := []rune(rucksack)
	m := len(r) / 2
	return string(r[0:m]), string(r[m:])
}

func commonItem(first, second string) (string, error) {
	firstRunes := []rune(first)
	secondRunes := []rune(second)
	for _, firstRune := range firstRunes {
		for _, secondRune := range secondRunes {
			if firstRune == secondRune {
				return string(firstRune), nil
			}
		}
	}
	return "", fmt.Errorf("no common item")
}

func itemTypeToPriority(i string) (int, error) {
	c := []rune(i)[0]
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1, nil
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27, nil
	}
	return 0, fmt.Errorf("invalid item %s", string(c))
}
