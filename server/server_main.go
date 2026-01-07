package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/upload/resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "resource upload stub")
	})
	fmt.Println("Management server running on :8080 (scaffold)")
	http.ListenAndServe(":8080", nil)
}