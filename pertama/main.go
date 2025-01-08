package main

import "fmt"

func main() {
	numberA := 5
	numberB := &numberA //note: menyimpan alamat numberA di memory komputer
	fmt.Println(numberA)
	fmt.Println(numberB)
	// note: Deferensing (dereference)
	fmt.Println(*numberB)

	// todo: reassign numberB
	*numberB = 10
	fmt.Printf("NumberB : %d, numberA: %d", *numberB, numberA)
	// fmt.Println(*numberB)
	// fmt.Println(numberA)
}
