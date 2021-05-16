package controllers

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON1(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := xml.NewEncoder(w)
	enc.Encode(data)
}