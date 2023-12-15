package listener

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io"
	"log"
	"net/http"
	"sync"
)

const ForDefiPublicKey = "-----BEGIN PUBLIC KEY-----\n" +
	"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQJ0NeDYQqqeCvgDofFsgtgaxk+dx\n" +
	"ybi63YGJwHz8Ebx7YQrmwNWnW3bG65E8wGHqZECjuaK2GKHbZx1EV2ws9A==\n" +
	"-----END PUBLIC KEY-----"

var keyCacheOnce sync.Once
var pubKeyCache *ecdsa.PublicKey

type EventHandler func(eventType string, transactionId string, state string)

func GetPubKey() *ecdsa.PublicKey {
	keyCacheOnce.Do(func() {
		// First Decode PEM
		blk, _ := pem.Decode([]byte(ForDefiPublicKey))
		if blk.Type != "PUBLIC KEY" {
			panic("Invalid PEM")
		}
		// Then Parse PKCS8
		key, err := x509.ParsePKIXPublicKey(blk.Bytes)
		if err != nil {
			panic(err)
		}
		// Finally Assert
		if _, ok := key.(*ecdsa.PublicKey); !ok {
			panic("Invalid key type")
		}
		pubKeyCache = key.(*ecdsa.PublicKey)
	})
	return pubKeyCache
}

func CreateHook(handler EventHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleWebHook(w, r, handler)
	}
}

func HandleWebHook(w http.ResponseWriter, r *http.Request, handler EventHandler) {
	pubKey := GetPubKey()
	// Get signature
	signature := r.Header.Get("X-Signature")
	if signature == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Decode signature
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hash := sha256.Sum256(body)

	// Verify signature
	pass := ecdsa.VerifyASN1(pubKey, hash[:], decodedSignature)
	if !pass {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req := map[string]interface{}{}
	err = json.Unmarshal(body, &req)
	eventType := req["event_type"].(string)
	transactionId := req["transaction_id"].(string)
	state := req["state"].(string)
	go handler(eventType, transactionId, state)
	w.WriteHeader(http.StatusOK)
}
