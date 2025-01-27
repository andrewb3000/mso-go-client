package tests

import (
	"testing"

	"github.com/andrewb3000/mso-go-client/client"
)

func TestClientAuthenticate(t *testing.T) {

	client := GetTestClient()
	err := client.Authenticate()
	if err != nil {
		t.Error(err)
	}

	if client.AuthToken.Token == "{}" {
		t.Error("Token is empty")
	}
	t.Error("all wrong")
}

func GetTestClient() *client.Client {
	return client.GetClient("https://173.36.219.193", "admin", client.Password("ins3965!ins3965!"), client.Insecure(true))

}
