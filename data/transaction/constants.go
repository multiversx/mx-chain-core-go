package transaction

// TxType represents a transaction type
type TxType string

const (
	// TxTypeNormal represents the identifier for a regular transaction
	TxTypeNormal TxType = "normal"

	// TxTypeUnsigned represents the identifier for a unsigned transaction
	TxTypeUnsigned TxType = "unsigned"

	// TxTypeReward represents the identifier for a reward transaction
	TxTypeReward TxType = "reward"

	// TxTypeInvalid represents the identifier for an invalid transaction
	TxTypeInvalid TxType = "invalid"
)

// MaskSignedWithHash - this mask is used to verify if LSB from last byte from field options from transaction is set
const MaskSignedWithHash = uint32(1)

// MaskGuardedTransaction - this mask is used to verify if the flag for guarded transaction is set in the transaction option field
const MaskGuardedTransaction = uint32(1) << 1
