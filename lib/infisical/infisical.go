package infisical

import (
	"log"
	"os"

	infisical "github.com/infisical/go-sdk"
)

type Client interface {
	Auth() infisical.AuthInterface
	Folders() infisical.FoldersInterface
	Secrets() infisical.SecretsInterface
}

type Infisical struct {
	client Client
}

func NewClient() *Infisical {
	i := &Infisical{
		client: infisical.NewInfisicalClient(infisical.Config{
			SiteUrl: "https://app.infisical.com",
		}),
	}

	return i
}

func (i *Infisical) Login(client_id, client_secret string) {

	_, err := i.client.Auth().UniversalAuthLogin(client_id, client_secret)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
