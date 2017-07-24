package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func bubbleSort(a []int64) ([]int64, int) {
	n := len(a)
	swaps := 0
	for i := 0; i < n; i++ {
		// Track number of elements swapped during a single array traversal
		numberOfSwaps := 0

		for j := 0; j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				numberOfSwaps += 1
				swaps += 1
			}
		}

		// If no elements were swapped during a traversal, array is sorted
		if numberOfSwaps == 0 {
			break
		}
	}
	return a, swaps
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan() // Ignore first line

	s.Scan()
	els, err := strSliceToInt64(strings.Split(s.Text(), " "))
	if err != nil {
		panic(err)
	}

	a, swaps := bubbleSort(els)

	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[len(a)-1])
}
