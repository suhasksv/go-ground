//package main
//
//import (
//	"html/template"
//	"net/http"
//)
//
//type user struct {
//	UserName string
//	Password string
//	FirstName string
//	LastName string
//}
//
//var tpl *template.Template
//var userDB = map[string]user{} // user ID, user
//var sessionDB = map[string]user{} // session ID, user ID
//
//func main() {
//	http.HandleFunc("/", index)
//	http.HandleFunc("/login", login)
//	http.ListenAndServe("8080", nil)
//}
//
//func init() {
//	tpl = template.Must(template.ParseGlob("templates/"))
//	userDB["qgk@qt.in"] = user{"QuantumG", "supersecret123", "Suhas", "K"}
//}
//
//func index(w *http.ResponseWriter, r *http.Request) {
//
//}
//
//func login(w *http.ResponseWriter, r *http.Request) {
//
//}

package main

import (
	"controllers/accountcontroller"
	// "login-page-api-server/my-utils/uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/accounts", accountcontroller.Index)
	http.HandleFunc("/accounts/", accountcontroller.Index)
	http.HandleFunc("/accounts/login", accountcontroller.Index)
	http.HandleFunc("/accounts/index", accountcontroller.Login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
