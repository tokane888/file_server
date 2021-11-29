package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
	}
}
