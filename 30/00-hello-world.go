package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()

	fmt.Println("Hello, World.")
	fmt.Println(input)
}
