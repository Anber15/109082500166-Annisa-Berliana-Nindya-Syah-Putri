package main

import "fmt"

func SequentialSearch(arrBuah [5]string, dataDicari string) int {
	idx := -1
	for i := 0; i < len(arrBuah); i++ {
		if arrBuah[i] == dataDicari {
			idx = i
			break
		}
	}
	return idx
}

func main() {
	var arrBuah [5]string

	for i := 0; i < len(arrBuah); i++ {
		fmt.Printf("Masukkan data buah indeks ke-%d: ", i)
		fmt.Scan(&arrBuah[i])
	}

	var dataCari string
	fmt.Println("Masukkan data buah yang ingin dicari:")
	fmt.Scan(&dataCari)

	index_data := SequentialSearch(arrBuah, dataCari)

	if index_data != -1 {
		fmt.Printf("Data %s ditemukan pada indeks ke-%d!", dataCari, index_data)
	} else {
		fmt.Printf("Data %s tidak ditemukan!", dataCari)
	}
}