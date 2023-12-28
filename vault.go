package fordefi

import (
	"encoding/json"
	"time"
)

type ListVaultsResp struct {
	Total  int `json:"total"`
	Page   int `json:"page"`
	Size   int `json:"size"`
	Vaults []struct {
		Id         string    `json:"id"`
		CreatedAt  time.Time `json:"created_at"`
		ModifiedAt time.Time `json:"modified_at"`
		Name       string    `json:"name"`
		CreatedBy  struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"created_by"`
		DerivationPath      string `json:"derivation_path"`
		PublicKeyCompressed string `json:"public_key_compressed"`
		DerivationInfo      struct {
			DerivationPath  string `json:"derivation_path"`
			MasterPublicKey struct {
				Id   string `json:"id"`
				Xpub string `json:"xpub"`
			} `json:"master_public_key"`
		} `json:"derivation_info"`
		Keyset struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Scope string `json:"scope"`
		} `json:"keyset"`
		KeyHolder struct {
			Id         string `json:"id"`
			UserType   string `json:"user_type"`
			ExternalId string `json:"external_id"`
			State      string `json:"state"`
		} `json:"key_holder"`
		VaultGroup struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			VaultCount int    `json:"vault_count"`
		} `json:"vault_group"`
		PendingVaultGroupAction struct {
			Type           string `json:"type"`
			VaultGroupId   string `json:"vault_group_id"`
			VaultGroupName string `json:"vault_group_name"`
		} `json:"pending_vault_group_action"`
		Type    string `json:"type"`
		Details struct {
			Type      string `json:"type"`
			PublicKey string `json:"public_key"`
			StarkKey  string `json:"stark_key"`
		} `json:"details"`
	} `json:"vaults"`
}

func (c *Client) ListVaults() (resp *ListVaultsResp, err error) {
	respRaw, err := c.get(ListVaultsPath)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	resp = &ListVaultsResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
