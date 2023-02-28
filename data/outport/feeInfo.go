package outport

import "math/big"

func (f *FeeInfo) SetInitialPaidFee(fee *big.Int) {
	f.InitialPaidFee = fee
}

func (f *FeeInfo) SetGasUsed(gasUsed uint64) {
	f.GasUsed = gasUsed
}

func (f *FeeInfo) SetFee(fee *big.Int) {
	f.Fee = fee
}
