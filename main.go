package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	var cost int

	flag.IntVar(&cost, "cost", bcrypt.DefaultCost, "BCrypt password cost.")
	flag.IntVar(&cost, "c", bcrypt.DefaultCost, "BCrypt password cost.")
	flag.Parse()

	fmt.Print("Enter password: ")

	password, err := term.ReadPassword(syscall.Stdin)

	fmt.Println()

	if err != nil {
		log.Fatalln(err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		log.Fatalln(err)
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(hashed))
}
