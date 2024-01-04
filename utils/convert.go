package utils

import (
	"errors"
	"github.com/DenrianWeiss/fordefi"
	"log"
	"time"
)

var ErrTimeout = errors.New("timeout")
var ErrAborted = errors.New("aborted")

var DefaultCountLimit = 10 // 10 * 3 sec

func WaitUntilSigned(txUuid string, client *fordefi.Client, countLimit ...int) (string, error) {
	// Create a ticker at 3 sec
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// Set count limit
	count := DefaultCountLimit
	if len(countLimit) > 0 {
		count = countLimit[0]
	}
	for {
		if count <= 0 {
			log.Printf("WaitUntilSigned timeout")
			return "", ErrTimeout
		}
		count--
		select {
		case <-ticker.C:
			{
				// Get transaction
				tx, err := client.GetTransactionEvmTx(txUuid)
				if err != nil {
					log.Printf("GetTransactionEvmTx error: %v", err)
					continue
				}
				if fordefi.OnChainStatuses[tx.State] {
					log.Printf("WaitUntilSigned success")
					return tx.Hash, nil
				}
				if tx.State == fordefi.StatusAborted {
					log.Printf("WaitUntilSigned aborted")
					return "", ErrAborted
				}
			}
		}
	}
}
