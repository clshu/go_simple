package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to AA website</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Contact Page</h1><br><h3>Please Contact <a href=\"mailto: support@lenslock.com\">support@lenslock.com</a></h3")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":4000", nil)
}