package main

import "fmt"

func main() {
	type user struct {
		ID int
		FirstName string
		LastName string
		Password string
	}

	var u user
	u.ID = 1
	u.FirstName = "Go"
	u.LastName = "Lang"
	u.Password = "google"
	fmt.Println(u)
	fmt.Println(u.LastName)

	u2 := user{ID: 2, FirstName: "Python", LastName: "cool", Password: "Guido van Rossum"}
	fmt.Println(u2)
}
