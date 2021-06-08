package main

import (
	"fmt"
	"github.com/pluralsight/webservice/controllers"
	"net/http"
)

func main() {
	fmt.Println("Server is starting")
	controllers.RegisterControllers()
	fmt.Println("Server has started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server encountered an Error")
		panic(err)
	}
}
