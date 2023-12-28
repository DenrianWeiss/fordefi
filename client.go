package fordefi

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const DefaultApiEndpoint = "https://api.fordefi.com"

type Client struct {
	Host     string
	Token    string
	EcdsaKey ecdsa.PrivateKey
}

type ClientJson struct {
	Host     string `json:"host,omitempty"`
	Token    string `json:"token,omitempty"`
	EcdsaKey string `json:"ecdsaKey,omitempty"` // Key in pem format
}

func New(clientConfig ClientJson) (*Client, error) {
	// Decode pem key
	// decode base 64
	rawBytes, err := base64.StdEncoding.DecodeString(clientConfig.EcdsaKey)
	if err != nil {
		return nil, err
	}
	decodedKey, err := x509.ParseECPrivateKey(rawBytes)
	if err != nil {
		return nil, err
	}
	// Init client
	client := &Client{
		Host:     clientConfig.Host,
		Token:    clientConfig.Token,
		EcdsaKey: *decodedKey,
	}
	if clientConfig.Host == "" {
		client.Host = DefaultApiEndpoint
	}
	return client, nil
}

func (c *Client) post(path string, req interface{}, noSignature ...bool) (resp []byte, err error) {
	timeStampMilli := time.Now().UnixMilli()
	var serBody []byte
	var reqHttp *http.Request
	if req != nil {
		serBody, err = json.Marshal(req)
	}
	if err != nil {
		return nil, err
	}
	// Build request
	reqHttp, err = http.NewRequest("POST", c.Host+path, bytes.NewBuffer(serBody))
	if err != nil {
		return nil, err
	}
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	reqHttp.Header.Set("X-Timestamp", fmt.Sprintf("%d", timeStampMilli))
	if len(noSignature) == 0 || !noSignature[0] {
		// Sign
		sigPayload := fmt.Sprintf("%d|%s|%s", timeStampMilli, path, serBody)
		sigDigest := sha256.Sum256([]byte(sigPayload))
		sig, err := c.EcdsaKey.Sign(rand.Reader, sigDigest[:], nil)
		if err != nil {
			return nil, err
		}
		reqHttp.Header.Set("X-Signature", base64.StdEncoding.EncodeToString(sig))
	}
	// Send request
	client := http.DefaultClient
	respHttp, err := client.Do(reqHttp)
	if err != nil {
		return nil, err
	}
	if respHttp.StatusCode == http.StatusNoContent {
		return nil, nil
	}
	resp, err = io.ReadAll(respHttp.Body)
	if err != nil {
		return nil, err
	}
	if respHttp.StatusCode > 300 || respHttp.StatusCode < 200 {
		return resp, fmt.Errorf("status code %d", respHttp.StatusCode)
	}
	return resp, nil
}

func (c *Client) get(path string) (resp []byte, err error) {
	reqHttp, err := http.NewRequest("GET", c.Host+path, nil)
	if err != nil {
		return nil, err
	}
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	// Send request
	client := http.DefaultClient
	respHttp, err := client.Do(reqHttp)
	if err != nil {
		return nil, err
	}
	resp, err = io.ReadAll(respHttp.Body)
	if err != nil {
		return nil, err
	}
	if respHttp.StatusCode > 300 || respHttp.StatusCode < 200 {
		return resp, fmt.Errorf("status code %d", respHttp.StatusCode)
	}
	return resp, nil
}
