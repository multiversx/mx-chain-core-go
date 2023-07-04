package transaction

import (
	"bytes"
	"sort"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/hashing"
)

// SortTransactionsBySenderAndNonceWithFrontRunningProtection - sorts the transactions by address and randomness source to protect from front running
func SortTransactionsBySenderAndNonceWithFrontRunningProtection(transactions []data.TransactionHandler, hasher hashing.Hasher, randomness []byte) {
	// make sure randomness is 32bytes and uniform
	randSeed := hasher.Compute(string(randomness))
	xoredAddresses := make(map[string][]byte)

	for _, tx := range transactions {
		xoredBytes := xorBytes(tx.GetSndAddr(), randSeed)
		xoredAddresses[string(tx.GetSndAddr())] = hasher.Compute(string(xoredBytes))
	}

	sorter := func(i, j int) bool {
		txI := transactions[i]
		txJ := transactions[j]

		delta := bytes.Compare(xoredAddresses[string(txI.GetSndAddr())], xoredAddresses[string(txJ.GetSndAddr())])
		if delta == 0 {
			delta = int(txI.GetNonce()) - int(txJ.GetNonce())
		}

		return delta < 0
	}

	sort.Slice(transactions, sorter)
}

// TODO remove duplicated function when will use the version of mx-chain-go which exports transaction order during processing

// SortTransactionsBySenderAndNonceWithFrontRunningProtectionExtendedTransactions - sorts the transactions by address and randomness source to protect from front running
func SortTransactionsBySenderAndNonceWithFrontRunningProtectionExtendedTransactions(transactions []data.TxWithExecutionOrderHandler, hasher hashing.Hasher, randomness []byte) {
	// make sure randomness is 32bytes and uniform
	randSeed := hasher.Compute(string(randomness))
	xoredAddresses := make(map[string][]byte)

	for _, tx := range transactions {
		txHandler := tx.GetTxHandler()

		xoredBytes := xorBytes(txHandler.GetSndAddr(), randSeed)
		xoredAddresses[string(txHandler.GetSndAddr())] = hasher.Compute(string(xoredBytes))
	}

	sorter := func(i, j int) bool {
		txI := transactions[i]
		txJ := transactions[j]
		txIHandler := txI.GetTxHandler()
		txJHandler := txJ.GetTxHandler()

		delta := bytes.Compare(xoredAddresses[string(txIHandler.GetSndAddr())], xoredAddresses[string(txJHandler.GetSndAddr())])
		if delta == 0 {
			delta = int(txIHandler.GetNonce()) - int(txJHandler.GetNonce())
		}

		return delta < 0
	}

	sort.Slice(transactions, sorter)
}

// SortTransactionsBySenderAndNonce - sorts the transactions by address without the front running protection
func SortTransactionsBySenderAndNonce(transactions []data.TransactionHandler) {
	sorter := func(i, j int) bool {
		txI := transactions[i]
		txJ := transactions[j]

		delta := bytes.Compare(txI.GetSndAddr(), txJ.GetSndAddr())
		if delta == 0 {
			delta = int(txI.GetNonce()) - int(txJ.GetNonce())
		}

		return delta < 0
	}

	sort.Slice(transactions, sorter)
}

// SortTransactionsBySenderAndNonceExtendedTransactions - sorts the transactions by address without the front running protection
func SortTransactionsBySenderAndNonceExtendedTransactions(transactions []data.TxWithExecutionOrderHandler) {
	sorter := func(i, j int) bool {
		txI := transactions[i]
		txJ := transactions[j]
		txIHandler := txI.GetTxHandler()
		txJHandler := txJ.GetTxHandler()

		delta := bytes.Compare(txIHandler.GetSndAddr(), txJHandler.GetSndAddr())
		if delta == 0 {
			delta = int(txIHandler.GetNonce()) - int(txJHandler.GetNonce())
		}

		return delta < 0
	}

	sort.Slice(transactions, sorter)
}

// parameters need to be of the same len, otherwise it will panic (if second slice shorter)
func xorBytes(a, b []byte) []byte {
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}
