package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	storage := make(map[string]int, 100)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputArray := strings.Split(scanner.Text(), " ")
		stringVal1 := inputArray[0]
		stringOperator := inputArray[1]
		stringVal2 := inputArray[2]
		stringIfVal1 := inputArray[4]
		stringIfOperator := inputArray[5]
		stringIfVal2 := inputArray[6]

		var operator bool
		if stringOperator == "inc" {
			operator = true
		} else {
			operator = false
		}

		val2, err := strconv.Atoi(stringVal2)
		if err != nil {
			var isNotEmpty bool
			val2, isNotEmpty = storage[stringVal2]
			if !isNotEmpty {
				val2 = 0
			}
		}

		ifVal1, err := strconv.Atoi(stringIfVal1)
		if err != nil {
			var isNotEmpty bool
			ifVal1, isNotEmpty = storage[stringIfVal1]
			if !isNotEmpty {
				ifVal1 = 0
			}
		}

		ifVal2, err := strconv.Atoi(stringIfVal2)
		if err != nil {
			var isNotEmpty bool
			ifVal2, isNotEmpty = storage[stringIfVal2]
			if !isNotEmpty {
				ifVal2 = 0
			}
		}

		switch stringIfOperator {
		case ">":
			if ifVal1 > ifVal2 {
				if operator {
					storage[stringVal1] += val2
				} else {
					storage[stringVal1] -= val2
				}
			}
		case ">=":
			if ifVal1 >= ifVal2 {
				if operator {
					storage[stringVal1] += val2
				} else {
					storage[stringVal1] -= val2
				}
			}
		case "<":
			if ifVal1 < ifVal2 {
				if operator {
					storage[stringVal1] += val2
				} else {
					storage[stringVal1] -= val2
				}
			}
		case "<=":
			if ifVal1 <= ifVal2 {
				if operator {
					storage[stringVal1] += val2
				} else {
					storage[stringVal1] -= val2
				}
			}
		case "==":
			if ifVal1 == ifVal2 {
				if operator {
					storage[stringVal1] += val2
				} else {
					storage[stringVal1] -= val2
				}
			}
		case "!=":
			if ifVal1 != ifVal2 {
				if operator {
					storage[stringVal1] += val2
				} else {
					storage[stringVal1] -= val2
				}
			}
		default:
			panic("error")
		}
	}
	var maxK string
	var maxV int
	for k, v := range storage {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	fmt.Printf("max is %d at %s\n", maxV, maxK)
}
