package fordefi

import (
	"encoding/json"
	"fmt"
)

// ApproveTx approve transaction.
// @param txId: transaction id(fordefi uuid).
// @return resp: response body can be nil if err is nil(http204).
func (c *Client) ApproveTx(txId string) (resp []byte, err error) {
	respRaw, err := c.post(fmt.Sprintf(ApproveTransactionPath, txId), nil, true)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	return respRaw, nil
}

// AbortTx abort transaction.
// @param txId: transaction id(fordefi uuid).
// @return resp: response body can be nil if err is nil(http204).
func (c *Client) AbortTx(txId string) (resp []byte, err error) {
	respRaw, err := c.post(fmt.Sprintf(AbortTransactionPath, txId), nil, true)
	if err != nil {
		return nil, err
	}
	// Deserialize response
	return respRaw, nil
}

// ReleaseTx release transaction.
// @param txId: transaction id(fordefi uuid).
// @param releaseType: release type, `ReleaseCancel` or `ReleaseAccelerate`.
// @return resp: response same as EvmTransactResp.
func (c *Client) ReleaseTx(txId string, releaseType string) (resp *EvmTransactResp, err error) {
	reqBody := map[string]interface{}{
		"type":         "evm_transaction",
		"release_type": releaseType,
	}
	respRaw, err := c.post(fmt.Sprintf(ReleaseTransactionPath, txId), reqBody, true)
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
