package sovereignTx_test

import (
	"math/big"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/sovereignTx"
	"github.com/stretchr/testify/assert"
)

func TestSovereignTx_SettersAndGetters(t *testing.T) {
	t.Parallel()

	nonce := uint64(5)
	gasPrice := uint64(1)
	gasLimit := uint64(10)
	stx := sovereignTx.SovereignTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}

	rcvAddr := []byte("rcv address")
	sndAddr := []byte("snd address")
	value := big.NewInt(37)
	txData := []byte("data")

	stx.SetRcvAddr(rcvAddr)
	stx.SetSndAddr(sndAddr)
	stx.SetValue(value)
	stx.SetData(txData)

	assert.Equal(t, sndAddr, stx.GetSndAddr())
	assert.Equal(t, rcvAddr, stx.GetRcvAddr())
	assert.Equal(t, value, stx.GetValue())
	assert.Equal(t, txData, stx.GetData())
	assert.Equal(t, gasLimit, stx.GetGasLimit())
	assert.Equal(t, gasPrice, stx.GetGasPrice())
	assert.Equal(t, nonce, stx.GetNonce())
}

func createValidSovereignTx() *sovereignTx.SovereignTx {
	return &sovereignTx.SovereignTx{
		Nonce:    1,
		Value:    big.NewInt(10),
		GasPrice: 1,
		GasLimit: 10,
		Data:     []byte("data"),
		RcvAddr:  []byte("rcv-address"),
		SndAddr:  []byte("snd-address"),
	}
}

func TestSovereignTx_CheckIntegrityShouldWork(t *testing.T) {
	t.Parallel()

	stx := createValidSovereignTx()
	err := stx.CheckIntegrity()

	assert.Nil(t, err)
}

func TestSovereignTx_CheckIntegrityShouldErr(t *testing.T) {
	t.Parallel()

	stx := createValidSovereignTx()
	stx.RcvAddr = nil
	err := stx.CheckIntegrity()

	assert.Equal(t, data.ErrNilRcvAddr, err)

	stx = createValidSovereignTx()
	stx.SndAddr = nil
	err = stx.CheckIntegrity()

	assert.Equal(t, data.ErrNilSndAddr, err)

	stx = createValidSovereignTx()
	stx.Value = nil
	err = stx.CheckIntegrity()

	assert.Equal(t, data.ErrNilValue, err)

	stx = createValidSovereignTx()
	stx.Value = big.NewInt(-1)
	err = stx.CheckIntegrity()

	assert.Equal(t, data.ErrNegativeValue, err)
}
