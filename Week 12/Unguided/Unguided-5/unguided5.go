package main

import "fmt"

const maxData int = 7919

type Buku struct {
	kode, judul, penulis, penerbit string
	jumlah, tahun, nilai           int
}

type KoleksiBuku [maxData]Buku

func inputBuku(data *KoleksiBuku, total *int) {
	fmt.Scan(total)

	for i := 0; i < *total; i++ {
		fmt.Scan(&data[i].kode, &data[i].judul,
			&data[i].penulis,
			&data[i].penerbit, &data[i].jumlah,
			&data[i].tahun, &data[i].nilai)
	}
}

func tampilTerbaik(data KoleksiBuku, total int) {
	if total == 0 {
		return
	}

	maks := data[0].nilai
	idx := 0

	for i := 1; i < total; i++ {
		if data[i].nilai > maks {
			maks = data[i].nilai
			idx = i
		}
	}

	fmt.Printf("%s %s %s %d\n", data[idx].judul,
		data[idx].penulis, data[idx].penerbit, data[idx].tahun)
}

func urutData(data *KoleksiBuku, total int) {
	for i := 1; i < total; i++ {
		simpan := data[i]
		pos := i

		for pos > 0 && simpan.nilai > data[pos-1].nilai {
			data[pos] = data[pos-1]
			pos--
		}

		data[pos] = simpan
	}
}

func tampil5Teratas(data KoleksiBuku, total int) {
	batas := 5

	if total < 5 {
		batas = total
	}

	for i := 0; i < batas; i++ {
		fmt.Println(data[i].judul)
	}
}

func cariBerdasarkanRating(data KoleksiBuku, total int, target int) {
	kiri := 0
	kanan := total - 1
	ketemu := false
	posisi := -1

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2

		if data[tengah].nilai == target {
			ketemu = true
			posisi = tengah
		} else if target > data[tengah].nilai {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu {
		fmt.Printf("%s %s %s %d %d %d\n", data[posisi].judul,
			data[posisi].penulis, data[posisi].penerbit, data[posisi].tahun,
			data[posisi].jumlah, data[posisi].nilai)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var dataBuku KoleksiBuku
	var jumlah, cari int

	inputBuku(&dataBuku, &jumlah)
	tampilTerbaik(dataBuku, jumlah)
	urutData(&dataBuku, jumlah)
	tampil5Teratas(dataBuku, jumlah)

	fmt.Scan(&cari)
	cariBerdasarkanRating(dataBuku, jumlah, cari)
}