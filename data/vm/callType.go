package vm

// CallType specifies the type of SC invocation (in terms of asynchronicity)
type CallType int

const (
	// DirectCall means that the call is an explicit SC invocation originating from a user Transaction
	DirectCall CallType = iota

	// AsynchronousCall means that the invocation was performed from within
	// another SmartContract from another Shard, using asyncCall
	AsynchronousCall

	// AsynchronousCallBack means that an AsynchronousCall was performed
	// previously, and now the control returns to the caller SmartContract's callBack method
	AsynchronousCallBack

	// ESDTTransferAndExecute means that there is a smart contract execution after the ESDT transfer
	// this is needed in order to skip the check whether a contract is payable or not
	ESDTTransferAndExecute
)
