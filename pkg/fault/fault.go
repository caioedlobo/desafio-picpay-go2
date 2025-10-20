package fault

import (
	"fmt"
	"net/http"
)

type Fault struct {
	Code    int
	Err     error
	Message string
}

func New(msg string, options ...func(*Fault)) *Fault {
	fault := Fault{
		Code:    http.StatusBadRequest,
		Err:     nil,
		Message: msg,
	}

	for _, fn := range options {
		fn(&fault)
	}
	return &fault
}

// WithHTTPCode sets the HTTP code for the fault
func WithHTTPCode(code int) func(*Fault) {
	return func(f *Fault) {
		f.Code = code
	}
}

// WithError sets the error for the fault
func WithError(err error) func(*Fault) {
	return func(f *Fault) {
		if err == nil {
			return
		}
		f.Err = err
	}
}

// GetCode returns the HTTP code for the fault
func (f *Fault) GetCode() int {
	return f.Code
}

func (f *Fault) Error() string {
	if f.Err != nil {
		return fmt.Sprintf("%s (caused by: %v)", f.Message, f.Err)
	}
	return fmt.Sprintf(" %s", f.Message)
}
