/*
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	usernames []string
	passwds =  map[string]string{}
	assignedSalt []int
	userSalt = map[string]int{}
)

func main() {
	ask_sign()
}

func ask_sign() {
	var ask string
	fmt.Print("Do you want to sign in or sign up to Quantum Tricks ?: ")
	fmt.Scan(&ask)
	ask = strings.ToLower(ask)

	if ask == "signup" || ask == "register"{
		createNewUser()
	} else if ask == "signin" {
		fmt.Println(userSignIn())
		ask_sign()
	} else {
		ask_sign()
	}
}

func createNewUser() {
	for true {
		var username string
		fmt.Print("Enter your username: ")
		fmt.Scan(&username)

		for _, i := range usernames {
			if username != i {

				fmt.Println("Note that the password should only be a string...! \n")
				fmt.Println("Enter Password: ")
				var passwd string
				fmt.Scan(&passwd)
				for true {
					var rePasswd string
					fmt.Println("Re-enter the password: ")
					fmt.Scan(&rePasswd)
					if passwd == rePasswd {
						fmt.Println("Congratulations your Quantum Tricks Account has been created successfully")
						fmt.Println("Please sign in to start working with your QT Account \n")
						saltAssign(passwd, username)
						usernames = append(usernames, username)
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
	}
}

func userSignIn() string {
	var username string
	fmt.Print("Enter Username: ")
	fmt.Scan(&username)

	for _, j := range usernames {
		if username == j {
			chk := passwordChk(username)
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

func passwordChk(username string) bool {
	fmt.Print("Enter your password: ")
	var ask_passwd string
	fmt.Scan(&ask_passwd)
	salt := findSaltOfUser(username)
	passwd := encryptPasswd(salt, ask_passwd)
	OPasswd := passwds[username]
	if passwd == OPasswd {
		return true
	}
	return false
}

func findSaltOfUser(user string) int {
	salt := userSalt[user]
	return salt
}

func encryptPasswd(salt int, passwd string) string {
	newPasswd := string(rune(salt)) + passwd
	return newPasswd
}

func saltAssign(passwd string, username string) {
	salt := random()
	for _, j := range assignedSalt {
		if salt != j {
			s := string(rune(salt)) + passwd
			assignedSalt = append(assignedSalt, salt)
			userSalt[username] = salt
			passwds[username] = s
		}
	}
}

func random() int {
	return rand.Intn(100000)
}
*/

package main

import "fmt"

type userStruct struct {
	id int
	FirstName string
	LastName string
	password string
	EmailAddress string
	Location string
}

var (
	users []*userStruct
	nid = 1
)

func main() {

}

func GetUsers() []*userStruct {
	return users
}

func addUser(u userStruct) userStruct {
	u.id = nid
	nid++
	users = append(users, &u)
	return u
}

func deleteUserById(id int) error {
	for i, u := range users {
		if u.id == id {
			users = append(users[:i], users[i+1:]...)
		}
	}

	return fmt.Errorf("User with ID '%d' does not exist", id)
}
