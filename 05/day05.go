package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Star1TopCrateList() (string, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer input.Close()

	var craneLines []string
	var moveLines []string
	doMoveLines := false
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			doMoveLines = true
			continue
		}
		if doMoveLines {
			moveLines = append(moveLines, line)
		} else {
			craneLines = append(craneLines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	stacks := parseCraneLinesToStacks(craneLines)
	// fmt.Println(string(stacks[3]))
	// fmt.Println(string(stacks[len(stacks)-1]))
	// fmt.Println("THEN movements")
	for _, line := range moveLines {
		quantity, from, to, err := parseMove(line)
		if err != nil {
			return "", err
		}
		stacks, err = doMove(stacks, quantity, from, to)
		if err != nil {
			return "", err
		}
		// fmt.Println(string(stacks[3]))
	}
	fmt.Println(topCrateList(stacks))
	return "", nil
}

func parseCraneLinesToStacks(lines []string) [][]rune {
	var stacks [][]rune
	stackNumberLine := []rune(lines[len(lines)-1])
	for i, maybeRawNumber := range stackNumberLine {
		_, err := runeToInt(maybeRawNumber)
		if err != nil {
			continue
		}
		stacks = append(stacks, stack(lines[:len(lines)-1], i))
	}
	return stacks
}

func stack(lines []string, index int) []rune {
	var stack []rune
	for _, line := range lines {
		runes := []rune(line)
		if len(runes) > index {
			if isValidCrane(runes[index]) {
				stack = append(stack, runes[index])
			}
		}
	}
	return stack
}

func runeToInt(r rune) (int, error) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), nil
	}
	return 0, fmt.Errorf("not valid int")
}

func isValidCrane(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func parseMove(line string) (int, int, int, error) {
	p := strings.Split(line, " ")
	if len(p) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid move input")
	}
	quantity, err := strconv.Atoi(p[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid quantity input")
	}
	from, err := strconv.Atoi(p[3])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid from input")
	}
	to, err := strconv.Atoi(p[5])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid to input")
	}
	return quantity, from - 1, to - 1, nil
}

func doMove(stacks [][]rune, quantity, from, to int) ([][]rune, error) {
	if len(stacks[from]) < quantity {
		return nil, fmt.Errorf("not enough crane")
	}
	for i := 0; i < quantity; i++ {
		c, s := stacks[from][0], stacks[from][1:]
		stacks[from] = s
		stacks[to] = append([]rune{c}, stacks[to]...)
	}
	return stacks, nil
}

func topCrateList(stacks [][]rune) string {
	var list []rune
	for _, stack := range stacks {
		if len(stack) > 0 {
			list = append(list, stack[0])
		} else {
			list = append(list, ' ')
		}
	}
	return string(list)
}
