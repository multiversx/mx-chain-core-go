package data

const (
	// WSRoute is the route which data will be sent over websocket
	WSRoute = "/save"
)

// WebSocketConfig holds the configuration needed for instantiating a new web socket server
type WebSocketConfig struct {
	URL             string
	WithAcknowledge bool
}

// OperationType defines the type to be used to group web socket operations
type OperationType uint8

// OperationTypeFromUint64 returns the operation type based on the provided uint64 value
func OperationTypeFromUint64(value uint64) OperationType {
	return OperationType(uint8(value))
}

// String will return the string representation of the operation
func (ot OperationType) String() string {
	switch ot {
	case 0:
		return "SaveBlock"
	case 1:
		return "RevertIndexedBlock"
	case 2:
		return "SaveRoundsInfo"
	case 3:
		return "SaveValidatorsPubKeys"
	case 4:
		return "SaveValidatorsRating"
	case 5:
		return "SaveAccounts"
	case 6:
		return "FinalizedBlock"
	default:
		return "Unknown"
	}
}

// Uint32 will return the uint32 representation of the operation
func (ot OperationType) Uint32() uint32 {
	return uint32(ot)
}

const (
	// OperationSaveBlock is the operation that triggers a block saving
	OperationSaveBlock OperationType = 0
	// OperationRevertIndexedBlock is the operation that triggers a reverting of an indexed block
	OperationRevertIndexedBlock OperationType = 1
	// OperationSaveRoundsInfo is the operation that triggers the saving of rounds info
	OperationSaveRoundsInfo OperationType = 2
	// OperationSaveValidatorsPubKeys is the operation that triggers the saving of validators' public keys
	OperationSaveValidatorsPubKeys OperationType = 3
	// OperationSaveValidatorsRating is the operation that triggers the saving of the validators' rating
	OperationSaveValidatorsRating OperationType = 4
	// OperationSaveAccounts is the operation that triggers the saving of accounts
	OperationSaveAccounts OperationType = 5
	// OperationFinalizedBlock is the operation that triggers the handling of a finalized block
	OperationFinalizedBlock OperationType = 6
)
