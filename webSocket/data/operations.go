package data

const (
	// WSRoute is the route which data will be sent over websocket
	WSRoute = "/save"
)

// WebSocketConfig holds the configuration needed for instantiating a new web socket server
type WebSocketConfig struct {
	URL                string
	WithAcknowledge    bool
	IsServer           bool
	RetryDurationInSec int
	BlockingAckOnError bool
}

// PayloadType defines the type to be used to group web socket operations
type PayloadType uint8

// PayloadTypeFromUint64 returns the payload type based on the provided uint64 value
func PayloadTypeFromUint64(value uint64) PayloadType {
	return PayloadType(uint8(value))
}

// String will return the string representation of the operation
func (ot PayloadType) String() string {
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
func (ot PayloadType) Uint32() uint32 {
	return uint32(ot)
}

const (
	// PayloadSaveBlock is the payload type that triggers a block saving
	PayloadSaveBlock PayloadType = 0
	// PayloadRevertIndexedBlock is the payload type that triggers a reverting of an indexed block
	PayloadRevertIndexedBlock PayloadType = 1
	// PayloadSaveRoundsInfo is the payload type that triggers the saving of rounds info
	PayloadSaveRoundsInfo PayloadType = 2
	// PayloadSaveValidatorsPubKeys is the payload type that triggers the saving of validators' public keys
	PayloadSaveValidatorsPubKeys PayloadType = 3
	// PayloadSaveValidatorsRating is the payload type that triggers the saving of the validators' rating
	PayloadSaveValidatorsRating PayloadType = 4
	// PayloadSaveAccounts is the payload type that triggers the saving of accounts
	PayloadSaveAccounts PayloadType = 5
	// PayloadFinalizedBlock is the payload type that triggers the handling of a finalized block
	PayloadFinalizedBlock PayloadType = 6
)
