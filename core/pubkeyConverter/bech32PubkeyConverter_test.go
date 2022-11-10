package pubkeyConverter_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/pubkeyConverter"
	"github.com/stretchr/testify/assert"
)

func TestNewBech32PubkeyConverter_InvalidSizeShouldErr(t *testing.T) {
	t.Parallel()

	bpc, err := pubkeyConverter.NewBech32PubkeyConverter(-1, "erd")
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(bpc))

	bpc, err = pubkeyConverter.NewBech32PubkeyConverter(0, "erd")
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(bpc))

	bpc, err = pubkeyConverter.NewBech32PubkeyConverter(3, "erd")
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(bpc))
}

func TestNewBech32PubkeyConverter_ShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 28
	bpc, err := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	assert.Nil(t, err)
	assert.False(t, check.IfNil(bpc))
	assert.Equal(t, addressLen, bpc.Len())
}

func TestBech32PubkeyConverter_DecodeInvalidStringShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	str, err := bpc.Decode("not a bech32 string")

	assert.Equal(t, 0, len(str))
	assert.NotNil(t, err)
}

func TestBech32PubkeyConverter_DecodePrefixMismatchShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	str, err := bpc.Decode("err1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqnyphvl")

	assert.Equal(t, 0, len(str))
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidErdAddress))
}

func TestBech32PubkeyConverter_DecodeWrongSizeShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	str, err := bpc.Decode("erd1xyerxdp4xcmnswfsxyeqqzq40r")

	assert.Equal(t, 0, len(str))
	assert.True(t, errors.Is(err, pubkeyConverter.ErrWrongSize))
}

func TestBech32PubkeyConverter_EncodeDecodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	buff := []byte("12345678901234567890123456789012")
	str, err := bpc.Encode(buff)

	assert.Nil(t, err)
	assert.Equal(t, 0, strings.Index(str, pubkeyConverter.Prefix))

	fmt.Printf("generated address: %s\n", str)

	recoveredBuff, err := bpc.Decode(str)

	assert.Nil(t, err)
	assert.Equal(t, buff, recoveredBuff)
}

func TestBech32PubkeyConverter_EncodeWrongLengthShouldReturnEmpty(t *testing.T) {
	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	buff := []byte("12345678901234567890")
	str, err := bpc.Encode(buff)

	assert.True(t, errors.Is(err, pubkeyConverter.ErrWrongSize))
	assert.Equal(t, "", str)

	buff = []byte{}
	str, err = bpc.Encode(buff)

	assert.True(t, errors.Is(err, pubkeyConverter.ErrWrongSize))
	assert.Equal(t, "", str)

	buff = []byte("1234567890123456789012345678901234567890")
	str, err = bpc.Encode(buff)

	assert.True(t, errors.Is(err, pubkeyConverter.ErrWrongSize))
	assert.Equal(t, "", str)
}

func TestBech32PubkeyConverter_EncodeSliceShouldWork(t *testing.T) {
	addressLen := 32
	sliceLn := 5
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, "erd")

	decodedSlice := make([][]byte, 0)

	buff := []byte("12345678901234567890123456789012")

	for i := 0; i < sliceLn; i++ {
		decodedSlice = append(decodedSlice, buff)
	}

	str, err := bpc.EncodeSlice(decodedSlice)
	assert.Nil(t, err)
	assert.Equal(t, sliceLn, len(str))
	assert.Equal(t, []string{"erd1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqlrqt99",
		"erd1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqlrqt99",
		"erd1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqlrqt99",
		"erd1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqlrqt99",
		"erd1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqlrqt99"}, str)
}
