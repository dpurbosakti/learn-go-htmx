package main

import (
	"fmt"
	"learn-htmx/test"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")

	http.Handle("/", templ.Handler(component))
	http.Handle("/template", templ.Handler(headerTemplate(test.GetName())))
	fmt.Println("Listening on :3000")
	fmt.Println("hello air asd")
	http.ListenAndServe(":3000", nil)
}
