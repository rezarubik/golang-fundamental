package main

import (
	//note: Standar library in golang,
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, belajar Golang...")
	var name string = "Golang"
	// notes: Langsung tipe data dan assign value ke variabel
	age := 20
	// todo: Jika ingin reassign
	age = 30
	fmt.Println(name)
	fmt.Println(age)
	x := 10
	y := 9
	add := calculation.Add(x, y)
	multiply := calculation.Multiply(x, y)
	fmt.Println(add)
	fmt.Println(multiply)
	// fmt.Println("Hello world! Welcome in Golang!")
	// sentence := TestAja()
	// namaSaya := namaSaya()
	// fmt.Println(sentence)
	// fmt.Println(namaSaya)
}
