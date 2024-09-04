package data

// LogData holds the data needed for indexing logs and events
type LogData struct {
	LogHandler
	TxHash string
}

// KeyValuePair is a tuple of (key, value)
type KeyValuePair struct {
	Key   []byte
	Value []byte
}

// HeaderProof holds an aggregated signature and the public keys bitmap
type HeaderProof struct {
	AggregatedSignature []byte
	PubKeysBitmap       []byte
	HeaderHash          []byte
	HeaderEpoch         uint32
}
