package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	ii, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	id, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	is := scanner.Text()

	fmt.Printf("%d\n", i+ii)
	fmt.Printf("%.1f\n", d+id)
	fmt.Printf("%s%s\n", s, is)
}
