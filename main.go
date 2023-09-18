package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello!, Naman")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Error: %v", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful!\n")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	fmt.Fprintf(w, "First Name: %v\n", firstname)
	fmt.Fprintf(w, "Last Name: %v\n", lastname)	
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server on PORT 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}	
}