package block

import (
	"sync"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data"
)

type emptyBlockCreatorsContainer struct {
	mut           sync.RWMutex
	blockCreators map[core.HeaderType]EmptyBlockCreator
}

// NewEmptyBlockCreatorsContainer creates a new block creators container
func NewEmptyBlockCreatorsContainer() *emptyBlockCreatorsContainer {
	return &emptyBlockCreatorsContainer{
		blockCreators: make(map[core.HeaderType]EmptyBlockCreator),
	}
}

// Add will add a new empty block creator
func (container *emptyBlockCreatorsContainer) Add(headerType core.HeaderType, creator EmptyBlockCreator) error {
	if check.IfNil(creator) {
		return data.ErrNilEmptyBlockCreator
	}

	container.mut.Lock()
	container.blockCreators[headerType] = creator
	container.mut.Unlock()

	return nil
}

// Get will try to get a new empty block creator. Errors if the empty block creator type is not found
func (container *emptyBlockCreatorsContainer) Get(headerType core.HeaderType) (EmptyBlockCreator, error) {
	container.mut.RLock()
	creator, ok := container.blockCreators[headerType]
	container.mut.RUnlock()

	if !ok {
		return nil, data.ErrInvalidHeaderType
	}

	return creator, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (container *emptyBlockCreatorsContainer) IsInterfaceNil() bool {
	return container == nil
}
