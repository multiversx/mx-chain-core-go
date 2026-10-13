package mock

import "github.com/multiversx/mx-chain-core-go/data/outport"

// OutportHandlerStub is a stub implementation of the gRPC outport handler.
type OutportHandlerStub struct {
	SaveBlockCalled             func(in *outport.OutportBlock) error
	RevertIndexedBlockCalled    func(in *outport.BlockData) error
	SaveRoundsInfoCalled        func(in *outport.RoundsInfo) error
	SaveValidatorsPubKeysCalled func(in *outport.ValidatorsPubKeys) error
	SaveValidatorsRatingCalled  func(in *outport.ValidatorsRating) error
	SaveAccountsCalled          func(in *outport.Accounts) error
	FinalizedBlockCalled        func(in *outport.FinalizedBlock) error
}

// SaveBlock -
func (stub *OutportHandlerStub) SaveBlock(in *outport.OutportBlock) error {
	return stub.SaveBlockCalled(in)
}

// RevertIndexedBlock -
func (stub *OutportHandlerStub) RevertIndexedBlock(in *outport.BlockData) error {
	return stub.RevertIndexedBlockCalled(in)
}

// SaveRoundsInfo -
func (stub *OutportHandlerStub) SaveRoundsInfo(in *outport.RoundsInfo) error {
	return stub.SaveRoundsInfoCalled(in)
}

// SaveValidatorsPubKeys -
func (stub *OutportHandlerStub) SaveValidatorsPubKeys(in *outport.ValidatorsPubKeys) error {
	return stub.SaveValidatorsPubKeysCalled(in)
}

// SaveValidatorsRating -
func (stub *OutportHandlerStub) SaveValidatorsRating(in *outport.ValidatorsRating) error {
	return stub.SaveValidatorsRatingCalled(in)
}

// SaveAccounts -
func (stub *OutportHandlerStub) SaveAccounts(in *outport.Accounts) error {
	return stub.SaveAccountsCalled(in)
}

// FinalizedBlock -
func (stub *OutportHandlerStub) FinalizedBlock(in *outport.FinalizedBlock) error {
	return stub.FinalizedBlockCalled(in)
}

// IsInterfaceNil -
func (stub *OutportHandlerStub) IsInterfaceNil() bool {
	return stub == nil
}
