/*
package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"time"
	"os"
)

const TimeFormat = "February 26, 2006"
func main() {	
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Enter your Date of Birth(format - DDMMYY):")
	
	scanner.Scan()
	text := scanner.Text()
	texts := strings.Split(text, "")
	intText := make([]int, 0)
	
	for _, i := range texts {
		i, _ := strconv.Atoi(i)
		intText = append(intText, i)
	}
	fmt.Print(intText)
	
	dob := getDOB(intText[0], intText[1], intText[2])	
	fmt.Println(dob.Format(TimeFormat))
}

func getDOB(day, month, year int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
*/

package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main () {
	fmt.Println("MySQL Database")
	
	db, err := sql.Open("mysql", "root:qwer1234$S@tcp(localhost:3306)/testdb")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()
	
	fmt.Println("Successfully connected to MySQL Database 'testdb'")
	
	insert, err := db.Query("INSERT IN	TO UserDB VALUES('Suhas', 'K', 740731063536, 9834180012, '2006-02-28', 15, 'Kalpataru Estate Ph2 Pimple Gurav Pune', 'suhas.ksv@gmail.com')")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer insert.Close()

	fmt.Println("Successfully inserted")
}

