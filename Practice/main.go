package main

import (
	"fmt"
	"net/http"
)
// Serve HTML file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "GEEEEEEEEEET")
	} else if r.Method == http.MethodPost { 
		r.ParseForm() // Parse the form data
		name := r.FormValue("name")
		fmt.Fprintf(w, "%s\n", name)
	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
	// fmt.Fprintln(w, "My Name Is Scarlx")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/user", Handler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
