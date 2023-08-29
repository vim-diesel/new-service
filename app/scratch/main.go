package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	pass, err := bcrypt.GenerateFromPassword([]byte("adminpass"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println(string(pass))

	pass, err = bcrypt.GenerateFromPassword([]byte("userpass"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println(string(pass))

	return nil
}
