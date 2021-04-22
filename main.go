package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env file")
	}

	fmt.Println("Output:", os.Getenv("DB_CONNECTION_STRING"))
}
