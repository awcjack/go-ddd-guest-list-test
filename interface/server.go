//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./codegen.config.yaml ../docs/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./typegen.config.yaml ../docs/openapi.yaml

package interfaces

import (
	"net/http"

	"github.com/awcjack/getground-backend-assignment/application"
	"github.com/awcjack/getground-backend-assignment/domain/guest"

	"github.com/go-chi/render"
)

// HTTP server that using the application to handle all operation
// Implementing ServerInterfaceWrapper for using the chi server interface generated from oapi-codegen
type HttpServer struct {
	app application.Application
}

// HttpServer constructor
func NewHttpServer(app application.Application) HttpServer {
	return HttpServer{
		app: app,
	}
}

// domain guest to http server format for response
func guestModelToResponse(models []guest.Guest) GuestList {
	var guestList GuestList
	var guests []Guest

	for _, g := range models {
		guests = append(guests, Guest{
			Name:               g.Name(),
			Table:              g.Table().Id(),
			AccompanyingGuests: g.GuestNumber(),
		})
	}

	guestList.Guests = guests
	return guestList
}

// Display guestlist
// (GET /guest_list)
func (h HttpServer) ListGuest(w http.ResponseWriter, r *http.Request) {
	// Call application handler to retrieve guest list
	guestsModel, err := h.app.Queries.ListGuest.Handle(r.Context())
	if err != nil {
		BadRequest(err, w, r)
		return
	}

	// convert guest format for response
	guests := guestModelToResponse(guestsModel)
	render.JSON(w, r, guests)
}

// Add guest to guestlist
// (POST /guest_list/{name})
func (h HttpServer) AddGuest(w http.ResponseWriter, r *http.Request, name GuestNameParam) {
	addGuest := AddGuestJSONBody{}
	// get guest info from request body
	if err := render.Decode(r, &addGuest); err != nil {
		BadRequest(err, w, r)
		return
	}

	// Call application handler to get table
	table, err := h.app.Queries.GetTable.Handle(r.Context(), addGuest.Table)
	if err != nil {
		BadRequest(err, w, r)
		return
	}

	// Call application handler to add guest
	if err := h.app.Commands.AddGuest.Handle(r.Context(), guest.NewGuest(name, addGuest.AccompanyingGuests+1, table)); err != nil {
		BadRequest(err, w, r)
		return
	}

	render.JSON(w, r, GuestName{
		Name: name,
	})
}

// domain guest to http server format for response
func arrivedGuestModelToResponse(models []guest.Guest) ArrivedGuests {
	var arrivedGuests ArrivedGuests
	var guests []ArrivedGuest

	for _, g := range models {
		guests = append(guests, ArrivedGuest{
			Name:               g.Name(),
			AccompanyingGuests: g.ArrivedNumber() - 1,
			TimeArrived:        g.ArrivedTime().Format("2006-01-02 15:04:05"), // datetime mysql format
		})
	}
	arrivedGuests.Guests = guests

	return arrivedGuests
}

// list arrived guest
// (GET /guests)
func (h HttpServer) ListArrivedGuest(w http.ResponseWriter, r *http.Request) {
	// Call application handler to retrieve arrived guest list
	guestsModel, err := h.app.Queries.ListArrived.Handle(r.Context())
	if err != nil {
		BadRequest(err, w, r)
		return
	}

	// convert guest forat for server response
	guests := arrivedGuestModelToResponse(guestsModel)

	render.JSON(w, r, guests)
}

// guest checkout
// (DELETE /guests/{name})
func (h HttpServer) CheckOutGuest(w http.ResponseWriter, r *http.Request, name GuestNameParam) {
	// Call application handler to checkout guest
	err := h.app.Commands.CheckOutGuest.Handle(r.Context(), name)
	if err != nil {
		BadRequest(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// guest checkin
// (PUT /guests/{name})
func (h HttpServer) CheckInGuest(w http.ResponseWriter, r *http.Request, name GuestNameParam) {
	checkInGuest := CheckInGuestJSONBody{}
	// get guest info from request body
	if err := render.Decode(r, &checkInGuest); err != nil {
		BadRequest(err, w, r)
		return
	}

	// Call application handler to check in guest
	err := h.app.Commands.CheckInGuest.Handle(r.Context(), name, checkInGuest.AccompanyingGuests)
	if err != nil {
		BadRequest(err, w, r)
		return
	}

	render.JSON(w, r, GuestName{
		Name: name,
	})
}

// Count empty seat
// (GET /seats_empty)
func (h HttpServer) GetEmptySeat(w http.ResponseWriter, r *http.Request) {
	// Call application handler to count empty seat
	count := h.app.Queries.CountEmptySeat.Handle(r.Context())

	render.JSON(w, r, EmptySeats{
		SeatsEmpty: count,
	})
}

// Add table to the system
// (POST /tables)
func (h HttpServer) AddTable(w http.ResponseWriter, r *http.Request) {
	addTable := AddTableJSONBody{}
	// get table info from request body
	if err := render.Decode(r, &addTable); err != nil {
		BadRequest(err, w, r)
		return
	}

	// Call application handler to add table
	table, err := h.app.Commands.AddTable.Handle(r.Context(), int(addTable.Capacity))
	if err != nil {
		BadRequest(err, w, r)
		return
	}

	render.JSON(w, r, AddTable{
		Id:       table.Id(),
		Capacity: int64(table.Capacity()),
	})
}
