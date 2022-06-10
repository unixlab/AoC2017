package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntArray []int

func (ia IntArray) AddUnique(n int) IntArray {
	for _, v := range ia {
		if v == n {
			return ia
		}
	}
	ia = append(ia, n)
	return ia
}

func (ia IntArray) Remove(n int) IntArray {
	var newIA IntArray
	for _, v := range ia {
		if v != n {
			newIA = append(newIA, v)
		}
	}
	return newIA
}

func countConnects(input map[int]IntArray, start int) int {
	counter := len(input[start])
	walk := input[start]
	delete(input, start)
	for k, v := range input {
		input[k] = v.Remove(start)
	}
	for _, w := range walk {
		counter += countConnects(input, w)
	}
	return counter
}

func main() {
	inputs := make(map[int]IntArray, 2000)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		firstWhitespace := strings.Index(line, " ")
		program := line[:firstWhitespace]

		inputSeparator := strings.Index(line, ">") + 2
		connected := strings.Replace(line[inputSeparator:], " ", "", -1)

		programCode, _ := strconv.Atoi(program)

		var connectedCode []int
		for _, codeString := range strings.Split(connected, ",") {
			codeInt, _ := strconv.Atoi(codeString)
			connectedCode = append(connectedCode, codeInt)
		}

		for _, input := range connectedCode {
			inputs[programCode] = inputs[programCode].AddUnique(input)
		}
	}

	fmt.Println(countConnects(inputs, 0))
}
