package main

import (
	"errors"
	"fmt"
	"pertama/management"
)

/* ------------------- // start: Struct sebagai parameter ------------------- */
// func displayUser(user User) string {
// 	return fmt.Sprintf("Nama: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
// }
/* -------------------- // end: Struct sebagai parameter -------------------- */

/* ------------------------ // start: Embedded struct ----------------------- */
// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count: %d", len(group.Users))
// 	fmt.Println()
// 	fmt.Println("Users name")
// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }
/* ------------------------ // emd: Embedded struct ----------------------- */

//note: Standar library in golang,

func main() {
	fmt.Println("Halo, belajar Golang...")
	/* ------------------------ // start: method ----------------------- */
	user := management.User{
		ID:        1,
		FirstName: "Reza",
		LastName:  "Pahlevi",
		Email:     "reza@gmail.com",
		IsActive:  true,
	}
	result := user.Display()
	fmt.Println(result)
	user2 := management.User{
		ID:        2,
		FirstName: "Nadiah",
		LastName:  "Tsamara",
		Email:     "nadiah@gmail.com",
		IsActive:  true,
	}
	fmt.Println(user2.Display())
	/* ------------------------ // end: method ----------------------- */

	/* ------------------------ // start: Embedded struct ----------------------- */
	// user := User{
	// 	1,
	// 	"Reza",
	// 	"Pahlevi",
	// 	"reza@gmail.com",
	// 	true,
	// }
	// user2 := User{
	// 	2,
	// 	"Nadiah",
	// 	"Tsamara",
	// 	"nadiah@gmail.com",
	// 	true,
	// }
	users := []management.User{user, user2}
	group := management.Group{
		Name:        "Gamer",
		Admin:       user,
		Users:       users,
		IsAvailable: true,
	}

	// displayGroup(group)
	group.DisplayGroup()

	/* ------------------------- // end: Embedded struct ------------------------ */

	/* ------------------- // start: Struct sebagai parameter ------------------- */
	// user := User{
	// 	1,
	// 	"Reza",
	// 	"Pahlevi",
	// 	"reza@gmail.com",
	// 	true,
	// }
	// displayUser := displayUser(user)
	// fmt.Println(displayUser)
	/* ------------------- // end: Struct sebagai parameter ------------------- */
	/* ------------------------- // start: Struct Basic ------------------------- */
	// user := User{}
	// todo: Mengisi sekaligus ketika pembuatan (bisa juga tanpa menuliskan atributenya tetapi harus urut berdasarkan atribut yg dibuat)
	// user := User{
	// 	ID:        1,
	// 	FirstName: "Reza",
	// 	LastName:  "Pahlevi",
	// 	IsActive:  true,
	// }
	// user := User{
	// 	1,
	// 	"Reza",
	// 	"Pahlevi",
	// 	true,
	// }
	// todo: Initiate value ke atribute:
	// user.ID = 1
	// user.FirstName = "Reza"
	// user.LastName = "Pahlevi"
	// user.IsActive = true
	// fmt.Println(user)
	/* ------------------------- // end: Struct Basic ------------------------- */

	/* ---------------------------- // todo: Function --------------------------- */
	/* ------------------------- // todo: Quiz function ------------------------- */
	// note: Quiz 1
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Println("total =", total)
	// fmt.Println("======")
	// note: Quiz 2
	// result, err := calculate(10, 2, "=")
	// if err != nil {
	// 	fmt.Println("Ooops, something error", err.Error())
	// }
	// fmt.Println("Result =", result, ", err =", err)
	// luas, keliling := calculate(10, 2)
	// note: Jika hanya perlu satu pengembalian: luas, _ := calculate(10,2)
	// fmt.Println("luas =", luas, ", keliling =", keliling)
	// resultAdd := add(10, 20)
	// fmt.Println(resultAdd)
	// sentence := printMyResult("Params 1")
	// fmt.Println(sentence)
	// printMyResult("Params 2")
	// printMyResult("Params 3")
	/* ---------------- // todo: Quiz Golang yang lebih Kompleks ---------------- */
	// todo: Hitung rata-rata
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// var total int
	// for _, score := range scores {
	// 	total = total + score
	// }
	// fmt.Println(total)
	// average := float64(total) / float64(len(scores))
	// fmt.Println(average)
	// note: Good scores >= 90
	// var goodScores []int
	// for _, score := range scores {
	// 	if score >= 90 {
	// 		goodScores = append(goodScores, score)
	// 	}
	// }
	// fmt.Println(goodScores)
	/* -------------------------- // todo: Slice of map ------------------------- */
	// students := []map[string]string{
	// 	{"name": "Agung", "score": "A"},
	// 	{"name": "Link", "score": "B"},
	// 	{"name": "Mario", "score": "C"},
	// }

	// for _, student := range students {
	// 	nameStudent := student["name"]
	// 	scoreStudent := student["score"]
	// 	fmt.Println(nameStudent, "get score", scoreStudent)
	// }
	/* ------------------------------ // todo: Map ------------------------------ */
	// map [key] tipe_data valuenya
	// note: Langsung assign
	// myMap := map[string]string{
	// 	"ruby":       "This is rubby",
	// 	"go":         "This is GO",
	// 	"javascript": "Hype",
	// }
	// var myMap map[string]int
	// myMap = map[string]int{}
	// myMap["GO"] = 10
	// myMap["Javascript"] = 9
	// myMap["Rubby"] = 8
	// todo: Map lebih lanjut
	// for key, value := range myMap {
	// 	fmt.Println("Key:", key, ", value:", value)
	// }
	// fmt.Println("=======")
	// delete(myMap, "ruby")
	// for key, value := range myMap {
	// 	fmt.Println("Key:", key, ", value:", value)
	// }
	// value, isAvailable := myMap["python"]
	// fmt.Println(value)
	// fmt.Println(isAvailable)
	// fmt.Println(myMap)
	// fmt.Println(myMap["GO"])
	/* ----------------------------- // todo: Slice ----------------------------- */
	// todo: Assign langsung
	// gamingConsole := []string{"PS 4", "PS 3", "PS 2", "PS 1"}
	// todo: Manual
	// var gamingConsole []string
	// gamingConsole = append(gamingConsole, "PS 4")
	// gamingConsole = append(gamingConsole, "PS 3")
	// gamingConsole = append(gamingConsole, "PS 2")
	// gamingConsole = append(gamingConsole, "PS 1")
	// for _, console := range gamingConsole {
	// 	fmt.Println("Console", console, ", panjang", len(gamingConsole))
	// }
	// fmt.Println(gamingConsole)
	/* ----------------------------- // todo: Array ----------------------------- */
	// var languages [5]string
	// var numbers [3]int
	// languages[0] = "GO"
	// languages[1] = "JavaScript"
	// languages[2] = "Ruby"
	// languages[3] = "PHP"
	// languages[4] = "Python"
	// todo: Bisa langsung juga
	// languages := [5]string{
	// 	"PHP",
	// 	"Javascript",
	// 	"Ruby",
	// 	"Python",
	// 	"GO",
	// }
	// note: Jika arraynya ingin tidak didefinisikan berapa panjangnya
	// languages := [...]string{
	// 	"PHP",
	// 	"Javascript",
	// 	"Ruby",
	// 	"Python",
	// 	"GO",
	// 	"Java",
	// }
	// numbers[0] = 1
	// numbers[1] = 2
	// numbers[2] = 3
	// lengthLanguages := len(languages)
	// lengthNumbers := len(numbers)
	// fmt.Println(languages)
	// fmt.Println("Panjang array languages adalah", lengthLanguages)
	// fmt.Println(numbers)
	// fmt.Println("Panjang array numbers adalah", lengthNumbers)
	// for index, lang := range languages {
	// 	fmt.Println("index:", index, ", language:", lang)
	// }

	/* ------------------------ // todo: Quiz perulangan ------------------------ */
	// title := "Golang the best language"
	// for index, letter := range title {
	// 	letterString := string(letter)
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
	// if index%2 == 0 && (letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o") {
	// 	fmt.Println("index: ", index, " , letter: ", letterString)
	// }
	// }
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

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " belajar Golang"
// 	// fmt.Println(sentence)
// 	return newSentence
// }

// todo: Quiz function
// todo: Quiz 1
func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}
	return total
}

// todo: Quiz 2
// todo: My Version
// func calculate(x, y int, symbol string) (result int, err string) {
// 	err = "-"
// 	if symbol == "+" {
// 		result = x + y
// 	} else if symbol == "-" {
// 		result = x - y
// 	} else if symbol == "*" {
// 		result = x * y
// 	} else if symbol == "/" {
// 		result = x / y
// 	} else {
// 		err = "Symbol is invalid"
// 	}
// 	return result, err
// }

// todo: Mas Agung Version
func calculate(x, y int, operation string) (int, error) {
	var result int
	var errorResult error
	switch operation {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	default:
		errorResult = errors.New("Unknown Operation")
	}
	return result, errorResult
}

// todo: Struct sebagai parameter

// todo: Multiple return (Lebih dari satu pengembalian)
// note: Predefined return value (bisa juga seperti ini)
// func calculate(panjang int, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)
// 	return luas, keliling
// }

// func calculate(panjang int, lebar int) (int, int) {
// 	luas := panjang * lebar
// 	keliling := 2 * (panjang + lebar)
// 	return luas, keliling
// }

// func add(x, y int) int {
// 	// note: Sama dengan add(x int, y int)
// 	result := x + y
// 	return result
// }
