package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(a arrBalita, n int, min, max *float64) {
	*min = a[0]
	*max = a[0]

	for i := 1; i < n; i++ {
		if a[i] < *min {
			*min = a[i]
		}
		if a[i] > *max {
			*max = a[i]
		}
	}
}

func rerata(a arrBalita, n int) float64 {
	jumlah := 0.0
	for i := 0; i < n; i++ {
		jumlah += a[i]
	}
	return jumlah / float64(n)
}

func main() {
	var a arrBalita
	var n int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&a[i])
	}

	var min, max float64
	hitungMinMax(a, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)

	r := rerata(a, n)
	fmt.Printf("Rerata berat balita: %.2f kg\n", r)
}