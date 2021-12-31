package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Database")

	db, err := sql.open("mysql", "quantum:qwer1234$S@tcp(127.0.0.1:3306)/testdb")
}
