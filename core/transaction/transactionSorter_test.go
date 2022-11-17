package transaction

import (
	"encoding/hex"
	"fmt"
	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/transaction"
)

func Example_sortTransactionsBySenderAndNonceWithFrontRunningProtection() {
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

	txs := []*transaction.Transaction{
		{Nonce: 1, SndAddr: senders[0]},
		{Nonce: 2, SndAddr: senders[2]},
		{Nonce: 1, SndAddr: senders[2]},
		{Nonce: 2, SndAddr: senders[0]},
		{Nonce: 7, SndAddr: senders[1]},
		{Nonce: 6, SndAddr: senders[1]},
		{Nonce: 1, SndAddr: senders[4]},
		{Nonce: 3, SndAddr: senders[3]},
		{Nonce: 3, SndAddr: senders[2]},
	}

	SortTransactionsBySenderAndNonceWithFrontRunningProtection(txs, hasher, []byte(randomness))

	for _, item := range txs {
		fmt.Println(item.GetNonce(), hex.EncodeToString(item.GetSndAddr()))
	}

	// Output:
	// 1 ffffffffffffffffffffffffffffff00
	// 2 ffffffffffffffffffffffffffffff00
	// 6 ffffffffffffffffffffffffffff00ff
	// 7 ffffffffffffffffffffffffffff00ff
	// 1 ffffffffffffffffffffffffff00ffff
	// 2 ffffffffffffffffffffffffff00ffff
	// 3 ffffffffffffffffffffffffff00ffff
	// 3 ffffffffffffffffffffffff00ffffff
	// 1 ffffffffffffffffffffff00ffffffff
}
