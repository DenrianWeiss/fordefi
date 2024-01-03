package test

import (
	"encoding/json"
	"github.com/DenrianWeiss/fordefi/test/common"
	"os"
	"testing"
)

func TestGetTxById(t *testing.T) {
	client := common.GetClient()
	txId, exist := os.LookupEnv("FORDEFI_TX_ID")
	if !exist {
		t.Errorf("FORDEFI_TX_ID is not set")
	}
	resp, err := client.GetTransactionEvmSign(txId)
	if err != nil {
		t.Errorf("GetTransactionEvmSign error: %v", err)
	}
	// Print response
	jData, _ := json.Marshal(resp)
	t.Logf("GetTransactionEvmSign response: %s", string(jData))
}
