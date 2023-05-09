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
