package main

import (
	"log"

	"github.com/ysrckr/currency-converter-cli/internal/config"
	"github.com/ysrckr/currency-converter-cli/internal/huh"
)

var Config *config.Config

func main() {

	Config = config.NewConfig()

	err := huh.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

}
