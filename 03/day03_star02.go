package day03

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
)

func SumOfGroupPriorities() (int, error) {

	groups, err := elfGroups()
	if err != nil {
		return 0, err
	}
	total := 0
	for _, group := range groups {
		commonItem, err := group.commonItem()
		if err != nil {
			return 0, err
		}
		priority, err := itemTypeToPriority(commonItem)
		if err != nil {
			return 0, err
		}
		total += priority
	}
	return total, nil
}

type RucksackGroup struct {
	FirstRucksack  string
	SecondRucksack string
	ThirdRucksack  string
}

func (g RucksackGroup) IsEmpty() bool {
	return len(g.FirstRucksack) == 0 && len(g.SecondRucksack) == 0 && len(g.ThirdRucksack) == 0
}

func (g RucksackGroup) IsValid() bool {
	return len(g.FirstRucksack) > 0 && len(g.SecondRucksack) > 0 && len(g.ThirdRucksack) > 0
}

func (g RucksackGroup) commonItem() (string, error) {
	f := func(r, s []rune) []rune {
		var output []rune
		for _, rc := range r {
			for _, sc := range s {
				if rc == sc && !slices.Contains(output, rc) {
					output = append(output, rc)
				}
			}
		}
		return output
	}
	intermediary := f([]rune(g.FirstRucksack), []rune(g.SecondRucksack))
	output := f(intermediary, []rune(g.ThirdRucksack))

	if len(output) != 1 {
		return "", fmt.Errorf("invalid number of common items")
	}
	return string(output), nil
}

func elfGroups() ([]RucksackGroup, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var groups []RucksackGroup
	for {
		g := scanGroup(scanner)
		if g.IsEmpty() {
			break
		}
		if !g.IsValid() {
			return nil, fmt.Errorf("invalid rucksack group")
		}
		groups = append(groups, g)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func scanGroup(scanner *bufio.Scanner) RucksackGroup {
	g := RucksackGroup{}
	if scanner.Scan() {
		g.FirstRucksack = scanner.Text()
	}
	if scanner.Scan() {
		g.SecondRucksack = scanner.Text()
	}
	if scanner.Scan() {
		g.ThirdRucksack = scanner.Text()
	}
	return g
}
