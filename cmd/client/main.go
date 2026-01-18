package client

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/nickemp1996/famchat/internal/database"
)

func main() {
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	dbQueries := database.New(db)

}
