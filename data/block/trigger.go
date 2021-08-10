//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. trigger.proto
package block

import "github.com/ElrondNetwork/elrond-go-core/data"

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (tr *TriggerRegistry) GetEpochStartHeaderHandler() data.HeaderHandler {
	return tr.GetEpochStartShardHeader()
}

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (trV2 *TriggerRegistryV2) GetEpochStartHeaderHandler() data.HeaderHandler {
	return trV2.GetEpochStartShardHeader()
}
