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
	s.Scan()

	size, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	pb := make(map[string]string, size)
	for i := 0; i < int(size); i++ {
		s.Scan()
		line := strings.Split(s.Text(), " ")
		if len(line) != 2 {
			panic("Malformed input line")
		}
		pb[line[0]] = line[1]
	}

	for s.Scan() {
		lookup := s.Text()
		if _, ok := pb[lookup]; ok {
			fmt.Printf("%s=%s\n", lookup, pb[lookup])
		} else {
			fmt.Println("Not found")
		}
	}
}
