package addressGenerator

import "errors"

// ErrAddressLengthNotCorrect signals that an account does not have the correct address
var ErrAddressLengthNotCorrect = errors.New("address length is not correct")

// ErrVMTypeLengthIsNotCorrect signals that the vm type length is not correct
var ErrVMTypeLengthIsNotCorrect = errors.New("vm type length is not correct")
