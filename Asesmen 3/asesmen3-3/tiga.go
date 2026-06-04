package main

import "fmt"

const MAXDATA = 1000000

type party struct {
	id   int
	vote int
}

type listParty [MAXDATA]party

func cari(arr listParty, n, x int) int {
	for i := 0; i < n; i++ {
		if arr[i].id == x {
			return i
		}
	}
	return -1
}

func main() {
	var data listParty
	var nData, x int

	fmt.Println("Masukkan proses input suara:")

	fmt.Scan(&x)

	for x != -1 {
		idx := cari(data, nData, x)

		if idx == -1 {
			data[nData].id = x
			data[nData].vote = 1
			nData++
		} else {
			data[idx].vote++
		}

		fmt.Scan(&x)
	}

	for i := 0; i < nData-1; i++ {
		for j := i + 1; j < nData; j++ {
			if data[j].vote > data[i].vote {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

	fmt.Println("\nHasil perhitungan suara:")

	for i := 0; i < nData; i++ {
		fmt.Printf("%d(%d) ", data[i].id, data[i].vote)
	}
}