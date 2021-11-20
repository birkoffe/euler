package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readMatrix(filename string) ([][]int, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	var matrix [][]int

	i := 0
	for _, line := range lines {
		if line == "" {
			break
		}

		matrix = append(matrix, nil)

		for _, s := range strings.Split(line, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			matrix[i] = append(matrix[i], n)
		}

		i++
	}

	return matrix, nil
}

func get(i, j int, grid [][]int, move []int) int {
	prod := 1

	for x := 0; x < 4; x++ {
		ii := i + move[x]
		jj := j + move[x+4]
		if ii < 0 || jj < 0 || ii > 19 || jj > 19 {
			return 0
		}
		prod *= grid[ii][jj]
	}

	return prod
}

func main() {
	matrix, err := readMatrix("011/input.txt")
	if err != nil {
		panic(err)
	}

	moves := [][]int{{0, 0, 0, 0, 0, 1, 2, 3}, {0, 1, 2, 3, 0, 0, 0, 0}, {0, 1, 2, 3, 0, 1, 2, 3}, {0, -1, -2, -3, 0, 1, 2, 3}}

	result := 0

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			for k := 0; k < len(moves); k++ {
				prod := get(i, j, matrix, moves[k])
				if prod > result {
					result = prod
				}
			}
		}
	}

	fmt.Println(int(result))
}
