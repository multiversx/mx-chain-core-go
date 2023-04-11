package clientServerReceiver

// WsMessagesReceiver defines what a messages receiver should be able to do
type WsMessagesReceiver interface {
	Start()
	Close()
}
