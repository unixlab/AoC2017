package main

import (
	"bufio"
	"fmt"
	"os"
)

func sumOfLine1(input string) int {
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

func sumOfLine2(input string) int {
	var sum int
	len := len(input)
	half := len / 2
	for i := 0; i < len; i++ {
		var j int
		j = i + half
		if j >= len {
			j = j - len
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
		fmt.Println(sumOfLine1(line))
		fmt.Println(sumOfLine2(line))
	}
}
