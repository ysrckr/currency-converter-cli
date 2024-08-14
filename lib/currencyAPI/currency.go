package currencyapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/ysrckr/currency-converter-cli/internal/config"
)

type Currency struct {
	Timestamp uint64             `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

var conf = config.NewConfig("/Users/yasar/github/projects/currency-converter/cli/.env")

func GetCurrency(base, target string) map[string]float64 {
	api_uri := conf.InfisicalClient.GetSecret("OPEN_EXCHANGE_API_URI", "dev")
	api_app_id := conf.InfisicalClient.GetSecret("OPEN_EXCHANGE_API_APP_ID", "dev")

	api_url, err := url.Parse(fmt.Sprintf("%s/latest.json?app_id=%s&base=%s", api_uri, api_app_id, base))
	if err != nil {
		log.Println(err)
	}

	currency := Currency{}

	resp, err := http.Get(api_url.String())
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &currency)
	if err != nil {
		log.Println(err)
	}

	if target == "" {
		return currency.Rates
	}

	result := map[string]float64{
		target: currency.Rates[target],
	}

	return result

}
