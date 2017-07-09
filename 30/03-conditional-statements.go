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
	if i%2 != 0 {
		fmt.Println("Weird")
	} else if 2 <= i && i <= 5 {
		fmt.Println("Not Weird")
	} else if 6 <= i && i <= 20 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
