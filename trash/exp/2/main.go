package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateOrg()
	fmt.Println(err)
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		fmt.Printf("integer: %d, string: %s, any value: %v\n", 123, "Hello , world!", "3.14, Hello, 12, %$#@")
		return fmt.Errorf("create user: %w", err)
	}
	//...
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	//...
	return nil
}
