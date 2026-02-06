package config

import (
	// driver import for sqlc, only imported for its side effects (DB communication)
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
)

type Config struct {
	DB *database.Queries
}

var AppConfig *Config

func Load() {
	_ = godotenv.Load()

	AppConfig = &Config{
		DB: getEnv("DB_URL"),
	}
}

// Refactor this later when we will need to fill more config fields
func getEnv(dbString string) *database.Queries {
	dbURL := os.Getenv(dbString)
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening databse: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("error connection to the db: %v", err)
	}

	dbQueries := database.New(dbConn)

	return dbQueries
}
