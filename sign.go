package fordefi

import "encoding/json"

type EvmSignRequest struct {
	SignerType string `json:"signer_type"`
	Type       string `json:"type"`
	Details    struct {
		Type    string `json:"type"`
		RawData string `json:"raw_data"`
		Chain   string `json:"chain"`
	} `json:"details"`
	VaultId string `json:"vault_id"`
	Note    string `json:"note"`
}

// EvmSignMessage sign message using EIP191 personal sign or EIP712 typed sign.
// @param vaultId: vault id.
// @param chain: chain name.
// @param data: message data.
// @param signType: sign type, `SignPersonal` or `SignTyped`.
// @param note: note.
func (c *Client) EvmSignMessage(vaultId string, chain string, data string, signType string, note ...string) (resp *EvmSignResp, err error) {
	reqBody := EvmSignRequest{}
	reqBody.SignerType = "api_signer"
	reqBody.Type = "evm_message"
	reqBody.Details.Type = signType
	reqBody.Details.RawData = data
	reqBody.Details.Chain = chain
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	resp = &EvmSignResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
