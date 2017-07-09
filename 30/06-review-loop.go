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

	records, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(records); i++ {
		s.Scan()
		str := s.Text()

		for j := 0; j < len(str); j += 2 {
			fmt.Printf("%c", str[j]) // Strings behave as slices of bytes
		}
		fmt.Print(" ")
		for j := 1; j < len(str); j += 2 {
			fmt.Printf("%c", str[j])
		}
		fmt.Print("\n")
	}
}
