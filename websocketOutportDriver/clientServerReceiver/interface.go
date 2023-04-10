package clientServerReceiver

type WsMessagesReceiver interface {
	Start()
	Close()
}
