package day07

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func SumOfDirectoryBelow100000() (int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	sum := 0
	current := root
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return 0, fmt.Errorf("empty line")
		}
		if isCmd(line) {
			if isCmdPopDirectory(line) {
				var size int
				current, size = current.popDirectory()
				if shouldAddSize(size) {
					sum += size
				}
				continue
			}
			if isCmdLs(line) {
				continue
			}
			if isCmdPushRootDirectory(line) {
				current = root
				continue
			}
			if yes, name := isCmdPushDirectory(line); yes {
				current = current.pushDirectory(name)
				continue
			}
			return 0, fmt.Errorf("unknown command: %s", line)
		} else {
			if yes, _ := isResultLsDirectory(line); yes {
				continue
			}
			if yes, name, size := isResultLsFile(line); yes {
				current.addFile(name, size)
				continue
			}
			return 0, fmt.Errorf("unknown result: %s", line)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	// Popping pushed directories
	for !current.isRoot() {
		var size int
		current, size = current.popDirectory()
		if shouldAddSize(size) {
			sum += size
		}
	}
	return sum, nil
}

func SmallestDirectorySizeToDelete() (int, error) {
	input, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var directorySizes []int
	current := root
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return 0, fmt.Errorf("empty line")
		}
		if isCmd(line) {
			if isCmdPopDirectory(line) {
				var size int
				current, size = current.popDirectory()
				directorySizes = append(directorySizes, size)
				continue
			}
			if isCmdLs(line) {
				continue
			}
			if isCmdPushRootDirectory(line) {
				current = root
				continue
			}
			if yes, name := isCmdPushDirectory(line); yes {
				current = current.pushDirectory(name)
				continue
			}
			return 0, fmt.Errorf("unknown command: %s", line)
		} else {
			if yes, _ := isResultLsDirectory(line); yes {
				continue
			}
			if yes, name, size := isResultLsFile(line); yes {
				current.addFile(name, size)
				continue
			}
			return 0, fmt.Errorf("unknown result: %s", line)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	// Popping pushed directories
	for !current.isRoot() {
		var size int
		current, size = current.popDirectory()
		directorySizes = append(directorySizes, size)
	}
	totalUsed := current.size
	freeSpace := 70000000 - totalUsed
	missingSpace := 30000000 - freeSpace
	if missingSpace < 0 {
		return 0, fmt.Errorf("already enough space")
	}
	sort.Ints(directorySizes)
	for _, size := range directorySizes {
		if size >= missingSpace {
			return size, nil
		}
	}
	return 0, fmt.Errorf("no directory big enough")
}

func shouldAddSize(size int) bool {
	return size <= 100000
}
