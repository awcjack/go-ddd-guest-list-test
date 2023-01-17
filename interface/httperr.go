package interfaces

import (
	"net/http"

	"github.com/awcjack/getground-backend-assignment/interface/logger"
	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	httpStatus int
}

// Bad request response
func BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, w, r, "Bad request", http.StatusBadRequest)
}

// Generic error response
func GeneralHTTPRespondWithError(err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, w, r, err.Error(), http.StatusBadRequest)
}

func httpRespondWithError(err error, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	logger.GetLogEntry(r).WithError(err).Warn(logMSg)
	resp := ErrorResponse{err.Error(), status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

// chi render function to render the error response
func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.httpStatus)
	return nil
}
