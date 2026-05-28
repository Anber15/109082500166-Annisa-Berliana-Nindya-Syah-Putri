package main

import "fmt"

func main() {
	var x int
	total := 0
	sah := 0

	var hitung [21]int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			hitung[x] = hitung[x] + 1
			sah++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)

	for i := 1; i <= 20; i++ {
		if hitung[i] != 0 {
			fmt.Println(i, ":", hitung[i])
		}
	}
}