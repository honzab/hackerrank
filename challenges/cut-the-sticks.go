package main

import (
	"bufio"
	"fmt"
	"log"
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
	sticks := strings.Split(s.Text(), " ")
	if len(sticks) != int(size) {
		log.Fatalf("Invalid input")
	}

	isticks := make([]int64, len(sticks))
	for i := 0; i < len(sticks); i++ {
		isticks[i], err = strconv.ParseInt(sticks[i], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	c := 0
	cut := false
	for {
		noOfCut := 0
		cutlen := getCutSize(isticks)
		for i := 0; i < len(isticks); i++ {
			if isticks[i] > 0 {
				cut = true
				noOfCut = noOfCut + 1
				isticks[i] = isticks[i] - cutlen
			}
		}
		c = c + 1
		if !cut {
			break
		}
		cut = false
		fmt.Println(noOfCut)
	}
}

func getCutSize(is []int64) int64 {
	var min int64
	min = 10000
	for _, v := range is {
		if v == 0 {
			continue
		}
		if v <= min {
			min = v
		}
	}
	return min
}
