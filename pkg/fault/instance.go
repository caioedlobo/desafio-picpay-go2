package fault

import (
	"encoding/json"
	"errors"
	"net/http"
)

// NewHTTPError receives an error and writes it to the response writer
// It sets the content type to application/json and writes the error
// If the error is not a Fault, it writes a new InternalServerError
func NewHTTPError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var fault *Fault
	if errors.As(err, &fault) {
		w.WriteHeader(fault.GetCode())
		_ = json.NewEncoder(w).Encode(fault)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(
		New(
			"an unexpected error occurred",
			WithHTTPCode(http.StatusInternalServerError),
			//WithTag(InternalServerError),
			WithError(err),
		),
	)
}

func NewBadRequest(msg string) *Fault {
	return New(msg, WithHTTPCode(http.StatusBadRequest))
}

func NewUnauthorized(msg string) *Fault {
	return New(msg, WithHTTPCode(http.StatusUnauthorized))
}
func NewInternalServerError(message string) *Fault {
	return New(
		message,
		WithHTTPCode(http.StatusInternalServerError),
	)
}

func NewConflict(message string) *Fault {
	return New(
		message,
		WithHTTPCode(http.StatusConflict),
	)
}
