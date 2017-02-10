package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func getDatabase() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	fmt.Printf("Database config: %s\n", dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	return db
}

func incrementCount() int {
	var hits int
	row := db.QueryRow("UPDATE hits SET total = total + 1 RETURNING total;")
	err := row.Scan(&hits)
	if err != nil {
		fmt.Printf("Database increment error: %v\n", err)
		return -1
	}
	return hits
}

func handler(w http.ResponseWriter, r *http.Request) {
	hits := incrementCount()
	t, err := template.ParseFiles("templates/page.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, hits)
}

// Browsers typically send a request for favicon.ico, which we don't want to include in our page count.
func faviconHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	db = getDatabase()
	defer db.Close()

	port := os.Getenv("WEB_SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	fmt.Printf("Serving on port %s\n", port)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("ListenAndServe exited with error %v\n", err)
	}
}
