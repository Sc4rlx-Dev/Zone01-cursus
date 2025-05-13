package main

import (
	"fmt"
	"html/template"
	"net/http"

	asciiartweb "ascii-art/ascii-art"
	"ascii-art/function"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "" {
			function.Handle405(w)
			return
		}
		fs.ServeHTTP(w, r)
	})))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			function.Handle404(w)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			function.Handle405(w)
			return
		}

		if err := r.ParseForm(); err != nil {
			function.Handle400(w)
			return
		}

		userInput := r.FormValue("user_input")
		option := r.FormValue("option")
		if len(userInput) > 402 || len(userInput) < 1 {
			function.Handle400(w)
			return
		}
		if userInput == "" || option == "" {
			function.Handle400(w)
			return
		}
		asciiArt, err := asciiartweb.AsciiArt(option, userInput)
		if err != nil {
			function.Handle500(w)
			return
		}
		tp, _ := template.ParseFiles("result.html")
		tp.Execute(w, asciiArt)
	})

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
