package transaction

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/transaction"
	"github.com/stretchr/testify/assert"
)

func Test_sortTransactionsBySenderAndNonceWithFrontRunningProtection(t *testing.T) {
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

	SortTransactionsBySenderAndNonceWithFrontRunningProtection(txs, hasher, []byte(randomness))

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
	}
}
