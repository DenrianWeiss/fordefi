package test

import (
	"encoding/json"
	"github.com/DenrianWeiss/fordefi/test/common"
	"testing"
)

func TestListVault(t *testing.T) {
	client := common.GetClient()
	resp, err := client.ListVaults()
	if err != nil {
		t.Errorf("ListVaults error: %v", err)
	}
	// Print response
	/// Marshal response
	respStr, _ := json.Marshal(resp)
	t.Logf("ListVaults response: %s", string(respStr))
}
