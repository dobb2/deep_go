package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize(person Person) string {
	var propertiesStr string
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	for idx := 0; idx < t.NumField(); idx++ {
		fieldTag := t.Field(idx).Tag
		fieldInterface := v.Field(idx).Interface()

		var omitempty bool
		propertiesTag, _ := fieldTag.Lookup("properties")

		if strings.Contains(propertiesTag, "omitempty") {
			omitempty = true
		}

		propertiesNameField := strings.SplitN(propertiesTag, ",", 2)[0]

		switch value := fieldInterface.(type) {
		case string:
			if omitempty && value == "" {
				continue
			}
			propertiesStr += fmt.Sprintf("%s=%s", propertiesNameField, value)
		case int:
			if omitempty && value == 0 {
				continue
			}
			propertiesStr += fmt.Sprintf("%s=%d", propertiesNameField, value)
		case bool:
			if omitempty && !value {
				continue
			}
			propertiesStr += fmt.Sprintf("%s=%t", propertiesNameField, value)
		}
		if idx != t.NumField()-1 {
			propertiesStr += fmt.Sprint("\n")
		}
	}

	return propertiesStr
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
