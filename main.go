package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0)
	arr = append(arr, 1)
	for i, v := range arr {
		fmt.Printf("key: %v, value: %v\n", i, v)
	}
	fmt.Printf("val: %v\n", arr[0:])
	fmt.Printf("val: %v\n", arr[1:])
}
