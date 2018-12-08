package main

import (
	"log"
	"math/rand"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/rate-engineering/dai-ly/go/zap"

	_ "github.com/gemnasium/migrate/driver/postgres"
	"github.com/gemnasium/migrate/migrate" //migration database driver
	"github.com/gocraft/dbr"
	"github.com/joho/godotenv"
	"github.com/rate-engineering/dai-ly/go/database"
	"github.com/rate-engineering/dai-ly/go/server"
)

// ENVIRONMENTS
const (
	EnvPGHost     = "PGHOST"
	EnvPGUser     = "PGUSER"
	EnvPGPassword = "PGPASSWORD"
	EnvPGDatabase = "PGDATABASE"
	EnvPGPort     = "PGPORT"
)

// SessionStore encodes and decodes session data stored in signed cookies
// Can rotate keys in the future by appending new pair of hash and block key

// Structurinag Applications in Go reference https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091#.j1lcytmzz
func main() {
	// This is for generating random referral and token
	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	logger, err := zap.New()
	if err != nil {
		log.Fatal(err)
	}

	var s *server.Server
	s = server.New()
	//
	// // TODO main here should initialise more structs and fill up interface field with actual structs
	// // TODO Can our third party http utility packages work with http.Server directly?
	// // TODO separate out db init to main.go
	pgPort := os.Getenv(EnvPGPort)
	if pgPort == "" {
		pgPort = "5432"
	}

	pgURL := "postgres://" + url.QueryEscape(os.Getenv(EnvPGUser)) + ":" + url.QueryEscape(os.Getenv(EnvPGPassword)) + "@" + os.Getenv(EnvPGHost) + ":" + os.Getenv(EnvPGPort) + "/" + os.Getenv(EnvPGDatabase) + "?sslmode=disable"

	db, err := dbr.Open("postgres", pgURL, nil)
	if err != nil {
		logger.Fatal(err)
	}

	// Application dialect for timezone
	// This is critical to allow EncodeTime() to consider timezone
	// The dbr default driver will pass to db a timestamp in UTC but without timezone
	db.Dialect = database.PGTimezone{}

	// use synchronous versions of migration functions
	// Graceful mode doesnt work. Does not respond to multiple ^c
	migrate.NonGraceful()
	wd, err := os.Getwd()
	if err != nil {
		logger.Fatal()
	}
	migrationsPath := path.Join(wd, "./database/scripts/migrations")
	allErrors, ok := migrate.UpSync(pgURL, migrationsPath)
	if !ok {
		logger.Info(allErrors)
		logger.Info(pgURL)

	}

	s.API = &server.API{
		Store:  &database.Store{Connection: db, Logger: logger},
		Logger: logger,
	}

	err = s.SetRouter()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Fatal(s.Start())
}
