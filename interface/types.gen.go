// Package interfaces provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package interfaces

// AddTable defines model for AddTable.
type AddTable struct {
	Capacity int64 `json:"capacity"`
	Id       int   `json:"id"`
}

// ArrivedGuest defines model for ArrivedGuest.
type ArrivedGuest struct {
	AccompanyingGuests int    `json:"accompanying_guests"`
	Name               string `json:"name"`
	TimeArrived        string `json:"time_arrived"`
}

// ArrivedGuests defines model for ArrivedGuests.
type ArrivedGuests struct {
	Guests []ArrivedGuest `json:"guests"`
}

// EmptySeats defines model for EmptySeats.
type EmptySeats struct {
	SeatsEmpty int `json:"seats_empty"`
}

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// Guest defines model for Guest.
type Guest struct {
	AccompanyingGuests int    `json:"accompanying_guests"`
	Name               string `json:"name"`
	Table              int    `json:"table"`
}

// GuestList defines model for GuestList.
type GuestList struct {
	Guests []Guest `json:"guests"`
}

// GuestName defines model for GuestName.
type GuestName struct {
	Name string `json:"name"`
}

// GuestNameParam defines model for GuestNameParam.
type GuestNameParam = string

// AddGuestJSONBody defines parameters for AddGuest.
type AddGuestJSONBody struct {
	AccompanyingGuests int `json:"accompanying_guests"`
	Table              int `json:"table"`
}

// CheckInGuestJSONBody defines parameters for CheckInGuest.
type CheckInGuestJSONBody struct {
	AccompanyingGuests int `json:"accompanying_guests"`
}

// AddTableJSONBody defines parameters for AddTable.
type AddTableJSONBody struct {
	Capacity int64 `json:"capacity"`
}

// AddGuestJSONRequestBody defines body for AddGuest for application/json ContentType.
type AddGuestJSONRequestBody AddGuestJSONBody

// CheckInGuestJSONRequestBody defines body for CheckInGuest for application/json ContentType.
type CheckInGuestJSONRequestBody CheckInGuestJSONBody

// AddTableJSONRequestBody defines body for AddTable for application/json ContentType.
type AddTableJSONRequestBody AddTableJSONBody
