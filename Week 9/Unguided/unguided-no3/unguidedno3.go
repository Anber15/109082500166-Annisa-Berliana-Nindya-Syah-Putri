package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	type Match struct {
		a int
		b int
	}

	var data []Match

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	i := 1

	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		data = append(data, Match{skorA, skorB})
		i++
	}

	for j, m := range data {
		if m.a > m.b {
			fmt.Printf("Hasil %d : %s\n", j+1, klubA)
		} else if m.b > m.a {
			fmt.Printf("Hasil %d : %s\n", j+1, klubB)
		} else {
			fmt.Printf("Hasil %d : Draw\n", j+1)
		}
	}

	fmt.Println("Pertandingan selesai")
}