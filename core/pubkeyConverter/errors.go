package pubkeyConverter

import "errors"

// ErrInvalidAddressLength signals that address length is invalid
var ErrInvalidAddressLength = errors.New("invalid address length")

// ErrWrongSize signals that a wrong size occurred
var ErrWrongSize = errors.New("wrong size")

// ErrInvalidErdAddress signals that the provided address is not an ERD address
var ErrInvalidErdAddress = errors.New("invalid ERD address")

// ErrBech32ConvertError signals that conversion the 5bit alphabet to 8bit failed
var ErrBech32ConvertError = errors.New("can't convert bech32 string")

// ErrHrpPrefix signals that the prefix is not human readable or empty
var ErrInvalidHrpPrefix = errors.New("invalid hrp prefix")

// ErrConvertBits signals that a configuration mistake has been introduced
var ErrConvertBits = errors.New("invalid fromBits or toBits when converting bits")
