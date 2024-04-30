package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Failed to parse form %v", err)
			return
		}
		fmt.Fprintf(w, "POST request succesfull")
		name := r.FormValue("name")
		password := r.FormValue("password")
		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Passeord: %s\n", password)
		return
	}
	http.ServeFile(w, r, "./static/form.html")

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method isn't supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)
	fmt.Println("8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
