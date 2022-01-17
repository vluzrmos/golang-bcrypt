package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	var cost int

	flag.IntVar(&cost, "cost", bcrypt.DefaultCost, "BCrypt password cost.")
	flag.IntVar(&cost, "c", bcrypt.DefaultCost, "BCrypt password cost.")
	flag.Parse()

	var password []byte

	info, err := os.Stdin.Stat()

	if err != nil {
		log.Fatalln(err)
	}

	if (info.Mode() & os.ModeCharDevice) == 0 {
		p, _ := ioutil.ReadAll(os.Stdin)
		password = p
	} else {
		fmt.Print("Enter password: ")

		p, err := term.ReadPassword(syscall.Stdin)

		if err != nil {
			log.Fatalln(err)
		}

		password = p

		fmt.Println()
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
