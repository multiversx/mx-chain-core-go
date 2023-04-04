package pubkeyConverter

import (
	"encoding/hex"
	"fmt"
	"runtime/debug"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
)

type config struct {
	prefix   string
	fromBits byte
	toBits   byte
	pad      bool
}

var bech32Config = config{
	fromBits: byte(8),
	toBits:   byte(5),
	pad:      true,
}

// bech32PubkeyConverter encodes or decodes provided public key as/from bech32 format
type bech32PubkeyConverter struct {
	prefix string
	len    int
}

// NewBech32PubkeyConverter returns a bech32PubkeyConverter instance
func NewBech32PubkeyConverter(addressLen int, prefix string) (*bech32PubkeyConverter, error) {
	if addressLen < 1 {
		return nil, fmt.Errorf("%w when creating address converter, addressLen should have been greater than 0",
			ErrInvalidAddressLength)
	}
	if addressLen%2 == 1 {
		return nil, fmt.Errorf("%w when creating address converter, addressLen should have been an even number",
			ErrInvalidAddressLength)
	}
	if !check.IfHrp(prefix) {
		return nil, fmt.Errorf("%w when creating address converter, prefix should have been human readable",
			ErrInvalidHrpPrefix)
	}

	return &bech32PubkeyConverter{
		prefix: prefix,
		len:    addressLen,
	}, nil
}

// Len returns the decoded address length
func (bpc *bech32PubkeyConverter) Len() int {
	return bpc.len
}

// Decode converts the provided public key string as bech32 decoded bytes
func (bpc *bech32PubkeyConverter) Decode(humanReadable string) ([]byte, error) {
	decodedPrefix, buff, err := bech32.Decode(humanReadable)
	if err != nil {
		return nil, err
	}
	if decodedPrefix != bpc.prefix {
		return nil, ErrInvalidErdAddress
	}

	// warning: mind the order of the parameters, those should be inverted
	decodedBytes, err := bech32.ConvertBits(buff, bech32Config.toBits, bech32Config.fromBits, !bech32Config.pad)
	if err != nil {
		return nil, ErrBech32ConvertError
	}

	if len(decodedBytes) != bpc.len {
		return nil, fmt.Errorf("%w when decoding address, expected length %d, received %d",
			ErrWrongSize, bpc.len, len(decodedBytes))
	}

	return decodedBytes, nil
}

// Encode converts the provided bytes in a bech32 form
func (bpc *bech32PubkeyConverter) Encode(pkBytes []byte) (string, error) {
	if len(pkBytes) != bpc.len {
		return "", fmt.Errorf("%w when encoding address, expected length %d, received %d",
			ErrWrongSize, bpc.len, len(pkBytes))
	}

	conv, err := bech32.ConvertBits(pkBytes, bech32Config.fromBits, bech32Config.toBits, bech32Config.pad)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrConvertBits, err.Error())
	}

	encodedBytes, err := bech32.Encode(bpc.prefix, conv)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrBech32ConvertError, err.Error())
	}

	return encodedBytes, nil
}

// SilentEncode converts the provided bytes in a bech32 form without returning any error
func (bpc *bech32PubkeyConverter) SilentEncode(pkBytes []byte, log core.Logger) string {
	encodedBytes, err := bpc.Encode(pkBytes)
	if err != nil {
		log.Warn("bech32PubkeyConverter.SilentEncode",
			"hex buff", hex.EncodeToString(pkBytes),
			"error", err,
			"stack trace", string(debug.Stack()))
		return ""
	}

	return encodedBytes
}

// EncodeSlice converts the provided bytes slice into a slice of bech32 addresses
func (bpc *bech32PubkeyConverter) EncodeSlice(pkBytesSlice [][]byte) ([]string, error) {
	encodedSlice := make([]string, 0, len(pkBytesSlice))

	for _, item := range pkBytesSlice {
		encoded, err := bpc.Encode(item)
		if err != nil {
			return nil, err
		}
		encodedSlice = append(encodedSlice, encoded)
	}

	return encodedSlice, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (bpc *bech32PubkeyConverter) IsInterfaceNil() bool {
	return bpc == nil
}
