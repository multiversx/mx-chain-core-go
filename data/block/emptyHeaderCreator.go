package block

import "github.com/multiversx/mx-chain-core-go/data"

type emptyHeaderCreator struct{}

// NewEmptyHeaderCreator is able to create empty header v1 instances
func NewEmptyHeaderCreator() *emptyHeaderCreator {
	return &emptyHeaderCreator{}
}

// CreateNewHeader creates a new empty header v1
func (creator *emptyHeaderCreator) CreateNewHeader() data.HeaderHandler {
	return &Header{}
}

// IsInterfaceNil returns true if there is no value under the interface
func (creator *emptyHeaderCreator) IsInterfaceNil() bool {
	return creator == nil
}
