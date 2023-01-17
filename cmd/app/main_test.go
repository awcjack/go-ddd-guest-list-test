//go:build integration
// +build integration

// Integration test global setting

package main_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/awcjack/getground-backend-assignment/application"
	"github.com/awcjack/getground-backend-assignment/config"
	"github.com/awcjack/getground-backend-assignment/infrastructure/datastore"
	"github.com/awcjack/getground-backend-assignment/infrastructure/datastore/testdb"
	interfaces "github.com/awcjack/getground-backend-assignment/interface"
	chiLogger "github.com/awcjack/getground-backend-assignment/interface/logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var rootRouter http.Handler
var db *sqlx.DB

func TestMain(m *testing.M) {
	os.Setenv("DB_USERNAME", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOSTNAME", "mysql")
	os.Setenv("DB_NAME", "getground")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("PORT", "3000")

	// Init logrus as logger
	logger := logrus.New()

	// Loading config from env
	config, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Not able to load config", err)
	}

	// set logger level from config
	logLevel, err := logrus.ParseLevel(config.Application.LogLevel)
	if err != nil {
		logger.Errorf("Not able to parse log level %v, changing to default level info level", err)
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	// Start MySQL connection
	// Create repository implementation (possible to switch to other datastore implementation)
	// repo := datastore.NewMemoryRepository(logger) // change repo to this line able to switch the repository implementation to memory storage
	logger.Info("Loading datastore Reopsitory")
	db, err = datastore.NewMYSQLConnection(config.Database)
	if err != nil {
		logger.Fatal("Not able to create new mysql connection", err)
	}

	// Clean guests and tables table to make sure the database is clean when start
	err = testdb.CleanGuests(db)
	if err != nil {
		logger.Fatalf("error truncating test database guests table: %v", err)
	}
	err = testdb.CleanTables(db)
	if err != nil {
		logger.Fatalf("error truncating test database tables table: %v", err)
	}

	repo := datastore.NewMySQLRepository(db, logger) // Start application from repository implementation
	logger.Info("Loading Application")
	app := application.NewApplication(repo, repo, logger)

	// Start http server based on application (possible to switch to grpc or other rpc server but haven't implemented in this test)
	logger.Info("Loading HTTP Server")
	httpServer := interfaces.NewHttpServer(app)

	// Setting path to router
	rootRouter = interfaces.HandlerWithOptions(httpServer, interfaces.ChiServerOptions{
		// add common middleware and logrus logger
		Middlewares: []interfaces.MiddlewareFunc{
			middleware.RequestID,
			middleware.RealIP,
			chiLogger.NewStructuredLogger(logrus.StandardLogger()),
			middleware.Recoverer,
			middleware.NoCache,
		},
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}
