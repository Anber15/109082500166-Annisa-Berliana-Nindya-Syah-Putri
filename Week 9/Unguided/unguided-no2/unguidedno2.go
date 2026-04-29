package main

func main() {

    arr := []int{2, 4, 6, 8, 10}
    N := len(arr)

    println("Isi array:")
    for i := 0; i < N; i++ {
        print(arr[i], " ")
    }
    println()

    println("Indeks ganjil:")
    for i := 0; i < N; i++ {
        if i%2 != 0 {
            print(arr[i], " ")
        }
    }
    println()

    println("Indeks genap:")
    for i := 0; i < N; i++ {
        if i%2 == 0 {
            print(arr[i], " ")
        }
    }
    println()

    x := 2
    println("Indeks kelipatan", x, ":")
    for i := 0; i < N; i++ {
        if i%x == 0 {
            print(arr[i], " ")
        }
    }
    println()

    idx := 2
    arr = append(arr[:idx], arr[idx+1:]...)
    println("Setelah dihapus:")
    for i := 0; i < len(arr); i++ {
        print(arr[i], " ")
    }
    println()

    sum := 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    mean := float64(sum) / float64(len(arr))
    println("Rata-rata:", mean)

    var total float64
    for i := 0; i < len(arr); i++ {
        diff := float64(arr[i]) - mean
        total += diff * diff
    }
    variance := total / float64(len(arr))
    println("Variansi:", variance)

    cari := 8
    freq := 0
    for i := 0; i < len(arr); i++ {
        if arr[i] == cari {
            freq++
        }
    }
    println("Frekuensi:", freq)
}