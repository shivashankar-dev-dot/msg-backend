package main

import (
	"fmt"
	"log"
	database "msg/internal/database"
	"net/http"
	"os"

	"github.com/subosito/gotenv"
)

func main() {

	fmt.Println("Hello world")

	err := gotenv.Load()

	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Print(os.Getenv("DATABASE_URL"))

	db, err := database.Connect()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
