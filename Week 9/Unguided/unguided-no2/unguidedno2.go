package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var a [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&a[i])
	}

	var wadah [1000]float64
	idx := 0
	sum := 0.0
	cnt := 0

	for i := 0; i < x; i++ {
		sum += a[i]
		cnt++

		if cnt == y {
			wadah[idx] = sum
			idx++
			sum = 0
			cnt = 0
		}
	}

	if cnt > 0 {
		wadah[idx] = sum
		idx++
	}

	total := 0.0
	for i := 0; i < idx; i++ {
		fmt.Printf("%.2f ", wadah[i])
		total += wadah[i]
	}
	fmt.Println()

	fmt.Printf("%.2f\n", total/float64(idx))
}