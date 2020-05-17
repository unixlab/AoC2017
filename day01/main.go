package main

import (
	"bufio"
	"fmt"
	"os"
)

func sumOfLine(input string) int {
	var sum int
	len := len(input)
	for i := 0; i < len; i++ {
		var j int
		if i == len-1 {
			j = 0
		} else {
			j = i + 1
		}
		if input[i] == input[j] {
			sum += int(input[i]) - 48
		}
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(sumOfLine(line))
	}
}
