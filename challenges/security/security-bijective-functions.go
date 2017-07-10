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

	bijective := true
	used := []int64{}
	for j := 0; j < len(a); j++ {
		k, err := strconv.ParseInt(a[j], 10, 64)
		if err != nil {
			panic(err)
		}
		if k > size {
			bijective = false
			break
		}
		for l := 0; l < len(used); l++ {
			if used[l] == k {
				bijective = false
				break
			}
		}
		used = append(used, k)
	}

	if bijective {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
