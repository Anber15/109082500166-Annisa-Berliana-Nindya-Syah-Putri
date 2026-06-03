package main

import "fmt"

func selectionSort(data []int, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		posMin := i

		for j := i + 1; j < jumlah; j++ {
			if data[j] < data[posMin] {
				posMin = j
			}
		}

		temp := data[posMin]
		data[posMin] = data[i]
		data[i] = temp
	}
}

func main() {
	var banyakKasus int
	fmt.Scan(&banyakKasus)

	for i := 0; i < banyakKasus; i++ {
		var jumlahData int
		fmt.Scan(&jumlahData)

		var angka []int

		for j := 0; j < jumlahData; j++ {
			var inputAngka int
			fmt.Scan(&inputAngka)
			angka = append(angka, inputAngka)
		}

		selectionSort(angka, jumlahData)

		for j := 0; j < jumlahData; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(angka[j])
		}

		fmt.Println()
	}
}