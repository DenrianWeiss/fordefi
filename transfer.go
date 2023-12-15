package fordefi

import (
	"encoding/json"
	"math/big"
)

func (r *EvmTransferReq) PrepareForEvmTransfer() {
	r.SignerType = "api_signer"
	r.Type = "evm_transaction"
	r.Details.Type = "evm_transfer"
}

func (r *SolanaTransferReq) PrepareForSolanaTransfer() {
	r.SignerType = "api_signer"
	r.Type = "solana_transaction"
	r.Details.Type = "solana_transfer"
}

// EvmTransferNative transfer native token on evm chains.
// @param vaultId: vault id.
// @param chain: chain name.
// @param gasOps: gas options.
// @param amount: amount to transfer, pass -1 to tran.
func (c *Client) EvmTransferNative(vaultId string, chain string, gasOps EvmGasOps, amount *big.Int, to string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmTransferReq{}
	(&reqBody).PrepareForEvmTransfer()
	reqBody.Details.Gas = gasOps
	reqBody.Details.To = to
	reqBody.Details.AssetIdentifier.Type = "evm"
	reqBody.Details.AssetIdentifier.Details = map[string]interface{}{
		"type":  "native",
		"chain": chain,
	}
	if amount.Cmp(big.NewInt(MaxRepr)) == 0 {
		reqBody.Details.Value.Type = "max"
	} else {
		reqBody.Details.Value.Type = "value"
		reqBody.Details.Value.Value = amount.String()
	}
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post("/v1/transaction", reqBody)
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

// EvmTransferErc20 transfer erc20 token on evm chains.
// @param vaultId: vault id.
// @param chain: chain name.
// @param tokenAddr: gas options.
// @param gasOps: gas options.
// @param amount: amount to transfer, pass -1 to transfer max.
func (c *Client) EvmTransferErc20(vaultId string, chain string, tokenAddr string, gasOps EvmGasOps, amount *big.Int, to string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmTransferReq{}
	(&reqBody).PrepareForEvmTransfer()
	reqBody.Details.Gas = gasOps
	reqBody.Details.To = to
	reqBody.Details.AssetIdentifier.Type = "evm"
	reqBody.Details.AssetIdentifier.Details = map[string]interface{}{
		"type": "erc20",
		"token": map[string]interface{}{
			"chain":    chain,
			"hex_repr": tokenAddr,
		},
	}
	if amount.Cmp(big.NewInt(MaxRepr)) == 0 {
		reqBody.Details.Value.Type = "max"
	} else {
		reqBody.Details.Value.Type = "value"
		reqBody.Details.Value.Value = amount.String()
	}
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

// EvmTransferErc721 transfer erc721 token on evm chains.
// @param vaultId: vault id.
// @param chain: chain name.
// @param tokenAddr: gas options.
// @param gasOps: gas options.
// @param tokenId: token id to transfer.
func (c *Client) EvmTransferErc721(vaultId string, chain string, tokenAddr string, gasOps EvmGasOps, tokenId *big.Int, to string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmTransferReq{}
	(&reqBody).PrepareForEvmTransfer()
	reqBody.Details.Gas = gasOps
	reqBody.Details.To = to
	reqBody.Details.AssetIdentifier.Type = "evm"
	reqBody.Details.AssetIdentifier.Details = map[string]interface{}{
		"type": "erc721",
		"token": map[string]interface{}{
			"chain":    chain,
			"hex_repr": tokenAddr,
		},
		"token_id": tokenId.String(),
	}
	reqBody.Details.Value.Type = "value"
	reqBody.Details.Value.Value = tokenId.String()
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	resp = &EvmTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// EvmTransferErc1155 transfer erc1155 token on evm chains.
// @param vaultId: vault id.
// @param chain: chain name.
// @param tokenAddr: gas options.
// @param gasOps: gas options.
// @param tokenId: token id to transfer.
func (c *Client) EvmTransferErc1155(vaultId string, chain string, tokenAddr string, gasOps EvmGasOps, tokenId *big.Int, to string, note ...string) (resp *EvmTransactResp, err error) {
	reqBody := EvmTransferReq{}
	(&reqBody).PrepareForEvmTransfer()
	reqBody.Details.Gas = gasOps
	reqBody.Details.To = to
	reqBody.Details.AssetIdentifier.Type = "evm"
	reqBody.Details.AssetIdentifier.Details = map[string]interface{}{
		"type": "erc1155",
		"token": map[string]interface{}{
			"chain":    chain,
			"hex_repr": tokenAddr,
		},
		"token_id": tokenId.String(),
	}
	reqBody.Details.Value.Type = "value"
	reqBody.Details.Value.Value = tokenId.String()
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	resp = &EvmTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SolanaTransferNative transfer native token on solana chains.
// @param vaultId: vault id.
// @param amount: amount to transfer, pass -1 to transfer max.
// @param to: to address.
func (c *Client) SolanaTransferNative(vaultId string, amount *big.Int, to string, note ...string) (resp *SolanaTransactResp, err error) {
	reqBody := SolanaTransferReq{}
	(&reqBody).PrepareForSolanaTransfer()
	reqBody.Details.To = to
	reqBody.Details.AssetIdentifier.Type = "solana"
	reqBody.Details.AssetIdentifier.Details = map[string]interface{}{
		"type":  "native",
		"chain": SolanaMainnet,
	}
	if amount.Cmp(big.NewInt(MaxRepr)) == 0 {
		reqBody.Details.Value.Type = "max"
	} else {
		reqBody.Details.Value.Type = "value"
		reqBody.Details.Value.Value = amount.String()
	}
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	// decode response
	resp = &SolanaTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SolanaTransferSpl transfer spl token on solana chains.
// @param vaultId: vault id.
// @param tokenAddr: token address.
// @param amount: amount to transfer, pass -1 to transfer max.
// @param to: to address.
func (c *Client) SolanaTransferSpl(vaultId string, tokenAddr string, amount *big.Int, to string, note ...string) (resp *SolanaTransactResp, err error) {
	reqBody := SolanaTransferReq{}
	(&reqBody).PrepareForSolanaTransfer()
	reqBody.Details.To = to
	reqBody.Details.AssetIdentifier.Type = "solana"
	reqBody.Details.AssetIdentifier.Details = map[string]interface{}{
		"type": "spl",
		"token": map[string]interface{}{
			"chain":       SolanaMainnet,
			"base58_repr": tokenAddr,
		},
	}
	if amount.Cmp(big.NewInt(MaxRepr)) == 0 {
		reqBody.Details.Value.Type = "max"
	} else {
		reqBody.Details.Value.Type = "value"
		reqBody.Details.Value.Value = amount.String()
	}
	reqBody.VaultId = vaultId
	if len(note) > 0 {
		reqBody.Note = note[0]
	}
	respRaw, err := c.post(CreateTransactionPath, reqBody)
	if err != nil {
		return nil, err
	}
	// decode response
	resp = &SolanaTransactResp{}
	err = json.Unmarshal(respRaw, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
