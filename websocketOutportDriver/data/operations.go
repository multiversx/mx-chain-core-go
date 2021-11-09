package data

type WebSocketConfig struct {
	URL             string
	WithAcknowledge bool
}

type OperationType uint16

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

func (ot OperationType) Uint32() uint32 {
	return uint32(ot)
}

const (
	OperationSaveBlock             OperationType = 0
	OperationRevertIndexedBlock    OperationType = 1
	OperationSaveRoundsInfo        OperationType = 2
	OperationSaveValidatorsPubKeys OperationType = 3
	OperationSaveValidatorsRating  OperationType = 4
	OperationSaveAccounts          OperationType = 5
	OperationFinalizedBlock        OperationType = 6
)
