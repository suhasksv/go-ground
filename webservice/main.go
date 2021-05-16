package main

import (
	"github.com/pluralsight/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":8080", nil)
}
