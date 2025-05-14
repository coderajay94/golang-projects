package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	koanf "github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func main() {
	fmt.Println("Hello service to read the configs")

	//load env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("got error", err)
	}

	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")

	fmt.Println("Port", port)
	fmt.Println("dbUser", dbUser)
	fmt.Println("dbPassword", dbPassword)

	// Load the JSON config file
	if err := k.Load(file.Provider("application.json"), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Access nested config values
	serverHost := k.String("server.host")
	serverPort := k.Int("server.port")
	dbUser2 := k.String("database.user")
	dbPass := k.String("database.password")

	// Print the values
	fmt.Println("Server Host:", serverHost)
	fmt.Println("Server Port:", serverPort)
	fmt.Println("Database User:", dbUser2)
	fmt.Println("Database Pass:", dbPass)

	
}
