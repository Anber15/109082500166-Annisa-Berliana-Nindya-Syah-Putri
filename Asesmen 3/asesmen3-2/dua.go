package main

import "fmt"

type player struct {
	fullname string
	goals    int
	assists  int
}

const MAX = 1000

type list [MAX]player

func main() {
	var data list
	var n, i, j int
	var fName, lName string

	fmt.Println("Masukkan Data Input:")

	fmt.Scan(&n)

	for i = 0; i < n && i < MAX; i++ {
		fmt.Scan(&fName, &lName, &data[i].goals, &data[i].assists)
		data[i].fullname = fName + " " + lName
	}

	for i = 0; i < n-1; i++ {
		best := i

		for j = i + 1; j < n; j++ {
			if data[j].goals > data[best].goals ||
				(data[j].goals == data[best].goals && data[j].assists > data[best].assists) {
				best = j
			}
		}

		data[i], data[best] = data[best], data[i]
	}

	fmt.Println("\nHasil Sorting:")

	for i = 0; i < n; i++ {
		fmt.Println(data[i].fullname, data[i].goals, data[i].assists)
	}
}