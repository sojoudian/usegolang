package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	HostName string
	Age      int
	Vote     UserMeta
}
type UserMeta struct {
	Option_a string
	Option_b string
}

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)

	userHostname := User{
		HostName: hostname,
		Age:      111,
		Vote: UserMeta{
			Option_a: "cat",
			Option_b: "dog",
		},
	}

	err = t.Execute(os.Stdout, userHostname)
	if err != nil {
		panic(err)
	}
}
