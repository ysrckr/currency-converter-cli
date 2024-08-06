package main

import (
	"log"

	"github.com/ysrckr/currency-converter-cli/internal/huh"
)

func main() {
	err := huh.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

}
