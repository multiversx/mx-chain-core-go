package addressGenerator

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/core/pubkeyConverter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AddressBytesLen represents the number of bytes of an address
const AddressBytesLen = 32

var AddressPublicKeyConverter, _ = pubkeyConverter.NewBech32PubkeyConverter(AddressBytesLen, &mock.LoggerMock{})

func TestBlockChainHookImpl_NewAddressLengthNoGood(t *testing.T) {
	t.Parallel()

	ag, err := NewAddressGenerator(AddressPublicKeyConverter)
	require.Nil(t, err)

	address := []byte("test")
	nonce := uint64(10)

	scAddress, err := ag.NewAddress(address, nonce, []byte("00"))
	assert.Equal(t, ErrAddressLengthNotCorrect, err)
	assert.Nil(t, scAddress)

	address = []byte("1234567890123456789012345678901234567890")
	scAddress, err = ag.NewAddress(address, nonce, []byte("00"))
	assert.Equal(t, ErrAddressLengthNotCorrect, err)
	assert.Nil(t, scAddress)
}

func TestBlockChainHookImpl_NewAddressVMTypeTooLong(t *testing.T) {
	t.Parallel()

	ag, err := NewAddressGenerator(AddressPublicKeyConverter)
	require.Nil(t, err)

	address := []byte("01234567890123456789012345678900")
	nonce := uint64(10)

	vmType := []byte("010")
	scAddress, err := ag.NewAddress(address, nonce, vmType)
	assert.Equal(t, ErrVMTypeLengthIsNotCorrect, err)
	assert.Nil(t, scAddress)
}

func TestBlockChainHookImpl_NewAddress(t *testing.T) {
	t.Parallel()

	ag, err := NewAddressGenerator(AddressPublicKeyConverter)
	require.Nil(t, err)

	address := []byte("01234567890123456789012345678900")
	nonce := uint64(10)

	vmType := []byte("11")
	scAddress1, err := ag.NewAddress(address, nonce, vmType)
	assert.Nil(t, err)

	for i := 0; i < 8; i++ {
		assert.Equal(t, scAddress1[i], uint8(0))
	}
	assert.True(t, bytes.Equal(vmType, scAddress1[8:10]))

	nonce++
	scAddress2, err := ag.NewAddress(address, nonce, []byte("00"))
	assert.Nil(t, err)

	assert.False(t, bytes.Equal(scAddress1, scAddress2))

	fmt.Printf("%s \n%s \n", hex.EncodeToString(scAddress1), hex.EncodeToString(scAddress2))
}
