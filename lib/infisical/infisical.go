package infisical

import (
	"log"
	"os"

	infisical "github.com/infisical/go-sdk"
)

type SecretStore interface {
	Login(string, string)
	GetSecret(string, string) string
}

type Infisical struct {
	client    infisical.InfisicalClientInterface
	projectId string
}

func NewInfisicalClient(projectId string) SecretStore {
	i := &Infisical{
		client: infisical.NewInfisicalClient(infisical.Config{
			SiteUrl: "https://app.infisical.com",
		}),
		projectId: projectId,
	}

	return i
}

func (i Infisical) Login(client_id, client_secret string) {

	_, err := i.client.Auth().UniversalAuthLogin(client_id, client_secret)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func (i Infisical) GetSecret(secretKey, environment string) string {
	secrets, err := i.client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		SecretKey:   secretKey,
		Environment: environment,
		ProjectID:   i.projectId,
	})
	if err != nil {
		log.Fatalf("cannot get secret of %s. the error is: %v", secretKey, err)
	}

	return secrets.SecretValue
}
