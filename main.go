package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boardArray := Solution(3, 3, -1)
	fmt.Println(boardArray)
}

func Solution(x int, y int, h int) [][]int {
	if h < x*y {
		boardArray := generateBoardArray(x, y, 0)
		for a := 0; a < h; a++ {
			for {
				X := random(0, x)
				Y := random(0, y)
				if boardArray[X][Y] != 1 {
					boardArray[X][Y] = 1
					break
				}
			}
		}
		return boardArray
	} else {
		boardArray := generateBoardArray(x, y, 1)
		return boardArray
	}
}

func generateBoardArray(x, y, value int) [][]int {
	var result [][]int

	for a := 0; a < x; a++ {
		var res []int
		for b := 0; b < y; b++ {
			res = append(res, value)
		}
		result = append(result, res)
	}
	return result
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
