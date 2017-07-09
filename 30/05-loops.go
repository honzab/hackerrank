package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	i, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	for j := 1; j <= 10; j++ {
		fmt.Printf("%d x %d = %d\n", i, j, i*int64(j))
	}
}
