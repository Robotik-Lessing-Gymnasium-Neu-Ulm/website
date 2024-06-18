package main

import "net/http"

func main() {
	// Just a simple webserver for testing
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.ListenAndServe(":8080", nil)
}
