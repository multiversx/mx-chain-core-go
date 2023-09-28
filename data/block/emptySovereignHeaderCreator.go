package block

import "github.com/multiversx/mx-chain-core-go/data"

type emptySovereignHeaderCreator struct{}

// NewEmptySovereignHeaderCreator is able to create empty sovereign header instances
func NewEmptySovereignHeaderCreator() *emptySovereignHeaderCreator {
	return &emptySovereignHeaderCreator{}
}

// CreateNewHeader creates a new empty sovereign header
func (creator *emptySovereignHeaderCreator) CreateNewHeader() data.HeaderHandler {
	return &SovereignChainHeader{}
}

// IsInterfaceNil returns true if there is no value under the interface
func (creator *emptySovereignHeaderCreator) IsInterfaceNil() bool {
	return creator == nil
}
