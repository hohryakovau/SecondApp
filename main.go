package main

import "fmt"

func main() {
<<<<<<< HEAD
	var A []int = []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println("Результат: нету пары у ", Solution(A))
=======
	var B []int = []int{1, 1, 3, 9, 3, 9, 7, 9, 9, 3, 3, 3, 3, 3, 3, 3, 3}

	fmt.Println(Solution(B))

>>>>>>> 0c262eb (Initial commit)
}

func Solution(A []int) int {
	len := len(A)
	var res int
<<<<<<< HEAD
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if A[i] == A[j] {
				break
			} else {
				res = A[i]
			}
		}
=======
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
>>>>>>> 0c262eb (Initial commit)
	}
	return res
}
