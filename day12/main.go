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

func (ia IntArray) Contains(n int) bool {
	for _, v := range ia {
		if v == n {
			return true
		}
	}
	return false
}

func isInProgram(p []IntArray, ia IntArray) bool {
	for _, v := range p {
		for _, v2 := range ia {
			if v.Contains(v2) {
				return true
			}
		}
	}
	return false
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

func getPath(input map[int]IntArray, start int) []int {
	var path []int
	walk := input[start]
	delete(input, start)
	path = append(path, start)
	for k, v := range input {
		input[k] = v.Remove(start)
	}
	for _, w := range walk {
		path = append(path, getPath(input, w)...)
	}
	return path
}

func main() {
	inputs := make(map[int]IntArray, 2000)
	inputsPart1 := make(map[int]IntArray, 2000)
	inputsPart2 := make(map[int]IntArray, 2000)

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

	for k, v := range inputs {
		inputsPart1[k] = v
		inputsPart2[k] = v
	}

	fmt.Printf("part 1 => %d\n", countConnects(inputsPart1, 0))

	var programs []IntArray
	for i, _ := range inputs {
		path := getPath(inputsPart2, i)
		var ia IntArray
		for _, v := range path {
			ia = ia.AddUnique(v)
		}
		if !isInProgram(programs, ia) {
			programs = append(programs, ia)
		}
	}

	fmt.Printf("part 2 => %d\n", len(programs))
}
