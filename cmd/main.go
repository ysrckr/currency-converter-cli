package main

import (
	"fmt"
	"log"
	"strconv"

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

	amount, err := strconv.ParseFloat(currencyForm.AmountToConvert, 64)
	if err != nil {
		log.Fatalf("error")
	}

	result := amount * currency[currencyForm.TargetCurrency] / currency[currencyForm.BaseCurrency]

	fmt.Printf("Result of conversion is %s %s = %.2f %s", currencyForm.AmountToConvert, currencyForm.BaseCurrency, result, currencyForm.TargetCurrency)

}
