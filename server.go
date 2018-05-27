package main 

import (
	"fmt"
	"log"
	"net/http"
	
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	
}


func saveHandler(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query()["input"]
	fmt.Println("input = ", val)
}

func main () {

	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/retrieve/", retrieveHandler)
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}