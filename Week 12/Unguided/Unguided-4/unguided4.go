package main

import "fmt"

func urutInsertion(data []int, panjang int) {
	for i := 1; i < panjang; i++ {
		simpan := data[i]
		pos := i

		for pos > 0 && simpan < data[pos-1] {
			data[pos] = data[pos-1]
			pos = pos - 1
		}

		data[pos] = simpan
	}
}

func main() {
	var angka []int
	var nilai int

	for {
		fmt.Scan(&nilai)

		if nilai < 0 {
			break
		}

		angka = append(angka, nilai)
	}

	jumlah := len(angka)

	if jumlah > 0 {
		urutInsertion(angka, jumlah)
	}

	for i := 0; i < jumlah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(angka[i])
	}
	fmt.Println()

	if jumlah < 2 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		selisih := angka[1] - angka[0]
		konsisten := true

		for i := 1; i < jumlah-1; i++ {
			if angka[i+1]-angka[i] != selisih {
				konsisten = false
				break
			}
		}

		if konsisten {
			fmt.Printf("Data berjarak %d\n", selisih)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}