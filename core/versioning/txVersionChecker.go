package versioning

import (
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data/transaction"
)

// TxVersionChecker represents transaction option decoder
type txVersionChecker struct {
	minTxVersion uint32
}

// NewTxVersionChecker will create a new instance of TxOptionsChecker
func NewTxVersionChecker(minTxVersion uint32) *txVersionChecker {
	return &txVersionChecker{
		minTxVersion: minTxVersion,
	}
}

// IsSignedWithHash will return true if transaction is signed with hash
func (tvc *txVersionChecker) IsSignedWithHash(tx *transaction.Transaction) bool {
	if tx.Version > core.InitialVersionOfTransaction {
		// transaction is signed with hash if LSB from last byte from options is set with 1
		return tx.HasOptionHashSignSet()
	}

	return false
}

// IsGuardedTransaction will return true if transaction also holds a guardian signature
func (tvc *txVersionChecker) IsGuardedTransaction(tx *transaction.Transaction) bool {
	if tx.Version > core.InitialVersionOfTransaction {
		return tx.HasOptionGuardianSet()
	}

	return false
}

// CheckTxVersion will check transaction version
func (tvc *txVersionChecker) CheckTxVersion(tx *transaction.Transaction) error {
	if (tx.Version == core.InitialVersionOfTransaction && tx.Options != 0) || tx.Version < tvc.minTxVersion {
		return core.ErrInvalidTransactionVersion
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (tvc *txVersionChecker) IsInterfaceNil() bool {
	return tvc == nil
}
