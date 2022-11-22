package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string `mask:"hash" json:"username"`
	Password string `mask:"empty2"`

	// private field
	secret int
}

func main() {
	value := User{Username: "test", Password: "123123", secret: 4}

	t := reflect.TypeOf(value)
	
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			fmt.Printf(
				"%s: %s\n",
				t.Field(i).Name,
				t.Field(i).Tag.Get("mask"))
		}
	}
}
