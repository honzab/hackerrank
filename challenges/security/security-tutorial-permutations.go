package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func f(iv map[int64]int64, v int64) int64 {
	return iv[v]
}

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
		iv[int64(k+1)] = ival
	}

	for i := 1; i <= len(els); i++ {
		fmt.Println(f(iv, f(iv, int64(i))))
	}
}
