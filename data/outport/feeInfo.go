package outport

import "math/big"

// SetInitialPaidFee sets the initial paid fee
func (f *FeeInfo) SetInitialPaidFee(fee *big.Int) {
	f.InitialPaidFee = fee
}

// SetGasUsed sets the used gas
func (f *FeeInfo) SetGasUsed(gasUsed uint64) {
	f.GasUsed = gasUsed
}

// SetFee sets the fee
func (f *FeeInfo) SetFee(fee *big.Int) {
	f.Fee = fee
}
