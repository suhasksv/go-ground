package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

const (
	TimeFormat = "2006-01-02"
)

var (
	fname   string
	lname   string
	aaddhar int
	phone   int
	DOB     string
	age     int
	address string
	email   string
)

func main() {
	fmt.Println("Welcome to ATM Bank")
	fmt.Println("MySQL Database")

	CREATE_AN_ACCOUNT()
}

func CREATE_AN_ACCOUNT() {
	fmt.Print("Enter your First Name:")
	fmt.Scan(&fname)
	fmt.Print("Enter your Last Name:")
	fmt.Scan(&lname)
	fmt.Print("Enter your AADDHAR Number:")
	fmt.Scan(&aaddhar)

	fmt.Print("Enter your Date of Birth(format - YYYY MM DD):")
	var day, month, year int
	fmt.Scan(&year, &month, &day)
	dob := getDOB(year, month, day)
	DOB = dob.Format(TimeFormat)

	fmt.Print("Enter your age:")
	fmt.Scan(&age)
	fmt.Print("Enter your Mobile or Phone Number:")
	fmt.Scan(&phone)
	fmt.Print("Enter you Primary Email Address:")
	fmt.Scan(&email)
	fmt.Print("Enter your Address:")
	fmt.Scan(&address)

	if verify(month, year) != true {
		fmt.Println("\nDetails verification failed\nThere are a few errors in the values or missed some of them....\nPlease Try Again...!\n")
		CREATE_AN_ACCOUNT()
	}
	fmt.Println("\nPlease Validate the Entered Values ...")
	fmt.Printf("\n\n First Name: %s\n Last Name: %s\n Aaddhar: %d\n Phone Number: %d\n Date of Birth: %s\n Age: %d\n Address: %s\n\n", fname, lname, aaddhar, phone, DOB, age, address)
	time.Sleep(500 * time.Millisecond)

	var ask string
	fmt.Print("You sure you want to Proceed with these inputs? :")
	fmt.Scan(&ask)
	INSERTINTOTABLE()
	if strings.ToLower(ask) != "no" {
		fmt.Println("Fantastic...!\nYour ATM BANK Account has been Successfully Created")
	} else {
		fmt.Println("Thanks for visiting ATM Bank")
	}
}

func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}

func verify(month, year int) bool {
	var g int
	// username should not be unique
	if DIGITS(aaddhar) == 12 {
		g++
	}

	CURRENTDATE := time.Now()
	if CURRENTDATE.Year() > year {
		g++
	}

	if CURRENTDATE.Year()-year == age {
		g++
	}

	if DIGITS(phone) == 10 {
		g++
	}

	if g == 4 {
		return true
	}
	return false
}

func DIGITS(n int) int {
	count := 0
	if n == 0 {
		return 1
	}
	for n != 0 {
		count += 1
		n /= 10
		n = int(n)
	}
	return count
}

func INSERTINTOTABLE() {
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

	insert, err := db.Query("INSERT INTO UserDB (fname, lname, aaddhar, phone, DOB, age, address, email) VALUES (fname, lname, aaddhar, phone, DOB, age, address, email)")

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

/*
func calcAge(year, cyear) int {
	return cyear - year
}
*/
