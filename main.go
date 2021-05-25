package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to AA website</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Contact Page</h1><br><h3>Please Contact <a href=\"mailto: support@lenslock.com\">support@lenslock.com</a></h3>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Page Not Found</h1><p>Email us if you keep receiving an invalid page.</p><p><a href=\"mailto: support@lenslock.com\">support@lenslock.com</a></p>")
	}
}

func handleHello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", handleHello)
	// mux := &http.ServeMux{}
	// mux.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":4000", router)
}
