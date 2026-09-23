package main

import (
	"fmt"
	"log"
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

}
