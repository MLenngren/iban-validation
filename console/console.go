package main

import (
	"log"
	"os"

	ibanValidator "github.com/mlenngren/iban-validator/ibanValidator"
)

func main() {

	if len(os.Args) == 1 {
		log.Print("No arguments, you must supply IBAN")
		os.Exit(1)
	}

	validated := ibanValidator.ValidateIban(os.Args[1])
	if validated == 1 {
		println("validated")
		os.Exit(0)
	}

	println("not validated")
	os.Exit(0)
}
