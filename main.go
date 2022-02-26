package main

import "fmt"

func main() {
	var A []int = []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println("Результат: нету пары у ", Solution(A))
}

func Solution(A []int) int {
	len := len(A)
	var res int
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if A[i] == A[j] {
				break
			} else {
				res = A[i]
			}
		}
	}
	return res
}
