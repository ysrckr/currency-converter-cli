package huh

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
)

// type Currency struct {
// 	USD, EUR, TRY, THB string
// }

type CurrencyForm struct {
	BaseCurrency    string
	TargetCurrency  string
	AmountToConvert string
	Form            *huh.Form
}

func NewCurrencyForm() *CurrencyForm {

	cf := &CurrencyForm{
		BaseCurrency:    "",
		TargetCurrency:  "",
		AmountToConvert: "",
	}

	return cf
}

func (cf *CurrencyForm) CreateForm() {
	cf.Form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your base currency").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("TRY", "TRY"),
					huh.NewOption("THB", "THB"),
				).
				Value(&cf.BaseCurrency),
		),
		huh.NewGroup(huh.NewSelect[string]().
			Title("Choose your target currency").
			Options(
				huh.NewOption("USD", "USD"),
				huh.NewOption("EUR", "EUR"),
				huh.NewOption("TRY", "TRY"),
				huh.NewOption("THB", "THB"),
			).
			Value(&cf.TargetCurrency).Validate(func(s string) error {
			if s == cf.BaseCurrency {
				return errors.New("base currency and target currency cannot be the same")
			}
			return nil
		}),
		),
		huh.NewGroup(
			huh.NewInput().Title("Amount to convert").Description("Please write as 10000.00").Value(&cf.AmountToConvert).Validate(func(s string) error {
				_, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return errors.New("please formatting guide lines for the number")
				}

				return nil
			}),
		),
	)
}
