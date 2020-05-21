package main

import (
	"fmt"
	"math"
)

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func main() {
	number := 277678
	i := 0
	layer := 0
	numberOfSquares := 1

	for numberOfSquares < number {
		layer++
		i = i + 2
		numberOfSquares += 4 * i
	}

	side := (numberOfSquares - number) / i
	pos := abs(i/2 - (numberOfSquares-number)%i)

	var x, y int

	switch side {
	case 0:
		x = pos
		y = layer
	case 1:
		x = layer
		y = pos
	case 2:
		x = pos
		y = layer
	case 3:
		x = layer
		y = pos
	}

	fmt.Println(getDistance(0, 0, x, y))
}
