package httputil

import (
	"desafio-picpay-go2/pkg/strutil"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const maxRequestBodyBytes = 1_048_576

var (
	ErrUnknownRequestBodyKey = errors.New("request body contains unknown key")
	ErrEmptyRequestBody      = errors.New("body cannot be empty")
)

func ReadRequestBody(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxRequestBodyBytes))

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			// JSON syntax error in the request body
			// Offset is the exact byte where the error occurred
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			// JSON value and struct type do not match
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			// io.EOF (End of File) indicates that there are no more bytes left to read
			return ErrEmptyRequestBody
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			return ErrUnknownRequestBodyKey
		case errors.As(err, &invalidUnmarshalError):
			// Received a non-nil pointer into Decode()
			panic(err)
		default:
			return err
		}
	}

	// Calling decode again to check if there's more data after the JSON object
	// This will return an io.EOF error, indicating that the client sent more data
	err = d.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

// WriteJSONResponse writes a JSON response with the specified HTTP status code and the provided data.
// The data to be encoded as JSON should be passed as the 'dst' parameter.
func WriteJSON(w http.ResponseWriter, code int, data strutil.Envelope) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}
