package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Starting Peril server...")

	url := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(url)
	if err != nil {
		log.Fatalln("Error creating connection:", err)
	}
	defer connection.Close()

	fmt.Println("Connection successful!")

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalln("Error creating connection channel:", err)
	}
}
