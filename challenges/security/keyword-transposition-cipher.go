package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func dedup(s string) string {
	var seen, out string
	for _, v := range s {
		if strings.ContainsRune(seen, v) {
			continue
		}
		out = out + string(v)
		seen = seen + string(v)
	}
	return out
}

// Prints out the alphabet / substitution alphabet
// with the desired wrapping for easy reading
func prn(s string, wrap int) {
	fmt.Printf("(%d) %d\n", len(s), wrap)
	for k, v := range strings.Split(s, "") {
		if k%wrap == 0 {
			fmt.Println("")
		}
		fmt.Printf("%s", string(v))
	}
	fmt.Println("")
}

func load(kw string) string {
	dedupKw := dedup(kw)
	intermediate := dedupKw

	for i := 0; i < len(alphabet); i++ {
		if !strings.ContainsRune(intermediate, rune(alphabet[i])) {
			intermediate = intermediate + string(alphabet[i])
		}
	}

	keyslice := strings.Split(dedupKw, "")
	sort.Strings(keyslice)

	var substitution string
	step := len(dedupKw)
	for i := 0; i < step; i++ {
		idx := strings.Index(intermediate, keyslice[i])
		for j := idx; j < len(alphabet); j += step {
			substitution = substitution + string(intermediate[j])
		}
	}
	return substitution
}

func decode(s, key string) string {
	var out string
	sub := load(key)
	for _, v := range s {
		idx := strings.Index(sub, string(v))
		if idx == -1 {
			out = out + " "
			continue
		}
		out = out + string(alphabet[idx])
	}
	return out
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	size, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(size); i++ {
		s.Scan()
		kw := s.Text()
		s.Scan()
		line := s.Text()
		decoded := decode(line, kw)
		fmt.Println(decoded)
	}
}
