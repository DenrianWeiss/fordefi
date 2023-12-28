package fordefi

import (
	"encoding/json"
	"log"
)

type ListChainResponse struct {
	Chains []struct {
		ChainType            string `json:"chain_type"`
		UniqueId             string `json:"unique_id"`
		Name                 string `json:"name"`
		NativeCurrencySymbol string `json:"native_currency_symbol"`
		NativeCurrencyName   string `json:"native_currency_name"`
		BlockchainExplorer   struct {
			TransactionUrl       string `json:"transaction_url"`
			AddressUrl           string `json:"address_url"`
			RootUrl              string `json:"root_url"`
			TransactionFormatUrl string `json:"transaction_format_url"`
			AddressFormatUrl     string `json:"address_format_url"`
		} `json:"blockchain_explorer"`
		BaseDenom string `json:"base_denom"`
	} `json:"chains"`
}

func (c *Client) ListChains() (*ListChainResponse, error) {
	get, err := c.get(ListChainsPath)
	if err != nil {
		log.Printf("ListChains error: %v", err)
		return nil, err
	}
	// Unmarshal response
	resp := &ListChainResponse{}
	err = json.Unmarshal(get, resp)
	if err != nil {
		log.Printf("ListChains error: %v", err)
		return nil, err
	}
	return resp, nil
}
