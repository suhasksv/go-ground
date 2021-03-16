package main

import (
	"fmt"
	"gihub.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Suhas",
		LastName: "K",
	}
	fmt.Println(u)
}
