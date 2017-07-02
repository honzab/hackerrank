package main

import (
	"bufio"
	"fmt"
	"log"
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

func main() {
	s := bufio.NewScanner(os.Stdin)
	var buf []byte
	s.Buffer(buf, 1024*1024)
	s.Scan()
	size, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	s.Scan()
	c, err := strSliceToInt64(strings.Split(s.Text(), " "))
	if err != nil {
		panic(err)
	}

	if int64(len(c)) != size {
		log.Fatalf("bad input")
	}

	counter := 0
	var max int64
	for _, v := range c {
		if v > max {
			max = v
			counter = 1
			continue
		}
		if v == max {
			counter = counter + 1
		}
	}

	fmt.Println(counter)
}
