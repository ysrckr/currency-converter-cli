package main

import (
	"log"

	"github.com/ysrckr/currency-converter-cli/internal/config"
	"github.com/ysrckr/currency-converter-cli/internal/huh"
)

var Config *config.Config

func main() {

	Config = config.NewConfig()

	currencyForm := huh.NewCurrencyForm()
	currencyForm.CreateForm()

	err := currencyForm.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

}
