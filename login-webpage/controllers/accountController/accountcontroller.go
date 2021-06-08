package accountcontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./view/accountController/index.html")
	err := tmp.Execute(w, nil)
	if err != nil {
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("./view/accountController/index.html")
	tmp.Execute(w, nil)
}
