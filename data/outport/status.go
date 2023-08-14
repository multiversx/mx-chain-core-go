package outport

// StatusInfo holds the fields for the transaction status
type StatusInfo struct {
	CompletedEvent bool   `json:"completedEvent"`
	ErrorEvent     bool   `json:"errorEvent"`
	Status         string `json:"status"`
}
