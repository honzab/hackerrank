package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(i int64) int64 {
	if i <= 1 {
		return int64(1)
	} else {
		return i * factorial(i-1)
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	i, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", factorial(i))
}
