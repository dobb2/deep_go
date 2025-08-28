package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	var resultErr string
	if len(e.errors) == 0 {
		return resultErr
	}

	resultErr += fmt.Sprintf("%d errors occured:\n", len(e.errors))

	if len(e.errors) == 1 {
		return e.errors[0].Error()
	}

	for i, err := range e.errors {
		if i == len(e.errors)-1 {
			resultErr += fmt.Sprintf("\t* %s\n", err.Error())
		} else {
			resultErr += fmt.Sprintf("\t* %s", err.Error())
		}

	}

	return resultErr
}

func Append(err error, errs ...error) *MultiError {
	multiError := MultiError{
		errors: make([]error, 0, len(errs)+1),
	}
	if err != nil {
		multiError.errors = append(multiError.errors, err)
	}

	for _, error := range errs {
		multiError.errors = append(multiError.errors, error)
	}

	return &multiError
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
