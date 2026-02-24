package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	// driver import for sqlc, only imported for its side effects (DB communication)
	_ "github.com/lib/pq"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
)

type Config struct {
	DB *database.Queries
	// Needed in case we want to use transactions in our code
	SQLDB     *sql.DB
	JWTSecret string
}

var AppConfig *Config

func Load() {
	_ = godotenv.Load()

	sqldb, queries, secret := getEnv("DB_URL", "JWT_SECRET")
	AppConfig = &Config{
		DB:        queries,
		SQLDB:     sqldb,
		JWTSecret: secret,
	}
}

func getEnv(dbString, jwtSecret string) (*sql.DB, *database.Queries, string) {
	dbURL := os.Getenv(dbString)
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	secret := os.Getenv(jwtSecret)
	if secret == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("error connecting to the db: %v", err)
	}

	dbQueries := database.New(dbConn)

	return dbConn, dbQueries, secret
}
