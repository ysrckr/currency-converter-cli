package main

import (
	"fmt"
	"log"

	"github.com/ysrckr/currency-converter-cli/internal/huh"
	currencyapi "github.com/ysrckr/currency-converter-cli/lib/currencyAPI"
)

func main() {

	currencyForm := huh.NewCurrencyForm()
	currencyForm.CreateForm()

	err := currencyForm.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

	currency := currencyapi.GetCurrency(currencyForm.BaseCurrency, currencyForm.TargetCurrency)

	currencyForm.SetAmount()

	result := currencyForm.Convert(currency)

	fmt.Println(result)

}
