package main

import "fmt"

func insertionSort(arr []int, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {
		temp = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}
}

func hitungMedian(arr []int, n int) int {
	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
}

func main() {
	var data []int

	for {
		var input int
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {
			insertionSort(data, len(data))
			median := hitungMedian(data, len(data))
			fmt.Println(median)
		} else {
			data = append(data, input)
		}
	}
}