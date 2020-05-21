package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func checkUniq(array []string, element string) bool {
	for _, v := range array {
		if v == element {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		var uniqWords []string
		uniq := true
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			if checkUniq(uniqWords, word) {
				uniqWords = append(uniqWords, word)
			} else {
				uniq = false
			}
		}
		if uniq {
			counter++
		}
	}
	fmt.Println(counter)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	counter = 0
	for scanner.Scan() {
		var uniqWords []string
		uniq := true
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			word = sortString(word)
			if checkUniq(uniqWords, word) {
				uniqWords = append(uniqWords, word)
			} else {
				uniq = false
			}
		}
		if uniq {
			counter++
		}
	}
	fmt.Println(counter)
}
