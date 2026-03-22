package main

import "fmt"

func hitungBiaya(jenis string, masuk, keluar int) int {
	total := 0
	for i := masuk; i < keluar; i++ {
		if i < 17 {
			if jenis == "motor" {
				total += 4000
			} else {
				total += 6000
			}
		} else {
			if jenis == "motor" {
				total += 5000
			} else {
				total += 7000
			}
		}
	}
	return total
}

func main() {
	var jenis string
	var masuk, keluar int
	total := 0
	no := 1

	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	for {
		fmt.Printf("\n*Kendaraan %d\n", no)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)
		total += biaya

		fmt.Printf("Biaya Parkir %s %d: %d\n", jenis, no, biaya)
		no++
	}

	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", total)
}