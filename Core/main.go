package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	var host, user, password, database string = "db", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, database),
	)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully created connection to database")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Xiaoxuan!")
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
