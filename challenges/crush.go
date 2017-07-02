package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min   int64
	max   int64
	value int64
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	nms := strings.Split(s.Text(), " ")

	if len(nms) != 2 {
		log.Fatalf("Bad format of first line")
	}

	arraysize, _ := strconv.ParseInt(nms[0], 10, 64)
	operations, _ := strconv.ParseInt(nms[1], 10, 64)

	if arraysize < 3 {
		log.Fatalf("Cannot be smaller than 3")
	}

	if arraysize > int64(math.Pow10(7)) {
		log.Fatalf("Cannot be bigger than 10e7")
	}

	a := make(map[int64]int64, arraysize)

	for i := 0; i < int(operations); i++ {
		s.Scan()
		cmd := strings.Split(s.Text(), " ")
		if len(cmd) != 3 {
			log.Fatalf("Malformed cmd line")
		}
		from, _ := strconv.ParseInt(cmd[0], 10, 64)
		to, _ := strconv.ParseInt(cmd[1], 10, 64)
		add, _ := strconv.ParseInt(cmd[2], 10, 64)

		a[from] = a[from] + add

		if to+1 <= arraysize {
			a[to+1] = a[to+1] - add
		}
	}

	var max int64
	var runningMax int64
	for i := 1; i <= int(arraysize); i++ {
		runningMax = runningMax + a[int64(i)]
		if max < runningMax {
			max = runningMax
		}

	}

	fmt.Println(max)
}
