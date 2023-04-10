package clientServerSender

type MessageSender interface {
	Send(counter uint64, payload []byte) error
	Close() error
}
