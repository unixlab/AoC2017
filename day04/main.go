package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
}
