package data

// PayloadData holds the arguments that should be parsed from a websocket payload
type PayloadData struct {
	WithAcknowledge bool
	Counter         uint64
	OperationType   OperationType
	Payload         []byte
}
