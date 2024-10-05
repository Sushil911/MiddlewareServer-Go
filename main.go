package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/open", logging(openhandler))
	http.HandleFunc("/close", logging(authentication(closehandler)))
	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error while starting the server")
	}
}

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method=%s URL=%s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer mysecrettoken" {
			fmt.Println("Not authenticated")
		}
		next.ServeHTTP(w, r)
	}
}

func openhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You are on the open route")
}
func closehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You are on the authenticated close route")
}
