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
	size, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	c := map[string]int{}

	for i := 0; i < int(size); i++ {
		s.Scan()
		line := s.Text()
		if _, ok := c[line]; ok {
			c[line] = c[line] + 1
		} else {
			c[line] = 1
		}
	}

	s.Scan()
	qlen, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	for i := 0; i < int(qlen); i++ {
		s.Scan()
		query := s.Text()
		if v, ok := c[query]; ok {
			fmt.Println(v)
		} else {
			fmt.Println("0")
		}
	}
}
