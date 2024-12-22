package main

import (
	"fmt"
)

//note: Standar library in golang,

func main() {
	fmt.Println("Halo, belajar Golang...")
	title := "Golang the best language"
	// todo: Quiz perulangan
	for index, letter := range title {
		letterString := string(letter)
		// (letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o")
		// if index%2 == 0 {
		// 	fmt.Println("index: ", index, " , letter: ", letterString)
		// }
		// switch letterString {
		// case "a", "i", "u", "e", "o":
		// 	fmt.Println("index: ", index, " , letter: ", letterString)
		// }
		// if index%2 == 0 {
		// 	switch letterString {
		// 	case "a", "i", "u", "e", "o":
		// 		fmt.Println("index: ", index, " , letter: ", letterString)
		// 	}
		// }
		if index%2 == 0 && (letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o") {
			fmt.Println("index: ", index, " , letter: ", letterString)
		}
	}
	// todo: Jika index tidak ingin digunakan, gunakan _
	// for _, letter := range title {
	// 	fmt.Println("Letter: ", string(letter))
	// }
	// for index, letter := range title {
	// 	fmt.Println("index: ", index, " , letter: ", string(letter))
	// }
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Halo, belajar Golang:", i)
	// 	i++
	// }
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Halo, belajar Golang:", i)
	// }
	// number := 5
	// switch number {
	// case 1:
	// 	fmt.Println("Satu")
	// case 2:
	// 	fmt.Println("Dua")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("Default")
	// }
	// score := 65
	// var grade string
	// if score < 50 {
	// 	grade = "E"
	// } else if score <= 60 {
	// 	grade = "D"
	// } else if score < 70 {
	// 	grade = "C"
	// } else {
	// 	grade = "null"
	// }
	// fmt.Println(grade)
	// age := 10
	// if age > 10 {
	// 	fmt.Println("Age more than 10")
	// } else {
	// 	fmt.Println("Age not more than 10")
	// }
	// var name string = "Golang"
	// notes: Langsung tipe data dan assign value ke variabel
	// age := 20
	// todo: Jika ingin reassign
	// age = 30
	// fmt.Println(name)
	// fmt.Println(age)
	// x := 10
	// y := 9
	// add := calculation.Add(x, y)
	// multiply := calculation.Multiply(x, y)
	// fullname := "Muhammad Reza Pahlevi Y"
	// fmt.Println(add)
	// fmt.Println(multiply)
	// fmt.Println(fullname)
	// fmt.Println("Hello world! Welcome in Golang!")
	// sentence := TestAja()
	// namaSaya := namaSaya()
	// fmt.Println(sentence)
	// fmt.Println(namaSaya)
}
