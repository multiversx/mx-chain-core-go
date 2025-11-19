package block

import "github.com/multiversx/mx-chain-core-go/data"

type emptyMetaBlockV3Creator struct{}

// NewEmptyMetaBlockV3Creator is able to create empty MetaBlockV3 instances
func NewEmptyMetaBlockV3Creator() *emptyMetaBlockV3Creator {
	return &emptyMetaBlockV3Creator{}
}

// CreateNewHeader creates a new empty metablock v3
func (creator *emptyMetaBlockV3Creator) CreateNewHeader() data.HeaderHandler {
	return &MetaBlockV3{}
}

// IsInterfaceNil returns true if there is no value under the interface
func (creator *emptyMetaBlockV3Creator) IsInterfaceNil() bool {
	return creator == nil
}
