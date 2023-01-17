package main

import (
	"net/http"

	"github.com/awcjack/getground-backend-assignment/application"
	"github.com/awcjack/getground-backend-assignment/config"
	"github.com/awcjack/getground-backend-assignment/infrastructure/datastore"
	interfaces "github.com/awcjack/getground-backend-assignment/interface"
	chiLogger "github.com/awcjack/getground-backend-assignment/interface/logger"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
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
	db, err := datastore.NewMYSQLConnection(config.Database)
	if err != nil {
		logger.Fatal("Not able to create new mysql connection", err)
	}
	repo := datastore.NewMySQLRepository(db, logger)

	// Start application from repository implementation
	logger.Info("Loading Application")
	app := application.NewApplication(repo, repo, logger)

	// Start http server based on application (possible to switch to grpc or other rpc server but haven't implemented in this test)
	logger.Info("Loading HTTP Server")
	httpServer := interfaces.NewHttpServer(app)

	// Setting path to router
	rootRouter := interfaces.HandlerWithOptions(httpServer, interfaces.ChiServerOptions{
		// add common middleware and logrus logger
		Middlewares: []interfaces.MiddlewareFunc{
			middleware.RequestID,
			middleware.RealIP,
			chiLogger.NewStructuredLogger(logrus.StandardLogger()),
			middleware.Recoverer,
			middleware.NoCache,
		},
	})

	// Start HTTP server
	logger.Info("Starting Server")
	err = http.ListenAndServe(":"+config.Server.Port, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
