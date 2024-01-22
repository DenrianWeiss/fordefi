package fordefi

import (
	"encoding/json"
	"math/big"
)

type EvmRawTransactReq struct {
	VaultId    string `json:"vault_id"`
	Note       string `json:"note"`
	SignerType string `json:"signer_type"`
	Type       string `json:"type"`
	Details    struct {
		Type          string                 `json:"type"`
		UseSecureNode bool                   `json:"use_secure_node"`
		Chain         string                 `json:"chain"`
		Gas           EvmGasOps              `json:"gas"`
		To            string                 `json:"to"`
		Value         string                 `json:"value"`
		Data          map[string]interface{} `json:"data"`
	} `json:"details"`
}

func (e *EvmRawTransactReq) Prepare() {
	e.SignerType = "api_signer"
	e.Type = "evm_transaction"
}

// EvmTransactByName using function name and args to call contract function.
// @param vaultId: vault id.
// @param chain: chain name.
// @param gasOps: gas options.
// @param to: tx to address.
// @param amount: value to transfer
// @param fnName: function name.
// @param args: function args, in `["name1:value1", "name2:value2"]` format.
// @param note: note.
func (c *Client) EvmTransactByName(vaultId string, chain string, gasOps EvmGasOps, to string, amount *big.Int, fnName string, args []string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmRawTransactReq{}
	(&reqBody).Prepare()
	reqBody.Details.Gas = gasOps
	reqBody.Details.Type = "evm_raw_transaction"
	reqBody.Details.Chain = chain
	reqBody.Details.To = to
	reqBody.Details.Data = map[string]interface{}{
		"type":             ModeByFunctionName,
		"method_name":      fnName,
		"method_arguments": args,
	}
	reqBody.Details.Value = amount.String()
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	resp = &EvmTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// EvmTransactByData using raw data to call contract function.
// @param vaultId: vault id.
// @param chain: chain name.
// @param gasOps: gas options.
// @param to: tx to address.
// @param amount: value to transfer
// @param data: raw data.
// @param note: note.
func (c *Client) EvmTransactByData(vaultId string, chain string, gasOps EvmGasOps, to string, amount *big.Int, data string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmRawTransactReq{}
	(&reqBody).Prepare()
	reqBody.Details.Type = "evm_raw_transaction"
	reqBody.Details.Gas = gasOps
	reqBody.Details.Chain = chain
	reqBody.Details.Data = map[string]interface{}{
		"type":     ModeByHexData,
		"hex_data": data,
	}
	reqBody.Details.To = to
	reqBody.Details.Value = amount.String()
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	resp = &EvmTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// EvmTransactByBase64Data using base64 encoded data to call contract function.
// @param vaultId: vault id.
// @param chain: chain name.
// @param gasOps: gas options.
// @param to: tx to address.
// @param amount: value to transfer
// @param data: base64 encoded data.
// @param note: note.
func (c *Client) EvmTransactByBase64Data(vaultId string, chain string, gasOps EvmGasOps, to string, amount *big.Int, data string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmRawTransactReq{}
	(&reqBody).Prepare()
	reqBody.Details.Gas = gasOps
	reqBody.Details.Type = "evm_raw_transaction"
	reqBody.Details.Chain = chain
	reqBody.Details.Data = map[string]interface{}{
		"type":     ModeByBase64Data,
		"raw_data": data,
	}
	reqBody.Details.To = to
	reqBody.Details.Value = amount.String()
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	resp = &EvmTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
