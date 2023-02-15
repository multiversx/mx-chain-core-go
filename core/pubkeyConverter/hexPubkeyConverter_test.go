package pubkeyConverter_test

import (
	"errors"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/core/pubkeyConverter"
	"github.com/stretchr/testify/assert"
)

func TestNewHexPubkeyConverter_InvalidSizeShouldErr(t *testing.T) {
	t.Parallel()

	hpc, err := pubkeyConverter.NewHexPubkeyConverter(-1)
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(hpc))

	hpc, err = pubkeyConverter.NewHexPubkeyConverter(0)
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(hpc))

	hpc, err = pubkeyConverter.NewHexPubkeyConverter(3)
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(hpc))
}

func TestNewHexPubkeyConverter_ShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 28
	hpc, err := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	assert.Nil(t, err)
	assert.False(t, check.IfNil(hpc))
	assert.Equal(t, addressLen, hpc.Len())
}

func TestHexPubkeyConverter_DecodeShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 4
	hpc, _ := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	buff, err := hpc.Decode("aaff")
	assert.True(t, errors.Is(err, pubkeyConverter.ErrWrongSize))
	assert.Equal(t, 0, len(buff))

	buff, err = hpc.Decode("not a hex")
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(buff))
}

func TestHexPubkeyConverter_DecodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 2
	hpc, _ := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	buff, err := hpc.Decode("aaff")

	assert.Nil(t, err)
	assert.Equal(t, []byte{170, 255}, buff)
}

func TestHexPubkeyConverter_EncodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 4
	hpc, _ := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	str, err := hpc.Encode([]byte{170, 255})
	assert.Nil(t, err)
	assert.Equal(t, "aaff", str)
}

func TestHexPubkeyConverter_SilentEncodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 4
	hpc, _ := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	str := hpc.SilentEncode([]byte{170, 255}, &mock.LoggerMock{})
	assert.Equal(t, "aaff", str)
}

func TestHexPubkeyConverter_EncodeDecodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 16
	hpc, _ := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	value := "123456789012345678901234567890af"
	buff, err := hpc.Decode(value)
	assert.Nil(t, err)

	recoveredValue, err := hpc.Encode(buff)
	assert.Nil(t, err)
	assert.Equal(t, value, recoveredValue)
}

func TestHexPubkeyConverter_EncodeSliceShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 16
	sliceLen := 2

	hpc, _ := pubkeyConverter.NewHexPubkeyConverter(addressLen)

	decodedSlice := make([][]byte, 0)

	hexPubkey1, _ := hpc.Decode("123456789012345678901234567890af")
	decodedSlice = append(decodedSlice, hexPubkey1)

	hexPubkey2, _ := hpc.Decode("123456789012345678901234567890af")
	decodedSlice = append(decodedSlice, hexPubkey2)

	str, err := hpc.EncodeSlice(decodedSlice)
	assert.Nil(t, err)
	assert.Equal(t, sliceLen, len(str))
	assert.Equal(t, []string{"123456789012345678901234567890af",
		"123456789012345678901234567890af"}, str)
}
