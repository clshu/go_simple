package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to AA website</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>Email us at </p><p><a href=\"mailto: support@lenslock.com\">support@lenslock.com</a></p>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1><p>What's suuport email? </p><p>Answer: <a href=\"mailto: support@lenslock.com\">support@lenslock.com</a></p>")
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Page Not Found</h1><p>Email us if you keep receiving an invalid page.</p><p><a href=\"mailto: support@lenslock.com\">support@lenslock.com</a></p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(handle404)
	http.ListenAndServe(":4000", r)
}
