package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func round(f float64) int {
	if f < -0.5 {
		return int(f - 0.5)
	}
	if f > 0.5 {
		return int(f + 0.5)
	}
	return 0
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	mc, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}

	s.Scan()
	ti, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}

	s.Scan()
	ta, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}

	tip := mc * (ti / 100.0)
	tax := mc * (ta / 100.0)

	fmt.Printf("The total meal cost is %d dollars.", round(mc+tip+tax))
}
