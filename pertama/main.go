package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang =", luas)

	persegi := Persegi{Sisi: 4}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas Persegi =", luas)

	asal := Asal{Angka: 5}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas Asal =", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}

// func SeberapaLuas2(bangunDatar Persegi) int {
// 	return bangunDatar.HitungLuas()
// }
