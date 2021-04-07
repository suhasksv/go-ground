package main

import (
	"fmt"
)

var usernames = []string{}
var passwds =  map[string]string{}
var assignedSalt = []int{}
var userSalt = map[string]int{}
var salt_qt = []int{}

func main() {
	salt_tp()
	ask_sign()
}

func ask_sign() {
	var ask string
	fmt.Print("Do you want to sign in or sign up to Quantum Tricks ?: ")
	fmt.Scan(&ask)
	signin := ask == "signin"
	signup := ask == "signup"
	if signup == true {
		create_new_user()
	} else if signin == true {
		is1 := user_sign_in()
		fmt.Println(is1)
		ask_sign()
	} else {
		ask_sign()
	}
	return
}

func create_new_user() {
	for true {
		var username string
		fmt.Print("Enter your username: ")
		fmt.Scan(&username)
		for _, i := range usernames {
			if username != i {
				usernames = append(usernames, username)

				fmt.Println("Note that the password should only be a string...! \n")
				fmt.Println("Enter Password: ")
				var passwd string
				fmt.Scan(&passwd)
				for true {
					var repasswd string
					fmt.Println("Re-enter the password: ")
					fmt.Scan(&repasswd)
					if passwd == repasswd {
						fmt.Println("Congratulations your Quantum Tricks Account has been created successfully")
						fmt.Println("Please sign in to start working with your QT Account \n")
						salt_assign(passwd, username)
						ask_sign()
						break
					}
					fmt.Println("Password is not matching, please try again")
					continue
				}
			}
			println("Username is already taken, please try another username...\n")
			continue
		}
		continue
	}
}

func user_sign_in() string {
	var username string
	fmt.Print("Enter Username: ")
	fmt.Scan(&username)

	for _, j := range usernames {
		if username == j {
			chk := password_chk(username)
			if chk == true {
				return "login success"
			} else if chk == false {
				return "login failed"
			} else {
				return "Could not find user"
			}
		}
	}
	return ""
}

func password_chk(username string) bool {
	fmt.Print("Enter your password: ")
	var ask_passwd string
	fmt.Scan(&ask_passwd)
	salt := find_salt_of_user(username)
	passwd := encrypt_passwd(salt, ask_passwd)
	passw := passwds[username]
	if passwd == passw {
		return true
	} else {
		return false
	}
}

func find_salt_of_user(user string) int {
	salt := userSalt[user]
	return salt
}

func encrypt_passwd(salt int, passwd string) string {
	new_passwd := string(salt) + passwd
	return new_passwd
}

func salt_assign(passwd string, username string) {
	for true {
		for i := 0; i <= 10000; i++ {
			salt := salt_qt[i]
			for _, j := range assignedSalt {
				if salt != j {
					s := string(salt) + passwd
					assignedSalt = append(assignedSalt, salt)
					userSalt[username] = salt
					passwds[username] = s
				}
			}
		}
	}
}

func salt_tp() {
	for i := 0; i <= 10000; i++ {
		salt_qt = append(salt_qt, i)
	}
}
