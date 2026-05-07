package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		for T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n *int) {
	*n = 0
	var bil int
	fmt.Scan(&bil)

	for *n < 2022 && !exist(*T, *n, bil) {
		T[*n] = bil
		*n++
		fmt.Scan(&bil)
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) {
			T3[*h] = T1[i]
			*h++
		}
	}
}

func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
