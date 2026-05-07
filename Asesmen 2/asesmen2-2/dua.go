package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func main() {
	var data arrayMahasiswa
	var N int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)

		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	var cariNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&cariNIM)

	nilaiPertama := cariNilaiPertama(data, N, cariNIM)

	nilaiTerbesar := cariNilaiTerbesar(data, N, cariNIM)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", cariNIM, nilaiPertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", cariNIM, nilaiTerbesar)
}

func cariNilaiPertama(data arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if data[i].NIM == nim {
			return data[i].nilai
		}
	}
	return -1 
}

func cariNilaiTerbesar(data arrayMahasiswa, N int, nim string) int {
	maxNilai := -1
	for i := 0; i < N; i++ {
		if data[i].NIM == nim {
			if data[i].nilai > maxNilai {
				maxNilai = data[i].nilai
			}
		}
	}
	return maxNilai
}