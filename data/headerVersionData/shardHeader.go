package headerVersionData

import "math/big"

// HeaderAdditionalData holds getters for the additional version related header data for new versions
// for future versions this interface can grow
type HeaderAdditionalData interface {
	GetScheduledRootHash() []byte
	GetScheduledAccumulatedFees() *big.Int
	GetScheduledDeveloperFees() *big.Int
	IsInterfaceNil() bool
}

// AdditionalData holds the additional version related header data
// for future header versions this structure can grow
type AdditionalData struct {
	ScheduledRootHash        []byte
	ScheduledAccumulatedFees *big.Int
	ScheduledDeveloperFees   *big.Int
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

// IsInterfaceNil returns true if there is no value under the interface
func (ad *AdditionalData) IsInterfaceNil() bool {
	return ad == nil
}
