package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeExclamationMark(s string) string {
	for strings.Index(s, "!") >= 0 {
		pos := strings.Index(s, "!")
		s = s[:pos] + s[pos+2:]
	}
	return s
}

func removeGarbage(s string) string {
	for strings.Index(s, "<") >= 0 {
		sPos := strings.Index(s, "<")
		ePos := strings.Index(s, ">")
		s = s[:sPos] + s[ePos+1:]
	}
	return s
}

func removeComma(s string) string {
	return strings.Replace(s, ",", "", -1)
}

func countGroups(s string) int {
	sum := 0
	depth := 0
	for _, char := range strings.Split(s, "") {
		if char == "{" {
			depth++
		} else {
			sum += depth
			depth--
		}
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		withoutExclamationMark := removeExclamationMark(scanner.Text())
		withoutGarbage := removeGarbage(withoutExclamationMark)
		withoutComma := removeComma(withoutGarbage)
		sum := countGroups(withoutComma)
		fmt.Printf("part 1 => %d\n", sum)
	}
}
