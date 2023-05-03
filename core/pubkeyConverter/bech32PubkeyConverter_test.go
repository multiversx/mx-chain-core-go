package pubkeyConverter_test

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/core/pubkeyConverter"
	"github.com/stretchr/testify/assert"
)

func TestNewBech32PubkeyConverter_InvalidSizeShouldErr(t *testing.T) {
	t.Parallel()

	bpc, err := pubkeyConverter.NewBech32PubkeyConverter(-1, core.DefaultAddressPrefix)
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(bpc))

	bpc, err = pubkeyConverter.NewBech32PubkeyConverter(0, core.DefaultAddressPrefix)
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(bpc))

	bpc, err = pubkeyConverter.NewBech32PubkeyConverter(3, core.DefaultAddressPrefix)
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidAddressLength))
	assert.True(t, check.IfNil(bpc))
}

func TestNewBech32PubkeyConverter_ShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 28
	bpc, err := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	assert.Nil(t, err)
	assert.False(t, check.IfNil(bpc))
	assert.Equal(t, addressLen, bpc.Len())
}

func TestBech32PubkeyConverter_DecodeInvalidStringShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	str, err := bpc.Decode("not a bech32 string")

	assert.Equal(t, 0, len(str))
	assert.NotNil(t, err)
}

func TestBech32PubkeyConverter_DecodePrefixMismatchShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	str, err := bpc.Decode("err1xyerxdp4xcmnswfsxyerxdp4xcmnswfsxyerxdp4xcmnswfsxyeqnyphvl")

	assert.Equal(t, 0, len(str))
	assert.True(t, errors.Is(err, pubkeyConverter.ErrInvalidErdAddress))
}

func TestBech32PubkeyConverter_DecodeWrongSizeShouldErr(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	str, err := bpc.Decode("erd1xyerxdp4xcmnswfsxyeqqzq40r")

	assert.Equal(t, 0, len(str))
	assert.True(t, errors.Is(err, pubkeyConverter.ErrWrongSize))
}

func TestBech32PubkeyConverter_SilentEncodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	buff := []byte("12345678901234567890123456789012")
	str := bpc.SilentEncode(buff, &mock.LoggerMock{})

	assert.Equal(t, 0, strings.Index(str, pubkeyConverter.Prefix))
}

func TestBech32PubkeyConverter_SilentEncodeShouldNotWork(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	buff := []byte("1234")
	str := bpc.SilentEncode(buff, &mock.LoggerMock{})

	assert.Equal(t, "", str)
}

func TestBech32PubkeyConverter_EncodeDecodeShouldWork(t *testing.T) {
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

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
	t.Parallel()

	addressLen := 32
	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

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
	t.Parallel()

	addressLen := 32
	sliceLen := 2

	bpc, _ := pubkeyConverter.NewBech32PubkeyConverter(addressLen, core.DefaultAddressPrefix)

	decodedSlice := make([][]byte, 0)

	alice, _ := hex.DecodeString("0139472eff6886771a982f3083da5d421f24c29181e63888228dc81ca60d69e1")
	decodedSlice = append(decodedSlice, alice)

	bob, _ := hex.DecodeString("8049d639e5a6980d1cd2392abcce41029cda74a1563523a202f09641cc2618f8")
	decodedSlice = append(decodedSlice, bob)

	str, err := bpc.EncodeSlice(decodedSlice)
	assert.Nil(t, err)
	assert.Equal(t, sliceLen, len(str))
	assert.Equal(t, []string{"erd1qyu5wthldzr8wx5c9ucg8kjagg0jfs53s8nr3zpz3hypefsdd8ssycr6th",
		"erd1spyavw0956vq68xj8y4tenjpq2wd5a9p2c6j8gsz7ztyrnpxrruqzu66jx"}, str)
}
