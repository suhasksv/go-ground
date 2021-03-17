package main

import "fmt"

var usernames = []string{}
var passwds =  map[string]int{}
var assignedSalt = []int{}
var userSalt = map[string]int{}

const (
	True = true
	False = false
)

func main() {
	ask_sign()
}

func ask_sign() {
	var ask string
	fmt.Print("Do you want to sign in or sign up to Quantum Tricks ? : ")
	fmt.Scan(&ask)
	if ask == "sign up" {
		create_new_user()
	} else if ask == "sign up" {
		fmt.Println(user_sign_in)
		ask_sign()
	} else {
		ask_sign()
	}
}

func create_new_user() {

}

func user_sign_in() string {
	var username string
	fmt.Print("Enter Username: ")
	fmt.Scanln(&username)

	for i, j := range usernames {
		if username == j {
			chk := password_chk(username)
			if chk == True {
				return "login success"
			} else if chk == False {
				return "login failed"
			} else {
				return "Could not find user"
			}
		}
	}
	return ""
}

func password_chk(username string) string {
	fmt.Print("Enter your password: ")
	var ask_passwd string
	fmt.Scanln(&ask_passwd)
	salt := find_salt_of_user(username)
}
