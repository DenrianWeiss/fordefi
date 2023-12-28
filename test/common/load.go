package common

import (
	"github.com/DenrianWeiss/fordefi"
	"os"
	"sync"
)

var initOnce sync.Once

var clientInstance *fordefi.Client

func GetClient() *fordefi.Client {
	initOnce.Do(func() {
		token, _ := os.LookupEnv("FORDEFI_TOKEN")
		ecdsaKey, _ := os.LookupEnv("FORDEFI_ECDSA_KEY")
		if token == "" || ecdsaKey == "" {
			panic("FORDEFI_TOKEN or FORDEFI_ECDSA_KEY is not set")
		}
		var err error
		clientInstance, err = fordefi.New(fordefi.ClientJson{
			Host:     "https://api.fordefi.com",
			Token:    token,
			EcdsaKey: ecdsaKey,
		})
		if err != nil {
			panic(err)
		}
	})
	return clientInstance
}
