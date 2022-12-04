package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Star1FullyContain() (int, error) {
	return commonApproach(func(p Pair) bool {
		return p.OneContainsTheOther()
	})
}

func Star2Overlapping() (int, error) {
	return commonApproach(func(p Pair) bool {
		return p.areOverlapping()
	})
}

func commonApproach(f func(Pair) bool) (int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return 0, fmt.Errorf("empty line")
		}
		p, err := textToPair(line)
		if err != nil {
			return 0, err
		}
		if f(p) {
			count += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil

}

type Pair struct {
	FirstRange  Range
	SecondRange Range
}

type Range struct {
	Min int
	Max int
}

func (p Pair) OneContainsTheOther() bool {
	firstWay := p.FirstRange.Min <= p.SecondRange.Min && p.FirstRange.Max >= p.SecondRange.Max
	secondWay := p.SecondRange.Min <= p.FirstRange.Min && p.SecondRange.Max >= p.FirstRange.Max
	return firstWay || secondWay
}

func (p Pair) areOverlapping() bool {
	if p.OneContainsTheOther() {
		return true
	}
	firstWay := p.FirstRange.Min <= p.SecondRange.Min && p.FirstRange.Max >= p.SecondRange.Min
	secondWay := p.SecondRange.Min <= p.FirstRange.Min && p.SecondRange.Max >= p.FirstRange.Min
	return firstWay || secondWay
}

func textToPair(text string) (Pair, error) {
	raw := strings.Split(text, ",")
	if len(raw) != 2 {
		return Pair{}, fmt.Errorf("invalid number of ranges")
	}
	firstRange, err := textToRange(raw[0])
	if err != nil {
		return Pair{}, err
	}
	secondRange, err := textToRange(raw[1])
	if err != nil {
		return Pair{}, err
	}
	return Pair{FirstRange: firstRange, SecondRange: secondRange}, nil
}

func textToRange(text string) (Range, error) {
	raw := strings.Split(text, "-")
	if len(raw) != 2 {
		return Range{}, fmt.Errorf("invalid number of limits")
	}
	min, err := textToLimit(raw[0])
	if err != nil {
		return Range{}, err
	}
	max, err := textToLimit(raw[1])
	if err != nil {
		return Range{}, err
	}
	return Range{Min: min, Max: max}, nil
}

func textToLimit(text string) (int, error) {
	return strconv.Atoi(text)
}
