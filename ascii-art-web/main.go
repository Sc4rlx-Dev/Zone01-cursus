package main

import (
	"ascii-art/function"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handle404(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			handle400(w, r)
			return
		}

		if err := r.ParseForm(); err != nil {
			handle400(w, r)
			return
		}

		userInput := r.FormValue("user_input")
		option := r.FormValue("option")
		if userInput == "" || option == "" {
			handle400(w, r)
			return
		}

		asciiArt, err := AsciiArt(option, userInput)
		if err != nil {
			handle500(w, r)
			return
		}

		escapedASCII := fmt.Sprintf("<pre>%s</pre>", html.EscapeString(asciiArt))
		resultHTML := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Ascii-Art Result</title>
			<link rel="stylesheet" href="/static/style-res.css">
		</head>
		<body>
			<div class="res-container">
				<img src="/static/assets/reslogo.svg" alt="Logo" width="500" height="200">
				<div class="ascii-container">
					%s
				</div>
				<a href="/"> <input type="image" src="/static/assets/backbutt.svg" height="55px"/> </a>
			</div>
		</body>
		</html>
		`, escapedASCII)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(resultHTML))
	})

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "static/404.html")
}

func handle400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	http.ServeFile(w, r, "static/400.html")
}

func handle500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	http.ServeFile(w, r, "static/500.html")
}
func AsciiArt(font string, input string) (string, error) {
	currhash, err := function.GetHash("standard.txt")
	if err != nil {
		return "", fmt.Errorf("failed to get file hash: %v", err)
	}
	validhash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	if currhash != validhash {
		return "", fmt.Errorf("invalid file hash, file is not the valid one")
	}

	ascii, err := os.ReadFile(font + ".txt")
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	lines := function.Split(string(ascii))
	asciiarr := make(map[rune][]string)
	chatr := 32
	for _, v := range lines {
		asciiarr[rune(chatr)] = v
		chatr++
	}

	var slice []string
	str := ""

	for _, x := range input {
		if x != '\n' {
			str += string(x)
		} else {
			if str != "" {
				slice = append(slice, str)
				str = ""
			}
			if x == '\n' {
				str = ""
			}
		}
	}
	if str != "" {
		slice = append(slice, str)
		str = ""
	}

	var result strings.Builder
	for i := 0; i < len(slice); i++ {
		if slice[i] != "\n" {
			lastElement := function.GetArr(slice[i], asciiarr)
			for i := 0; i < len(lastElement); i++ {
				result.WriteString(lastElement[i] + "\n")
			}
		} else {
			if i+1 < len(slice) {
				if slice[i] == "\n" && slice[i+1] == "\n" {
					fmt.Println("double new line")
					result.WriteString("\n")
				}
				if i == 0 {
					if slice[i] == "\n" {
						result.WriteString("\n")
					}
				}
			}
		}
	}
	return result.String(), nil
}
