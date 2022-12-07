package day07

import (
	"strconv"
	"strings"
)

func isCmd(line string) bool {
	return strings.HasPrefix(line, "$ ")
}

func isCmdLs(line string) bool {
	return line == "$ ls"
}

func isCmdPopDirectory(line string) bool {
	return line == "$ cd .."
}

func isCmdPushDirectory(line string) (bool, string) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return false, ""
	}
	if parts[0] != "$" || parts[1] != "cd" {
		return false, ""
	}
	return true, parts[2]
}

func isCmdPushRootDirectory(line string) bool {
	yes, name := isCmdPushDirectory(line)
	return yes && name == "/"
}

func isResultLsDirectory(line string) (bool, string) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return false, ""
	}
	if parts[0] != "dir" {
		return false, ""
	}
	return true, parts[1]
}

func isResultLsFile(line string) (bool, string, int) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return false, "", 0
	}
	size, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, "", 0
	}
	return true, parts[1], size
}
