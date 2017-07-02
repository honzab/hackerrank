package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func strSliceToInt64(s []string) ([]int64, error) {
	is := make([]int64, len(s))
	for k, v := range s {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		is[k] = i
	}

	return is, nil
}

func abs(i int64) int64 {
	return int64(math.Abs(float64(i)))

}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var buf []byte
	s.Buffer(buf, 1024*1024)
	s.Scan()
	size, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	var i int64
	for i = 0; i < size; i++ {
		s.Scan()
		c, err := strSliceToInt64(strings.Split(s.Text(), " "))
		if err != nil {
			panic(err)
		}

		catB := abs(c[2] - c[1])
		catA := abs(c[2] - c[0])

		if catB == catA {
			fmt.Println("Mouse C")
		} else if catA > catB {
			fmt.Println("Cat B")
		} else {
			fmt.Println("Cat A")
		}
	}
}
