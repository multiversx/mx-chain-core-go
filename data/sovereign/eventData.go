package sovereign

type TransferData struct {
	GasLimit uint64
	Function []byte
	Args     [][]byte
}

type EventData struct {
	Nonce uint64
	*TransferData
}
