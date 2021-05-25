package main

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func bcrypt_hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {

	fmt.Print("Enter password: ")

	password, err := term.ReadPassword(syscall.Stdin)

	fmt.Println()

	if err != nil {
		log.Fatalln(err)
	}

	hashed, err := bcrypt_hash(string(password))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(hashed)
}
