package main

import "fmt"

// Interface is a Contract / Aturan Main
type BangunDatar interface {
	HitungLuas() int
}

// Cara Memenuhi Contract tersebut bagaimana? Maka dibutuhkan Struct
type Persegi struct {
	Sisi int
}
func (p Persegi) HitungLuas() int {
	return p.Sisi * p.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}
func (pp PersegiPanjang) HitungLuas() int {
	return pp.Panjang * pp.Lebar
}

func main() {
	bangunDatar1 := Persegi{ Sisi: 4 }
	fmt.Println(SeberapaLuas(bangunDatar1))
	
	bangunDatar2 := PersegiPanjang{ Panjang: 4, Lebar: 4}
	fmt.Println(SeberapaLuas(bangunDatar2))
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}