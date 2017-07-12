package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxConsecutiveOnes(i int64) int64 {
	s := fmt.Sprintf("%b", i)
	var consecutive, max int64
	for _, v := range s {
		if string(v) == "1" {
			consecutive += 1
			if consecutive > max {
				max = consecutive
			}
		} else {
			consecutive = 0
		}
	}
	return max
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	i, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", maxConsecutiveOnes(i))
}
