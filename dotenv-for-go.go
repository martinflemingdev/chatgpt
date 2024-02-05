package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()

	//err := godotenv.Load("path/to/your/.env")


    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Use environment variables
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")

    fmt.Printf("DB_HOST: %s\n", dbHost)
    fmt.Printf("DB_USER: %s\n", dbUser)
    fmt.Printf("DB_PASSWORD: %s\n", dbPassword)
}
