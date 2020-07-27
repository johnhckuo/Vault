package vault

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

type Vault struct {
	ServiceName   string
	token         string
	vaultAddr     string
	client        *api.Client
	ParticipantId string
	ENV           string
}

// hard-codede vaule for now, To be changed
const envTokenName = "TOKEN"
const envAddrName = "VAULT_ADDR"
const serviceName = "send-service"
const participantId = "ibm01"
const envStage = "sandbox"

func InitializeVault() (*Vault, error) {

	var token, vaultAddr string
	var exists bool

	if token, exists = os.LookupEnv(envTokenName); !exists {
		return nil, errors.New("No Vault API token detected")
	}

	if vaultAddr, exists = os.LookupEnv(envAddrName); !exists {
		return nil, errors.New("No Vault API address detected")
	}

	var vaultClient = &Vault{
		token:         token,
		vaultAddr:     vaultAddr,
		ServiceName:   serviceName,
		ParticipantId: participantId,
		ENV:           envStage,
	}
	var err error

	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	config := &api.Config{
		Address:    vaultClient.vaultAddr,
		HttpClient: httpClient,
	}
	vaultClient.client, err = api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	vaultClient.client.SetToken(vaultClient.token)
	return vaultClient, nil
}
