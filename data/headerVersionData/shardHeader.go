package headerVersionData

import "math/big"

// HeaderAdditionalData holds getters for the additional version related header data for new versions
// for future versions this interface can grow
type HeaderAdditionalData interface {
	GetScheduledRootHash() []byte
	GetScheduledAccumulatedFees() *big.Int
	GetScheduledDeveloperFees() *big.Int
	GetGasProvided() uint64
	GetGasPenalized() uint64
	GetGasRefunded() uint64
	IsInterfaceNil() bool
}

// AdditionalData holds the additional version related header data
// for future header versions this structure can grow
type AdditionalData struct {
	ScheduledRootHash        []byte
	ScheduledAccumulatedFees *big.Int
	ScheduledDeveloperFees   *big.Int
	GasProvided              uint64
	GasPenalized             uint64
	GasRefunded              uint64
}

// GetScheduledRootHash returns the scheduled RootHash
func (ad *AdditionalData) GetScheduledRootHash() []byte {
	if ad == nil {
		return nil
	}

	return ad.ScheduledRootHash
}

// GetScheduledAccumulatedFees returns the accumulated fees on scheduled SC calls
func (ad *AdditionalData) GetScheduledAccumulatedFees() *big.Int {
	if ad == nil {
		return nil
	}

	return ad.ScheduledAccumulatedFees
}

// GetScheduledDeveloperFees returns the developer fees on scheduled SC calls
func (ad *AdditionalData) GetScheduledDeveloperFees() *big.Int {
	if ad == nil {
		return nil
	}

	return ad.ScheduledDeveloperFees
}

// GetGasProvided returns the gas provided on scheduled SC calls for previous block
func (ad *AdditionalData) GetGasProvided() uint64 {
	if ad == nil {
		return 0
	}

	return ad.GasProvided
}

// GetGasPenalized returns the gas penalized on scheduled SC calls for previous block
func (ad *AdditionalData) GetGasPenalized() uint64 {
	if ad == nil {
		return 0
	}

	return ad.GasPenalized
}

// GetGasRefunded returns the gas refunded on scheduled SC calls for previous block
func (ad *AdditionalData) GetGasRefunded() uint64 {
	if ad == nil {
		return 0
	}

	return ad.GasRefunded
}

// IsInterfaceNil returns true if there is no value under the interface
func (ad *AdditionalData) IsInterfaceNil() bool {
	return ad == nil
}
