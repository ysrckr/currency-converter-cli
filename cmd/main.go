package main

import (
	"fmt"
	"log"

	"github.com/ysrckr/currency-converter-cli/internal/config"
	"github.com/ysrckr/currency-converter-cli/internal/huh"
)

func main() {

	c := config.NewConfig("/Users/yasar/github/projects/currency-converter/cli/.env")

	secret := c.InfisicalClient.GetSecret("OPEN_EXCHANGE_API_URI", "dev")

	fmt.Println(secret)

	currencyForm := huh.NewCurrencyForm()
	currencyForm.CreateForm()

	err := currencyForm.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

}
