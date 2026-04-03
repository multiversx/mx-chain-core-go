package block

import "github.com/multiversx/mx-chain-core-go/data"

type emptyHeaderV3Creator struct{}

// NewEmptyHeaderV3Creator is able to create empty header v3 instances
func NewEmptyHeaderV3Creator() *emptyHeaderV3Creator {
	return &emptyHeaderV3Creator{}
}

// CreateNewHeader creates a new empty header v3
func (creator *emptyHeaderV3Creator) CreateNewHeader() data.HeaderHandler {
	return &HeaderV3{}
}

// IsInterfaceNil returns true if there is no value under the interface
func (creator *emptyHeaderV3Creator) IsInterfaceNil() bool {
	return creator == nil
}
