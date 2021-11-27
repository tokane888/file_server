package main

import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.ListenAndServe(":8081", nil)
}
