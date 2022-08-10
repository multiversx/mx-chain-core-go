package core

import (
	"time"

	"github.com/ElrondNetwork/elrond-go-core/data"
)

// AppStatusHandler interface will handle different implementations of monitoring tools, such as term-ui or status metrics
type AppStatusHandler interface {
	IsInterfaceNil() bool
	Increment(key string)
	AddUint64(key string, val uint64)
	Decrement(key string)
	SetInt64Value(key string, value int64)
	SetUInt64Value(key string, value uint64)
	SetStringValue(key string, value string)
	Close()
}

// ConnectedAddressesHandler interface will be used for passing the network component to AppStatusPolling
type ConnectedAddressesHandler interface {
	ConnectedAddresses() []string
}

// PubkeyConverter can convert public key bytes to/from a human readable form
type PubkeyConverter interface {
	Len() int
	Decode(humanReadable string) ([]byte, error)
	Encode(pkBytes []byte) string
	IsInterfaceNil() bool
}

// TimersScheduler exposes functionality for scheduling multiple timers
type TimersScheduler interface {
	Add(callback func(alarmID string), duration time.Duration, alarmID string)
	Cancel(alarmID string)
	Close()
	Reset(alarmID string)
	IsInterfaceNil() bool
}

// NodeTypeProviderHandler defines the actions needed for a component that can handle the node type
type NodeTypeProviderHandler interface {
	SetType(nodeType NodeType)
	GetType() NodeType
	IsInterfaceNil() bool
}

// WatchdogTimer is used to set alarms for different components
type WatchdogTimer interface {
	Set(callback func(alarmID string), duration time.Duration, alarmID string)
	SetDefault(duration time.Duration, alarmID string)
	Stop(alarmID string)
	Reset(alarmID string)
	IsInterfaceNil() bool
}

// Throttler can monitor the number of the currently running go routines
type Throttler interface {
	CanProcess() bool
	StartProcessing()
	EndProcessing()
	IsInterfaceNil() bool
}

// KeyValueHolder is used to hold a key and an associated value
type KeyValueHolder interface {
	Key() []byte
	Value() []byte
	ValueWithoutSuffix(suffix []byte) ([]byte, error)
}

// EpochSubscriberHandler defines the behavior of a component that can be notified if a new epoch was confirmed
type EpochSubscriberHandler interface {
	EpochConfirmed(epoch uint32, timestamp uint64)
	IsInterfaceNil() bool
}

// Accumulator defines the interface able to accumulate data and periodically evict them
type Accumulator interface {
	AddData(data interface{})
	OutputChannel() <-chan []interface{}
	Close() error
	IsInterfaceNil() bool
}

// GasScheduleSubscribeHandler defines the behavior of a component that can be notified if a the gas schedule was changed
type GasScheduleSubscribeHandler interface {
	GasScheduleChange(gasSchedule map[string]map[string]uint64)
	IsInterfaceNil() bool
}

// EpochNotifier can notify upon an epoch change and provide the current epoch
type EpochNotifier interface {
	RegisterNotifyHandler(handler EpochSubscriberHandler)
	IsInterfaceNil() bool
}

// GasScheduleNotifier can notify upon a gas schedule change
type GasScheduleNotifier interface {
	RegisterNotifyHandler(handler GasScheduleSubscribeHandler)
	LatestGasSchedule() map[string]map[string]uint64
	UnRegisterAll()
	IsInterfaceNil() bool
}

// Queue is an interface for queue implementations that evict the first element when the queue is full
type Queue interface {
	Add(hash []byte) []byte
}

// SafeCloser represents a subcomponent used for signaling a closing event. Its Close method is considered to be
// concurrent safe
type SafeCloser interface {
	Close()
	ChanClose() <-chan struct{}
	IsInterfaceNil() bool
}

// Logger defines the behavior of a data logger component
type Logger interface {
	Trace(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	LogIfError(err error, args ...interface{})
	IsInterfaceNil() bool
}

// MessageP2P defines what a p2p message can do (should return)
type MessageP2P interface {
	From() []byte
	Data() []byte
	Payload() []byte
	SeqNo() []byte
	Topic() string
	Signature() []byte
	Key() []byte
	Peer() PeerID
	Timestamp() int64
	IsInterfaceNil() bool
}

// InterceptedDebugger defines an interface for debugging the intercepted data
type InterceptedDebugger interface {
	LogReceivedHashes(topic string, hashes [][]byte)
	LogProcessedHashes(topic string, hashes [][]byte, err error)
	IsInterfaceNil() bool
}

// Interceptor defines what a data interceptor should do
// It should also adhere to the p2p.MessageProcessor interface so it can wire to a p2p.Messenger
type Interceptor interface {
	ProcessReceivedMessage(message MessageP2P, fromConnectedPeer PeerID) error
	SetInterceptedDebugHandler(handler InterceptedDebugger) error
	RegisterHandler(handler func(topic string, hash []byte, data interface{}))
	Close() error
	IsInterfaceNil() bool
}

// Validator defines a node that can be allocated to a shard for participation in a consensus group as validator
// or block proposer
type Validator interface {
	PubKey() []byte
	Chances() uint32
	Index() uint32
	Size() int
}

// EpochStartActionHandler defines the action taken on epoch start event
type EpochStartActionHandler interface {
	EpochStartAction(hdr data.HeaderHandler)
	EpochStartPrepare(metaHdr data.HeaderHandler, body data.BodyHandler)
	NotifyOrder() uint32
}
