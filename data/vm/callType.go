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

	// ExecOnDestByCaller means that the call is an invocation of a built in function / smart contract from
	// another smart contract but the caller is from the previous caller
	ExecOnDestByCaller
)

const (
	DirectCallStr             = "directCall"
	AsynchronousCallStr       = "asynchronousCall"
	AsynchronousCallBackStr   = "asynchronousCallBack"
	ESDTTransferAndExecuteStr = "esdtTransferAndExecute"
	ExecOnDestByCallerStr     = "execOnDestByCaller"
	UnknownStr                = "unknown"
)

func (ct CallType) ToString() string {
	switch ct {
	case DirectCall:
		return DirectCallStr
	case AsynchronousCall:
		return AsynchronousCallStr
	case AsynchronousCallBack:
		return AsynchronousCallBackStr
	case ESDTTransferAndExecute:
		return ESDTTransferAndExecuteStr
	case ExecOnDestByCaller:
		return ExecOnDestByCallerStr
	default:
		return UnknownStr
	}
}
