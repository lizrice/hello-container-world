package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/page.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func main() {
	port := os.Getenv("WEB_SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	fmt.Printf("Serving on port %s\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("ListenAndServe exited with error %v\n", err)
	}
}
