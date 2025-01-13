package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

type Gamer struct {
	Name  string
	Games []string
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory setelah diubah:", &old)
	fmt.Println("Nilai setelah diubah:", *old)
}

func graduate(student *Student) {
	student.Name = student.Name + ", S.Tr.Kom"
}

func (student *Student) graduate() {
	student.Name = student.Name + ", S.Tr.Kom"
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	/* --------------------- // start: Quiz --------------------- */
	gamer := Gamer{
		Name: "Zelda",
	}
	gamer.AddGame("PS 3")
	gamer.AddGame("PS 4")
	gamer.AddGame("PS 5")

	fmt.Println("List Game")
	for _, game := range gamer.Games {
		fmt.Println(game)
	}
	/* --------------------- // end: Quiz --------------------- */

	/* --------------------- // start: Method dengan pointer receiver --------------------- */
	// student := Student{
	// 	ID:   1,
	// 	Name: "Reza Pahlevi",
	// 	GPA:  3.45,
	// }
	// fmt.Println("Nama sebelum:", student.Name)
	// student.graduate()
	// fmt.Println("Nama sesudah:", student.Name)
	/* --------------------- // end: Method dengan pointer receiver --------------------- */

	/* --------------------- // start: Pointer struct sebagai parameter --------------------- */
	// student := Student{
	// 	ID:   1,
	// 	Name: "Reza Pahlevi",
	// 	GPA:  3.45,
	// }
	// fmt.Println("Nama sebelum:", student.Name)
	// graduate(&student)
	// fmt.Println("Nama sesudah:", student.Name)
	/* --------------------- // end: Pointer struct sebagai parameter --------------------- */

	/* --------------------- // start: Contoh kasus penggunaan pointer --------------------- */
	// note: Mengubah nilai suatu variable
	// number := 5
	// fmt.Println("Alamat memory number awal:", &number)
	// fmt.Println("Nilai awal:", number)

	// change(&number, 100)
	// fmt.Println("Alamat memory number akhir:", &number)
	// fmt.Println("Nilai akhir:", number)
	/* --------------------- // end: Contoh kasus penggunaan pointer --------------------- */

	/* --------------------- // start: Pointer lebih lanjut --------------------- */
	// var numberA int = 5
	// var numberB *int = &numberA
	// fmt.Println("number A:", numberA)
	// fmt.Println("number B (alamat numberA):", numberB)
	// fmt.Println("number B (value numberA):", *numberB)
	// numberA = 20
	// fmt.Println("number A:", numberA)
	// fmt.Println("number B (alamat numberA):", numberB)
	// fmt.Println("number B (value numberA):", *numberB)
	/* ---------------------- // end: Pointer lebih lanjut ---------------------- */

	// numberA := 5
	// numberB := &numberA //note: menyimpan alamat numberA di memory komputer
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// note: Deferensing (dereference)
	// fmt.Println(*numberB)

	// todo: reassign numberB
	// *numberB = 10
	// fmt.Printf("NumberB : %d, numberA: %d", *numberB, numberA)
	// fmt.Println(*numberB)
	// fmt.Println(numberA)
}
