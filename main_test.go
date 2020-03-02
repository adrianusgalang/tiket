package main

import (
	"testing"

	"github.com/testify/assert"
)

func TestRandom(t *testing.T) {
	result := random(0, 10)
	assert.NotNil(t, result, "result should not null")
}

func TestGenerateBoardArray(t *testing.T) {
	x := 3
	y := 2
	testBoard := generateBoardArray(x, y, 0)
	for a := 0; a < x; a++ {
		for b := 0; b < y; b++ {
			assert.Equal(t, testBoard[a][b], 0, "value should be 0")
		}
	}
}

func TestSolution(t *testing.T) {
	x := 3
	y := 2
	h := 2
	sum := 0
	result := Solution(x, y, h)
	for a := 0; a < x; a++ {
		for b := 0; b < y; b++ {
			sum += result[a][b]
		}
	}
	assert.Equal(t, sum, h, "value should be same")
}

func TestSolutionOutRange(t *testing.T) {
	x := 3
	y := 2
	h := 10
	sum := 0
	result := Solution(x, y, h)
	for a := 0; a < x; a++ {
		for b := 0; b < y; b++ {
			sum += result[a][b]
		}
	}
	assert.Equal(t, sum, x*y, "value should be same - out range")
}

func TestSolutionMinus(t *testing.T) {
	x := 3
	y := 2
	h := -10
	sum := 0
	result := Solution(x, y, h)
	for a := 0; a < x; a++ {
		for b := 0; b < y; b++ {
			sum += result[a][b]
		}
	}
	assert.Equal(t, sum, 0, "value should zero")
}
