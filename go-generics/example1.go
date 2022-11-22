package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	Password string
	Email    string

	// private field
	secret int
}

func main() {
	value := User{Username: "test",Email: "test@arst", Password: "123123", secret: 4}

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	for i := 0; i < t.NumField(); i++ {
		// Reflection panics if you access a value that is unexported
		if t.Field(i).IsExported() {
			fmt.Printf("%s: %v\n", t.Field(i).Name, v.Field(i).Interface())
		}
	}
}
