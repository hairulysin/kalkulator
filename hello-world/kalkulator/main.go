package main

import "fmt"

func Tambah(a, b float64) float64{
	return a + b
}
func Kurang(a, b float64) float64{
	return a - b
}
func Kali(a, b float64) float64{
	return a * b
}
func Bagi(a, b float64) float64{
	return a / b
}
func Modulus(a, b int) int{
	return a % b
}

func main() {
	hasil1 := Tambah(3,4)
	hasil2 := Kurang(9,2)
	hasil3 := Kali(3,3)
	hasil4 := Bagi(8,3)
	hasil5 := Modulus(9,2)

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
	fmt.Println(hasil4)
	fmt.Println(hasil5)
}
