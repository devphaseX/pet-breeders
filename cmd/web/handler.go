package main

import (
	"fmt"
	"net/http"
)

func (a *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
