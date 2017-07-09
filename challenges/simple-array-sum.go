package main

import "fmt"

func main() {
	var size, sum, j int64
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(size); i++ {
		_, err := fmt.Scanf("%d", &j)
		if err != nil {
			panic(err)
		}
		sum += j
	}
	fmt.Println(sum)
}
