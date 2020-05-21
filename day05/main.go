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
	// save memory
	savedMemory := make([]int, len(memory))
	copy(savedMemory, memory)

	offset := 0
	counter := 0
	for offset < len(memory) {
		jump := memory[offset]
		memory[offset]++
		offset += jump
		counter++
	}
	fmt.Println(counter)

	// reset memory, counter and offset
	offset = 0
	counter = 0
	copy(memory, savedMemory)
	for offset < len(memory) {
		jump := memory[offset]
		if jump >= 3 {
			memory[offset]--
		} else {
			memory[offset]++
		}
		offset += jump
		counter++
	}
	fmt.Println(counter)
}
