package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MIN int64 = -9 * 7 // Minimal value is -9, hourglass sum is then times 7

func strSliceToInt64(s []string) ([]int64, error) {
	is := make([]int64, len(s))
	for k, v := range s {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		is[k] = i
	}
	return is, nil
}

func loadFromInput(s *bufio.Scanner) [][]int64 {
	s.Scan()
	i := strings.Split(s.Text(), " ")
	size := len(i)
	matrix := make([][]int64, size)
	for i := 0; i < size; i++ {
		l, err := strSliceToInt64(strings.Split(s.Text(), " "))
		if err != nil {
			panic(err)
		}
		matrix[i] = l
		s.Scan()
	}
	return matrix
}

// Return zero if there is no hourglass on those coordinates
func findHourglassSum(i, j int, m [][]int64) int64 {
	if i > len(m)-3 {
		// Too far to the right
		return MIN
	}
	if j > len(m)-3 {
		// Too far down
		return MIN
	}
	return m[i][j] + m[i][j+1] + m[i][j+2] + m[i+1][j+1] + m[i+2][j] + m[i+2][j+1] + m[i+2][j+2]
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	matrix := loadFromInput(s)
	size := len(matrix)
	max := int64(MIN)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum := findHourglassSum(i, j, matrix)
			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
