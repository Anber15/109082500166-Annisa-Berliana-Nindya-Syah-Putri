package main

import "fmt"

const batasData = 1000000

type dataSet [batasData]int

func urutkan(d *dataSet, ukuran int) {
	for i := 0; i < ukuran-1; i++ {
		minPos := i

		for j := i + 1; j < ukuran; j++ {
			if d[j] < d[minPos] {
				minPos = j
			}
		}

		d[i], d[minPos] = d[minPos], d[i]
	}
}

func cariMedian(d dataSet, ukuran int) float64 {
	tengah := ukuran / 2

	if ukuran%2 == 0 {
		return float64(d[tengah-1]+d[tengah]) / 2.0
	}
	return float64(d[tengah])
}

func main() {
	var angka dataSet
	var nilai, jumlah int

	fmt.Println("Input data masukan:")

	fmt.Scan(&nilai)

	for nilai != -5313541 {

		if nilai == 0 {
			urutkan(&angka, jumlah)
			fmt.Println("Median:")
			fmt.Println(cariMedian(angka, jumlah))

		} else {
			angka[jumlah] = nilai
			jumlah++
		}

		fmt.Scan(&nilai)
	}
}
