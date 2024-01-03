package test

import (
	"encoding/json"
	"github.com/DenrianWeiss/fordefi"
	"github.com/DenrianWeiss/fordefi/test/common"
	"os"
	"strings"
	"testing"
)

const KeyDerivationPayload = "{\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"}],\"dYdX\":[{\"name\":\"action\",\"type\":\"string\"}]},\"primaryType\":\"dYdX\",\"domain\":{\"name\":\"dYdX\",\"version\":\"1.0\",\"chainId\":\"0xaa36a7\"},\"message\":{\"action\":\"dYdX STARK Key\"}} \n"

func TestSignEip712(t *testing.T) {
	client := common.GetClient()
	vaultId, exist := os.LookupEnv("FORDEFI_VAULT_ID")
	if !exist {
		t.Errorf("FORDEFI_VAULT_ID is not set")
	}
	chainId := strings.TrimPrefix(fordefi.SepoliaTestnet, "evm_")
	// Assemble EVM Sign Request
	resp, err := client.EvmSignMessage(vaultId, chainId, KeyDerivationPayload, fordefi.SignTyped, "test signing request")
	if err != nil {
		t.Errorf("EvmSignMessage error: %v", err)
	}
	// Print response
	/// Marshal response
	respStr, _ := json.Marshal(resp)
	t.Logf("EvmSignMessage response: %s", string(respStr))
}
