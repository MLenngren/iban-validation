package main

import (
	"log"

	"github.com/mlenngren/iban-validator/ibanService/routers"
)

func main() {
	r := routers.Setup()
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}
