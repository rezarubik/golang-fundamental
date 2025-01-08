package management

import "fmt"

// note: Initiate Struct (Struct bisa dibilang hampir sama dengan model pada laravel)
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// note: Method
func (user User) Display() string {
	return fmt.Sprintf("Nama: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

// note: Struct dengan isi property struct yg lain
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

/* ------------------------ // start: Embedded struct ----------------------- */
func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println()
	fmt.Println("Users name")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
