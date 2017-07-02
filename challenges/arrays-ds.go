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

	values := make([]string, size)
	s.Scan()
	sp := strings.Split(s.Text(), " ")

	for i := 0; i < int(size); i++ {
		values[i] = sp[int(size)-i-1]
	}

	fmt.Println(strings.Join(values, " "))
}
