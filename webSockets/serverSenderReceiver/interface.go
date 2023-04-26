package serverSenderReceiver

import "github.com/multiversx/mx-chain-core-go/webSockets/utils"

type ReceiversHolder interface {
	AddReceiver(id string, rec utils.Receiver)
	RemoveReceiver(id string)
	GetAll() map[string]utils.Receiver
}
