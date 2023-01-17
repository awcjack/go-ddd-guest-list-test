package application

import (
	"github.com/awcjack/getground-backend-assignment/application/command"
	"github.com/awcjack/getground-backend-assignment/application/query"
	"github.com/awcjack/getground-backend-assignment/domain/guest"
	"github.com/awcjack/getground-backend-assignment/domain/table"
)

// Query handlers (querying data via repository)
type Queries struct {
	ListGuest      *query.ListGuestHandler
	ListArrived    *query.ListArrivedGuestHandler
	CountEmptySeat *query.CountEmptySeatHandler
	GetTable       *query.GetTableHandler
}

// Command handlers (updating data via repository)
type Commands struct {
	AddGuest      *command.AddGuestHandler
	CheckInGuest  *command.CheckInGuestHandler
	CheckOutGuest *command.CheckOutGuestHandler
	AddTable      *command.AddTableHandler
}

// expose all feature from the application
type Application struct {
	Queries  Queries
	Commands Commands
}

// logger interface that pass the handlers
type logger interface {
	Panicf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

// Application constructor
func NewApplication(guestRepo guest.Repository, tableRepo table.Repository, logger logger) Application {
	listGuestHandler := query.NewListGuestHandler(guestRepo, logger)
	listArrivedHandler := query.NewListArrivedGuestHandler(guestRepo, logger)
	countEmptySeatHandler := query.NewCountEmptySeatHandler(tableRepo, logger)
	getTableHandler := query.NewGetTableHandler(tableRepo, logger)
	addGuestHandler := command.NewAddGuestHandler(guestRepo, logger)
	checkInGuestHandler := command.NewCheckInGuestHandler(guestRepo, logger)
	checkOutGUestHandler := command.NewCheckOutGuestHandler(guestRepo, logger)
	addTableHandler := command.NewAddTableHandler(tableRepo, logger)

	return Application{
		Queries: Queries{
			ListGuest:      listGuestHandler,
			ListArrived:    listArrivedHandler,
			CountEmptySeat: countEmptySeatHandler,
			GetTable:       getTableHandler,
		},
		Commands: Commands{
			AddGuest:      addGuestHandler,
			CheckInGuest:  checkInGuestHandler,
			CheckOutGuest: checkOutGUestHandler,
			AddTable:      addTableHandler,
		},
	}
}
