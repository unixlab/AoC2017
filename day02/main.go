package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile("([0-9]+)")

	var checkSum int

	for scanner.Scan() {

		min := -1
		max := 0
		line := scanner.Text()
		regexRes := regex.FindAllString(line, -1)
		for i := 0; i < len(regexRes); i++ {
			number, _ := strconv.Atoi(regexRes[i])
			if number < min || min == -1 {
				min = number
			}
			if number > max {
				max = number
			}
		}
		checkSum += max - min
	}
	fmt.Println(checkSum)
}
