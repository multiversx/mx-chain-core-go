package transaction_test

import (
	"encoding/json"
	"errors"
	"math/big"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/mock"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_SettersAndGetters(t *testing.T) {
	t.Parallel()

	nonce := uint64(37)
	txData := []byte("data")
	value := big.NewInt(12)
	gasPrice := uint64(1)
	gasLimit := uint64(5)
	sender := []byte("sndr")
	receiver := []byte("receiver")

	tx := &transaction.Transaction{
		Nonce:    nonce,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}
	assert.False(t, check.IfNil(tx))

	tx.SetSndAddr(sender)
	tx.SetData(txData)
	tx.SetValue(value)
	tx.SetRcvAddr(receiver)

	assert.Equal(t, nonce, tx.GetNonce())
	assert.Equal(t, value, tx.GetValue())
	assert.Equal(t, txData, tx.GetData())
	assert.Equal(t, gasPrice, tx.GetGasPrice())
	assert.Equal(t, gasLimit, tx.GetGasLimit())
	assert.Equal(t, sender, tx.GetSndAddr())
	assert.Equal(t, receiver, tx.GetRcvAddr())
}

func TestTransaction_MarshalUnmarshalJsonShouldWork(t *testing.T) {
	t.Parallel()

	value := big.NewInt(445566)
	tx := &transaction.Transaction{
		Nonce:     112233,
		Value:     new(big.Int).Set(value),
		RcvAddr:   []byte("receiver"),
		SndAddr:   []byte("sender"),
		GasPrice:  1234,
		GasLimit:  5678,
		Data:      []byte("data"),
		Signature: []byte("signature"),
	}

	buff, err := json.Marshal(tx)
	assert.Nil(t, err)
	txRecovered := &transaction.Transaction{}
	err = json.Unmarshal(buff, txRecovered)
	assert.Nil(t, err)
	assert.Equal(t, tx, txRecovered)

	buffAsString := string(buff)
	assert.Contains(t, buffAsString, value.String())
}

func TestTransaction_TrimsSlicePtr(t *testing.T) {
	t.Parallel()

	tx1 := transaction.Transaction{
		Nonce:     1,
		Value:     big.NewInt(10),
		RcvAddr:   []byte("rcv"),
		SndAddr:   []byte("snd"),
		GasPrice:  1,
		GasLimit:  10,
		Data:      []byte("data"),
		Signature: []byte("sign"),
	}

	tx2 := tx1
	tx2.Nonce = 2

	input := make([]*transaction.Transaction, 0, 5)
	input = append(input, &tx1)
	input = append(input, &tx2)

	assert.Equal(t, 2, len(input))
	assert.Equal(t, 5, cap(input))

	input = transaction.TrimSlicePtr(input)

	assert.Equal(t, 2, len(input))
	assert.Equal(t, 2, cap(input))
}

func TestTransaction_TrimsSliceHandler(t *testing.T) {
	t.Parallel()

	tx1 := transaction.Transaction{
		Nonce:     1,
		Value:     big.NewInt(10),
		RcvAddr:   []byte("rcv"),
		SndAddr:   []byte("snd"),
		GasPrice:  1,
		GasLimit:  10,
		Data:      []byte("data"),
		Signature: []byte("sign"),
	}

	tx2 := tx1
	tx2.Nonce = 2

	input := make([]data.TransactionHandler, 0, 5)
	input = append(input, &tx1)
	input = append(input, &tx2)

	assert.Equal(t, 2, len(input))
	assert.Equal(t, 5, cap(input))

	input = transaction.TrimSliceHandler(input)

	assert.Equal(t, 2, len(input))
	assert.Equal(t, 2, cap(input))
}

func TestTransaction_GetDataForSigningNilPubkeyConverterShouldErr(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{}

	buff, err := tx.GetDataForSigning(nil, &mock.MarshalizerStub{}, &mock.HasherMock{})

	assert.Nil(t, buff)
	assert.Equal(t, transaction.ErrNilEncoder, err)
}

func TestTransaction_GetDataForSigningNilMarshalizerShouldErr(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{}

	buff, err := tx.GetDataForSigning(&mock.PubkeyConverterStub{}, nil, &mock.HasherMock{})

	assert.Nil(t, buff)
	assert.Equal(t, transaction.ErrNilMarshalizer, err)
}

func TestTransaction_GetDataForSigningNilHasherShouldErr(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{}

	buff, err := tx.GetDataForSigning(&mock.PubkeyConverterStub{}, &mock.MarshalizerStub{}, nil)

	assert.Nil(t, buff)
	assert.Equal(t, transaction.ErrNilHasher, err)
}

func TestTransaction_GetDataForSigningMarshalizerErrShouldErr(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{}

	numEncodeCalled := 0
	expectedErr := errors.New("expected error")
	buff, err := tx.GetDataForSigning(
		&mock.PubkeyConverterStub{
			EncodeCalled: func(pkBytes []byte) (string, error) {
				numEncodeCalled++
				return "", nil
			},
		},
		&mock.MarshalizerStub{
			MarshalCalled: func(obj interface{}) (bytes []byte, err error) {
				return nil, expectedErr
			},
		},
		&mock.HasherMock{},
	)

	assert.Nil(t, buff)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, 2, numEncodeCalled)
}

func TestTransaction_HasOptionHashSignSet(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{Options: 0}
	assert.False(t, tx.HasOptionHashSignSet())

	tx.Options ^= transaction.MaskSignedWithHash
	assert.True(t, tx.HasOptionHashSignSet())
}

func TestTransaction_HasOptionGuardianSet(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{Options: 0}
	assert.False(t, tx.HasOptionGuardianSet())

	tx.Options ^= transaction.MaskGuardedTransaction
	assert.True(t, tx.HasOptionGuardianSet())
}

func TestTransaction_GetDataForSigningShouldWork(t *testing.T) {
	t.Parallel()

	t.Run("without hash sign option set", func(t *testing.T) {
		t.Parallel()

		tx := &transaction.Transaction{}

		numEncodeCalled := 0
		marshalizerWasCalled := false
		hasherWasCalled := false
		buff, err := tx.GetDataForSigning(
			&mock.PubkeyConverterStub{
				EncodeCalled: func(pkBytes []byte) (string, error) {
					numEncodeCalled++
					return "", nil
				},
			},
			&mock.MarshalizerStub{
				MarshalCalled: func(obj interface{}) (bytes []byte, err error) {
					marshalizerWasCalled = true

					return make([]byte, 0), nil
				},
			},
			&mock.HasherStub{
				ComputeCalled: func(s string) []byte {
					hasherWasCalled = true

					return make([]byte, 0)
				},
			},
		)

		assert.Equal(t, 0, len(buff))
		assert.Nil(t, err)
		assert.True(t, marshalizerWasCalled)
		assert.False(t, hasherWasCalled)
		assert.Equal(t, 2, numEncodeCalled)
	})

	t.Run("with hash sign option set", func(t *testing.T) {
		t.Parallel()

		tx := &transaction.Transaction{}
		tx.Version = uint32(2)
		tx.Options ^= transaction.MaskSignedWithHash

		numEncodeCalled := 0
		marshalizerWasCalled := false
		hasherWasCalled := false
		expectedHash := []byte("expectedHash")
		buff, err := tx.GetDataForSigning(
			&mock.PubkeyConverterStub{
				EncodeCalled: func(pkBytes []byte) (string, error) {
					numEncodeCalled++
					return "", nil
				},
			},
			&mock.MarshalizerStub{
				MarshalCalled: func(obj interface{}) (bytes []byte, err error) {
					marshalizerWasCalled = true

					return make([]byte, 0), nil
				},
			},
			&mock.HasherStub{
				ComputeCalled: func(s string) []byte {
					hasherWasCalled = true

					return expectedHash
				},
			},
		)

		assert.Nil(t, err)
		assert.True(t, marshalizerWasCalled)
		assert.True(t, hasherWasCalled)
		assert.Equal(t, expectedHash, buff)
		assert.Equal(t, 2, numEncodeCalled)
	})
}

func TestTransaction_CheckIntegrityShouldWork(t *testing.T) {
	t.Parallel()

	tx := &transaction.Transaction{
		Nonce:       1,
		Value:       big.NewInt(10),
		GasPrice:    1,
		GasLimit:    10,
		Data:        []byte("data"),
		Signature:   []byte("signature"),
		RcvUserName: []byte("rcv-username"),
		SndUserName: []byte("snd-username"),
	}

	err := tx.CheckIntegrity()
	assert.Nil(t, err)
}

func TestTransaction_CheckIntegrityShouldErr(t *testing.T) {
	t.Parallel()

	invalidUsername := []byte("invalid-username-length-exceeds-max-allowed-length")
	validUsername := []byte("unittest")
	tx := &transaction.Transaction{
		Nonce: 1,
		Data:  []byte("data"),
	}

	err := tx.CheckIntegrity()
	assert.Equal(t, data.ErrNilSignature, err)

	tx.Signature = []byte("signature")

	err = tx.CheckIntegrity()
	assert.Equal(t, data.ErrNilValue, err)

	tx.Value = big.NewInt(-1)

	err = tx.CheckIntegrity()
	assert.Equal(t, data.ErrNegativeValue, err)

	tx.Value = big.NewInt(10)
	tx.RcvUserName = invalidUsername

	err = tx.CheckIntegrity()
	assert.Equal(t, data.ErrInvalidUserNameLength, err)

	tx.RcvUserName = validUsername
	tx.SndUserName = invalidUsername

	err = tx.CheckIntegrity()
	assert.Equal(t, data.ErrInvalidUserNameLength, err)
}

func TestTransaction_ImplementsGuardedTransactionHandler(t *testing.T) {
	t.Parallel()

	var tx data.TransactionHandler = &transaction.Transaction{}

	_, ok := tx.(data.GuardedTransactionHandler)
	assert.True(t, ok)
}
