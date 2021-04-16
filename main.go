package main

import (
	"fmt"
)

func main() {
	_ = Bubble([]int{1, 3, 4, 67, 5, 39, 77}, 7)
}

func Bubble(a []int, n int) []int {

	var temp int
	for k := 0; k < n-1; k++ {
		for i := 0; i < n-k-1; i++ {
			if a[i] > a[i+1] {
				temp = a[i]
				a[i] = a[i+1]
				a[i+1] = temp
			}
		}
	}
	fmt.Println(a)

	return a
}
