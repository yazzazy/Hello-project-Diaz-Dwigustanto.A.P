package main

import "fmt"

func main() {
	var s string
	var a, b, hasil_penjumlahan int

	fmt.Scan(&s, &a, &b)
	hasil_penjumlahan = a + b

	fmt.Print("Kata =", s)
	fmt.Print("Jumlah=", hasil_penjumlahan)
}
