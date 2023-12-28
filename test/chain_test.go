package test

import (
	"encoding/json"
	"github.com/DenrianWeiss/fordefi/test/common"
	"testing"
)

func TestGetChainList(t *testing.T) {
	client := common.GetClient()
	resp, err := client.ListChains()
	if err != nil {
		t.Errorf("ListChains error: %v", err)
	}
	// Print response
	/// Marshal response
	respStr, _ := json.Marshal(resp)
	t.Logf("ListChains response: %s", string(respStr))
}
