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
	s.Scan()
	a := strings.Split(s.Text(), " ")

	if len(a) != int(size) {
		panic("Size not as declared")
	}

	for i := size - 1; i >= 0; i-- {
		fmt.Printf("%s ", a[i])
	}
}
