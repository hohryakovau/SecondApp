package main

import "fmt"

func main() {

	var B []int = []int{1, 1, 3, 9, 3, 9, 7, 9, 9, 3, 3, 3, 3, 3, 3, 3, 3}

	fmt.Println(Solution(B))

}

func Solution(A []int) int {
	len := len(A)
	var res int
	var z int
	var arr = make([]int, len)
	for i := 0; i < len; i++ {
		z = 0
		for j := 0; j < len; j++ {
			if A[i] == A[j] {
				z = z + 1
			}
		}
		arr[i] = z
	}
	for y, x := range arr {
		if x == 1 {
			res = A[y]
		}
	}
	return res

}
