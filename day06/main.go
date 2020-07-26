package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intArrayToString(array []int) string {
	removeArray := strings.NewReplacer("[", "", "]", "", " ", "")
	return removeArray.Replace(fmt.Sprint(array))
}

func main() {
	var banks []int
	seen := make(map[string]int, 100)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stringBanks := strings.Split(scanner.Text(), "\t")
		for _, v := range stringBanks {
			number, _ := strconv.Atoi(v)
			banks = append(banks, number)
		}
	}

	runs := 0
	for {
		runs++

		maxK := -1
		maxV := -1
		for k, v := range banks {
			if v > maxV {
				maxK = k
				maxV = v
			}
		}

		banks[maxK] = 0
		pointer := maxK
		for i := maxV; i > 0; i-- {
			pointer++
			if pointer == len(banks) {
				pointer = 0
			}
			banks[pointer]++
		}

		if loopSize, exists := seen[intArrayToString(banks)]; exists {
			fmt.Println(runs)
			fmt.Println(loopSize)
			break
		}

		seen[intArrayToString(banks)] = 0
		for k, v := range seen {
			seen[k] = v + 1
		}
	}
}
