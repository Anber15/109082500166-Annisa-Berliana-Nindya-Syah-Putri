package main

import "fmt"

func sortNaik(data []int, panjang int) {
	for i := 0; i < panjang-1; i++ {
		posMin := i
		for j := i + 1; j < panjang; j++ {
			if data[j] < data[posMin] {
				posMin = j
			}
		}
		temp := data[posMin]
		data[posMin] = data[i]
		data[i] = temp
	}
}

func sortTurun(data []int, panjang int) {
	for i := 0; i < panjang-1; i++ {
		posMax := i
		for j := i + 1; j < panjang; j++ {
			if data[j] > data[posMax] {
				posMax = j
			}
		}
		temp := data[posMax]
		data[posMax] = data[i]
		data[i] = temp
	}
}

func main() {
	var jumlahKasus int
	fmt.Scan(&jumlahKasus)

	for i := 0; i < jumlahKasus; i++ {
		var banyakData int
		fmt.Scan(&banyakData)

		var dataGanjil []int
		var dataGenap []int

		for j := 0; j < banyakData; j++ {
			var angkaInput int
			fmt.Scan(&angkaInput)

			if angkaInput%2 != 0 {
				dataGanjil = append(dataGanjil, angkaInput)
			} else {
				dataGenap = append(dataGenap, angkaInput)
			}
		}

		sortNaik(dataGanjil, len(dataGanjil))
		sortTurun(dataGenap, len(dataGenap))

		awal := true

		for j := 0; j < len(dataGanjil); j++ {
			if !awal {
				fmt.Print(" ")
			}
			fmt.Print(dataGanjil[j])
			awal = false
		}

		for j := 0; j < len(dataGenap); j++ {
			if !awal {
				fmt.Print(" ")
			}
			fmt.Print(dataGenap[j])
			awal = false
		}

		fmt.Println()
	}
}