package versioning

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
)

// There are different options based on the tx version. If version is the initial version of transaction no options can be used.
// Otherwise, if version is higher than the initial version, several options can be used:
// bit 0: signed over hash - if set, the transaction signature was applied over the hash of the marshalled tx, otherwise it was applied directly on the marshalled tx
// bit 1: guarded Tx - if set, the transaction is guarded (co signed) by the sender's configured guardian which should be the same as the one configured in the Tx GuardianAddr field. The guardian signature is represented by field GuardianSignature

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
