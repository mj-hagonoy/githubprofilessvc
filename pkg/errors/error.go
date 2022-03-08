package errors

import (
	"fmt"
	"log"
	"net/http"
)

var (
	HttpNotFoundError = &HttpError{StatusCode: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound)}
	MaxLenghtError    = func(expected int, actual int) error {
		return fmt.Errorf("input error: expected %d, got %d", expected, actual)
	}
)

type HttpError struct {
	StatusCode int
	Message    string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error: statuscode=%d, %s", e.StatusCode, e.Message)
}

var (
	errors = make(chan error, 1)
)

func Send(err error) {
	errors <- err
}

func Run() {
	go func() {
		for err := range errors {
			log.Printf("ERROR: %v\n", err)
		}
	}()
}

func Stop() {
	close(errors)
}
