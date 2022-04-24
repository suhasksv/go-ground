//package main
//
//import (
//	"fmt"
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//const attempts = 3
//
//func main() {
//	fmt.Println("Welcome to ATM Banking")
//
//	fmt.Print("Do you have a Banking Account? Have - yes, Don't have - no")
//	var ask1 string
//	fmt.Scan(&ask1)
//
//	if ask1 == "yes" {
//		accountHave()
//	}
//	accountCreate()
//}
//
//func accountHave() {
//
//}
//
//func accountCreate()  {
//
//}
package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}