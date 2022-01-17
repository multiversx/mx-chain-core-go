package headerVersionData

import "math/big"

// HeaderAdditionalData holds getters for the additional version related header data for new versions
// for future versions this interface can grow
type HeaderAdditionalData interface {
	GetScheduledRootHash() []byte
	GetScheduledAccumulatedFees() *big.Int
	GetScheduledDeveloperFees() *big.Int
	GetScheduledGasProvided() uint64
	GetScheduledGasPenalized() uint64
	GetScheduledGasRefunded() uint64
	IsInterfaceNil() bool
}

// AdditionalData holds the additional version related header data
// for future header versions this structure can grow
type AdditionalData struct {
	ScheduledRootHash        []byte
	ScheduledAccumulatedFees *big.Int
	ScheduledDeveloperFees   *big.Int
	ScheduledGasProvided     uint64
	ScheduledGasPenalized    uint64
	ScheduledGasRefunded     uint64
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
	if ad.ScheduledAccumulatedFees == nil {
		return big.NewInt(0)
	}
	return big.NewInt(0).Set(ad.ScheduledAccumulatedFees)
}

// GetScheduledDeveloperFees returns the developer fees on scheduled SC calls
func (ad *AdditionalData) GetScheduledDeveloperFees() *big.Int {
	if ad == nil {
		return nil
	}
	if ad.ScheduledDeveloperFees == nil {
		return big.NewInt(0)
	}

	return big.NewInt(0).Set(ad.ScheduledDeveloperFees)
}

// GetScheduledGasProvided returns the gas provided on scheduled SC calls for previous block
func (ad *AdditionalData) GetScheduledGasProvided() uint64 {
	if ad == nil {
		return 0
	}

	return ad.ScheduledGasProvided
}

// GetScheduledGasPenalized returns the gas penalized on scheduled SC calls for previous block
func (ad *AdditionalData) GetScheduledGasPenalized() uint64 {
	if ad == nil {
		return 0
	}

	return ad.ScheduledGasPenalized
}

// GetScheduledGasRefunded returns the gas refunded on scheduled SC calls for previous block
func (ad *AdditionalData) GetScheduledGasRefunded() uint64 {
	if ad == nil {
		return 0
	}

	return ad.ScheduledGasRefunded
}

// IsInterfaceNil returns true if there is no value under the interface
func (ad *AdditionalData) IsInterfaceNil() bool {
	return ad == nil
}
