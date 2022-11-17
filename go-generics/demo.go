package main

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"reflect"
)

func main() {
	user := User{
		ID: 2,

		Credentials: &Credentials{
			Email: "test@test.com",
			Username: "test",
			Password: "123123",
		},
		Anything: "asrt",
	}

	maskedUser := maskStruct(&user)

	json.NewEncoder(os.Stdout).Encode(user)
	json.NewEncoder(os.Stdout).Encode(maskedUser)
}

type User struct {
	ID	int

	Credentials *Credentials
	Anything any    `mask:"empty"`

}

type Credentials struct {
	Email 	 string `mask:"hash"`
	Username string `mask:"hash"`
	Password string `mask:"empty"`

}

func maskStruct(value any) any {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	newV := reflect.New(t).Elem()

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			switch t.Field(i).Tag.Get("mask") {
			case "hash":
				newV.Field(i).Set(maskHash(v.Field(i)))
			case "empty":
				newV.Field(i).Set(reflect.Zero(t.Field(i).Type))
			default:
				switch {
				case t.Field(i).Type.Kind() == reflect.Struct:
					newV.Field(i).Set(reflect.ValueOf(maskStruct(v.Field(i).Interface())))
				case t.Field(i).Type.Kind() == reflect.Pointer && t.Field(i).Type.Elem().Kind() == reflect.Struct:
					newV.Field(i).Elem().Set(reflect.ValueOf(maskStruct(v.Field(i).Interface())))
				default:
					newV.Field(i).Set(v.Field(i))
				}
			}
		}
	}

	return newV.Interface()
}

func maskHash(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.String {
		panic("invalid field type to hash")
	}

	dst := base64.StdEncoding.EncodeToString([]byte(v.String()))
	return reflect.ValueOf(dst)
}
