package transaction

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/stretchr/testify/assert"
)

func Test_SortTransactionsBySenderAndNonceWithFrontRunningProtection(t *testing.T) {
	randomness := "randomness"
	nbSenders := 5

	hasher := &mock.HasherStub{
		ComputeCalled: func(s string) []byte {
			if s == randomness {
				return []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
			}

			return []byte(s)
		},
	}

	usedRandomness := hasher.Compute(randomness)
	senders := make([][]byte, 0)
	for i := 0; i < nbSenders; i++ {
		sender := make([]byte, len(usedRandomness))
		copy(sender, usedRandomness)
		sender[len(usedRandomness)-1-i] = 0
		senders = append(senders, sender)
	}

	txs := []data.TransactionHandler{
		&transaction.Transaction{Nonce: 1, SndAddr: senders[0]},
		&transaction.Transaction{Nonce: 2, SndAddr: senders[2]},
		&transaction.Transaction{Nonce: 1, SndAddr: senders[2]},
		&transaction.Transaction{Nonce: 2, SndAddr: senders[0]},
		&transaction.Transaction{Nonce: 7, SndAddr: senders[1]},
		&transaction.Transaction{Nonce: 6, SndAddr: senders[1]},
		&transaction.Transaction{Nonce: 1, SndAddr: senders[4]},
		&transaction.Transaction{Nonce: 3, SndAddr: senders[3]},
		&transaction.Transaction{Nonce: 3, SndAddr: senders[2]},
	}
	wrappedTxs := make([]data.TxWithExecutionOrderHandler, 0, len(txs))
	for _, tx := range txs {
		wrappedTxs = append(wrappedTxs, &outport.TxInfo{
			Transaction: tx.(*transaction.Transaction),
			FeeInfo:     &outport.FeeInfo{Fee: big.NewInt(0)}})
	}
	SortTransactionsBySenderAndNonceWithFrontRunningProtection(txs, hasher, []byte(randomness))
	SortTransactionsBySenderAndNonceWithFrontRunningProtectionExtendedTransactions(wrappedTxs, hasher, []byte(randomness))

	expectedOutput := []string{
		"1 ffffffffffffffffffffffffffffff00",
		"2 ffffffffffffffffffffffffffffff00",
		"6 ffffffffffffffffffffffffffff00ff",
		"7 ffffffffffffffffffffffffffff00ff",
		"1 ffffffffffffffffffffffffff00ffff",
		"2 ffffffffffffffffffffffffff00ffff",
		"3 ffffffffffffffffffffffffff00ffff",
		"3 ffffffffffffffffffffffff00ffffff",
		"1 ffffffffffffffffffffff00ffffffff",
	}

	for i, item := range txs {
		assert.Equal(t, expectedOutput[i], fmt.Sprintf("%d %s", item.GetNonce(), hex.EncodeToString(item.GetSndAddr())))
		assert.Equal(t, expectedOutput[i], fmt.Sprintf("%d %s", wrappedTxs[i].GetTxHandler().GetNonce(), hex.EncodeToString(wrappedTxs[i].GetTxHandler().GetSndAddr())))
	}
}

func Test_SortTransactionsBySenderAndNonceLegacy(t *testing.T) {
	txs := []data.TransactionHandler{
		&transaction.Transaction{Nonce: 3, SndAddr: []byte("bbbb")},
		&transaction.Transaction{Nonce: 1, SndAddr: []byte("aaaa")},
		&transaction.Transaction{Nonce: 5, SndAddr: []byte("bbbb")},
		&transaction.Transaction{Nonce: 2, SndAddr: []byte("aaaa")},
		&transaction.Transaction{Nonce: 7, SndAddr: []byte("aabb")},
		&transaction.Transaction{Nonce: 6, SndAddr: []byte("aabb")},
		&transaction.Transaction{Nonce: 3, SndAddr: []byte("ffff")},
		&transaction.Transaction{Nonce: 3, SndAddr: []byte("eeee")},
	}
	wrappedTxs := make([]data.TxWithExecutionOrderHandler, 0, len(txs))
	for _, tx := range txs {
		wrappedTxs = append(wrappedTxs, &outport.TxInfo{
			Transaction: tx.(*transaction.Transaction),
			FeeInfo:     &outport.FeeInfo{Fee: big.NewInt(0)}})
	}

	SortTransactionsBySenderAndNonce(txs)
	SortTransactionsBySenderAndNonceExtendedTransactions(wrappedTxs)

	expectedOutput := []string{
		"1 aaaa",
		"2 aaaa",
		"6 aabb",
		"7 aabb",
		"3 bbbb",
		"5 bbbb",
		"3 eeee",
		"3 ffff",
	}

	for i, item := range txs {
		assert.Equal(t, expectedOutput[i], fmt.Sprintf("%d %s", item.GetNonce(), string(item.GetSndAddr())))
		assert.Equal(t, expectedOutput[i], fmt.Sprintf("%d %s", wrappedTxs[i].GetTxHandler().GetNonce(), string(wrappedTxs[i].GetTxHandler().GetSndAddr())))
	}
}
