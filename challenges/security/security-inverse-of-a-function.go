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
	s.Scan() // Ignore the first number
	s.Scan()
	els := strings.Split(s.Text(), " ")
	iv := make(map[int64]int64, len(els))

	for k, v := range els {
		ival, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		iv[ival] = int64(k + 1)
	}

	for i := 1; i <= len(els); i++ {
		fmt.Println(iv[int64(i)])
	}
}
