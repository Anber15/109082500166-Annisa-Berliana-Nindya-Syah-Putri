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
			hitung[x]++
			sah++
		}
	}

	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		if hitung[i] > hitung[ketua] {
			wakil = ketua
			ketua = i
		} else if hitung[i] > hitung[wakil] && i != ketua {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}