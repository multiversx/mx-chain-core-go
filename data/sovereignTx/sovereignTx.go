//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. sovereignTx.proto
package sovereignTx

import (
	"github.com/multiversx/mx-chain-core-go/data"
	"math/big"
)

var _ = data.TransactionHandler(&SovereignTx{})

// IsInterfaceNil verifies if underlying object is nil
func (stx *SovereignTx) IsInterfaceNil() bool {
	return stx == nil
}

// SetValue sets the value of the sovereign transaction
func (stx *SovereignTx) SetValue(value *big.Int) {
	stx.Value = value
}

// SetData sets the data of the sovereign transaction
func (stx *SovereignTx) SetData(data []byte) {
	stx.Data = data
}

// SetRcvAddr sets the receiver address of the sovereign transaction
func (stx *SovereignTx) SetRcvAddr(addr []byte) {
	stx.RcvAddr = addr
}

// SetSndAddr sets the sender address of the sovereign transaction
func (stx *SovereignTx) SetSndAddr(addr []byte) {
	stx.SndAddr = addr
}

// GetRcvUserName returns nil for sovereign transaction
func (stx *SovereignTx) GetRcvUserName() []byte {
	return nil
}

// CheckIntegrity checks the integrity of sovereign transaction's fields
func (stx *SovereignTx) CheckIntegrity() error {
	if len(stx.RcvAddr) == 0 {
		return data.ErrNilRcvAddr
	}
	if len(stx.SndAddr) == 0 {
		return data.ErrNilSndAddr
	}
	if stx.Value == nil {
		return data.ErrNilValue
	}
	if stx.Value.Sign() < 0 {
		return data.ErrNegativeValue
	}

	return nil
}
