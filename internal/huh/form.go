package huh

import (
	"github.com/charmbracelet/huh"
)

// type Currency struct {
// 	USD, EUR, TRY, THB string
// }

type CurrencyForm struct {
	BaseCurrency   string
	TargetCurrency string
	Form           *huh.Form
}

func NewCurrencyForm() *CurrencyForm {

	cf := &CurrencyForm{
		BaseCurrency:   "",
		TargetCurrency: "",
	}

	return cf
}

func (cf *CurrencyForm) CreateForm() {
	cf.Form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your base currency").
				Options(
					huh.NewOption("USD", "usd"),
					huh.NewOption("EUR", "eur"),
					huh.NewOption("TRY", "try"),
					huh.NewOption("THB", "thb"),
				).
				Value(&cf.BaseCurrency),

			huh.NewSelect[string]().
				Title("Choose your target currency").
				Options(
					huh.NewOption("USD", "usd"),
					huh.NewOption("EUR", "eur"),
					huh.NewOption("TRY", "try"),
					huh.NewOption("THB", "thb"),
				).
				Value(&cf.TargetCurrency),
		),
	)
}
