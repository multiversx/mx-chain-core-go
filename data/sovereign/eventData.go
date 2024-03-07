package sovereign

// TransferData holds the required data for a transfer to smart contract
type TransferData struct {
	GasLimit uint64
	Function []byte
	Args     [][]byte
}

// EventData holds the full event data structure
type EventData struct {
	Nonce uint64
	*TransferData
}
