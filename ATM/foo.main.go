package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Database")
	db, err := sql.Open("mysql", "quantum:qwer1234$S@tcp(localhost:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	fmt.Println("Successfully connected to MySQL Database 'testdb'")

	insert, err := db.Query("INSERT INTO UserDB (fname, lname, aaddhar, phone, DOB, age, address, email) VALUES ('suhas', 'k', 740731063536, 9834180012, '2006 02 28', 16, 'KE PH2 PNQ', 'suhas.ksv@gmail.com')")

	if err != nil {
		panic(err.Error())
	}

	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			panic(err.Error())
		}
	}(insert)

	fmt.Println("\nSuccessfully Updated Database\n\n")

}
