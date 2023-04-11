package clientServerSender

// MessageSender defines what a message sender should be able to do
type MessageSender interface {
	Send(counter uint64, payload []byte) error
	Close() error
}
