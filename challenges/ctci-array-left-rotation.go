package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var buf []byte
	s.Buffer(buf, 1024*1024) // Should be enough for 10^5 integers
	s.Scan()

	setup := strings.Split(s.Text(), " ")

	size, _ := strconv.ParseInt(setup[0], 10, 64)
	shift, _ := strconv.ParseInt(setup[1], 10, 64)

	array := make([]int64, size)
	s.Scan()
	if err := s.Err(); err != nil {
		panic(err)
	}
	strs := strings.Split(s.Text(), " ")
	for i := 0; i < int(size); i++ {
		value, err := strconv.ParseInt(strs[i], 10, 64)
		if err != nil {
			panic(err)
		}
		array[i] = value
	}

	for i := 0; i < int(shift); i++ {
		x := array[0]
		array = array[1:]
		array = append(array, x)
	}

	for i := 0; i < int(size); i++ {
		fmt.Printf("%d ", array[i])
	}
}
