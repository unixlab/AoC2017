package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"regexp"
	"strconv"
)

func Combinations(set []string, n int) (subsets [][]string) {
	length := uint(len(set))
	if n > len(set) {
		n = len(set)
	}
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}
		var subset []string
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

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

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	checkSum = 0
	for scanner.Scan() {
		line := scanner.Text()
		regexRes := regex.FindAllString(line, -1)
		for _, v := range Combinations(regexRes, 2) {
			number1, _ := strconv.Atoi(v[0])
			number2, _ := strconv.Atoi(v[1])
			if number1%number2 == 0 {
				checkSum += number1 / number2
			}
			if number2%number1 == 0 {
				checkSum += number2 / number1
			}
		}
	}
	fmt.Println(checkSum)
}
