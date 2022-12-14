package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TotalScore() (int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return 0, fmt.Errorf("empty line")
		}
		plays := strings.Split(line, " ")
		if len(plays) != 2 {
			return 0, fmt.Errorf("number of play != 2")
		}
		score += roundScore(plays[1], plays[0])
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return score, nil
}

func TotalScoreSecondCode() (int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return 0, fmt.Errorf("empty line")
		}
		plays := strings.Split(line, " ")
		if len(plays) != 2 {
			return 0, fmt.Errorf("number of play != 2")
		}
		score += roundSecondStrategyScore(plays[1], plays[0])
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return score, nil
}

func roundScore(me, them string) int {
	switch me {
	case "X":
		return xScore(them)
	case "Y":
		return yScore(them)
	case "Z":
		return zScore(them)
	default:
		panic("invalid me input")
	}
}

// xScore computes the score when I play rock
func xScore(them string) int {
	letterScore := 1
	switch them {
	case "A":
		return 3 + letterScore
	case "B":
		return 0 + letterScore
	case "C":
		return 6 + letterScore
	default:
		panic("invalid input")
	}
}

// yScore computes the score when I play paper
func yScore(them string) int {
	letterScore := 2
	switch them {
	case "A":
		return 6 + letterScore
	case "B":
		return 3 + letterScore
	case "C":
		return 0 + letterScore
	default:
		panic("invalid input")
	}
}

// zScore computes the score when I play scissors
func zScore(them string) int {
	letterScore := 3
	switch them {
	case "A":
		return 0 + letterScore
	case "B":
		return 6 + letterScore
	case "C":
		return 3 + letterScore
	default:
		panic("invalid input")
	}
}

func roundSecondStrategyScore(me, them string) int {
	switch me {
	case "X":
		return xLosingScore(them)
	case "Y":
		return yDrawScore(them)
	case "Z":
		return zWinningScore(them)
	default:
		panic("invalid me input")
	}
}

// xLosingScore computes the score when I need to loose
func xLosingScore(them string) int {
	switch them {
	case "A":
		return 3
	case "B":
		return 1
	case "C":
		return 2
	default:
		panic("invalid input")
	}
}

// yDrawScore computes the score when I need to draw
func yDrawScore(them string) int {
	resultScore := 3
	switch them {
	case "A":
		return 1 + resultScore
	case "B":
		return 2 + resultScore
	case "C":
		return 3 + resultScore
	default:
		panic("invalid input")
	}
}

// zWinningScore computes the score when I play scissors
func zWinningScore(them string) int {
	resultScore := 6
	switch them {
	case "A":
		return 2 + resultScore
	case "B":
		return 3 + resultScore
	case "C":
		return 1 + resultScore
	default:
		panic("invalid input")
	}
}
