package main

import (
	"log"
	"math/rand"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/rate-engineering/dai-ly/go/ethereum"
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
	EnvPGHost        = "PGHOST"
	EnvPGUser        = "PGUSER"
	EnvPGPassword    = "PGPASSWORD"
	EnvPGDatabase    = "PGDATABASE"
	EnvPGPort        = "PGPORT"
	EnvTokenContract = "TOKEN_CONTRACT"
	EnvRPCClient     = "RPC_CLIENT"
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

	// set up ethereum client
	client, err := ethereum.NewClient(os.Getenv(EnvTokenContract), os.Getenv(EnvRPCClient))
	if err != nil {
		logger.Fatal()
	}

	s.API = &server.API{
		Store:          &database.Store{Connection: db, Logger: logger},
		Logger:         logger,
		EthereumClient: client,
	}

	// tx := daily.Transaction{
	// 	Sender:     "0x2b522cabe9950d1153c26c1b399b293caa99fcf9",
	// 	Receiver:   "0x3644b986b3f5ba3cb8d5627a22465942f8e06d09",
	// 	FeeAmount:  "2",
	// 	SendAmount: "100",
	// 	Nonce:      "0",
	// 	Signature:  "0xa1773b6e73877a02f4681c278fe0c3361de5c964b18e02eb95b8a3c82f1f78972cecabeb1679b31b57586e18a035e0c97044c1bac7d3ebf681df2a468c46798c00",
	// }
	// println("submitting")
	// submitted, err := s.API.EthereumClient.ExecuteTransaction(tx)
	// println(submitted)
	// if err != nil {
	// 	log.Print(err)
	// }

	err = s.SetRouter()
	if err != nil {
		logger.Fatal(err)
	}

	job, _ := scheduler.Every(10).Seconds().Run(s.API.FindAndSubmitQueuedTransaction)
	job.SkipWait <- true
	job, _ = scheduler.Every(5).Seconds().Run(s.API.IterateProcessingTransactions)
	job.SkipWait <- true

	go s.API.ListenEvents()

	logger.Fatal(s.Start())
}
