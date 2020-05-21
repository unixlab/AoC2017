package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var memory []int

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data, _ := strconv.Atoi(scanner.Text())
		memory = append(memory, data)
	}

	offset := 0
	counter := 0
	for offset < len(memory) {
		jump := memory[offset]
		memory[offset]++
		offset += jump
		counter++
	}

	fmt.Println(counter)
}
