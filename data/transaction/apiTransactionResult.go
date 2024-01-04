package transaction

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/vm"
)

// ApiTransactionResult is the data transfer object which will be returned on the get transaction by hash endpoint
type ApiTransactionResult struct {
	Tx                                data.TransactionHandler   `json:"-"`
	Type                              string                    `json:"type"`
	ProcessingTypeOnSource            string                    `json:"processingTypeOnSource,omitempty"`
	ProcessingTypeOnDestination       string                    `json:"processingTypeOnDestination,omitempty"`
	Hash                              string                    `json:"hash,omitempty"`
	HashBytes                         []byte                    `json:"-"`
	Nonce                             uint64                    `json:"nonce"`
	Round                             uint64                    `json:"round"`
	Epoch                             uint32                    `json:"epoch"`
	Value                             string                    `json:"value,omitempty"`
	Receiver                          string                    `json:"receiver,omitempty"`
	Sender                            string                    `json:"sender,omitempty"`
	SenderUsername                    []byte                    `json:"senderUsername,omitempty"`
	ReceiverUsername                  []byte                    `json:"receiverUsername,omitempty"`
	GasPrice                          uint64                    `json:"gasPrice,omitempty"`
	GasLimit                          uint64                    `json:"gasLimit,omitempty"`
	GasUsed                           uint64                    `json:"gasUsed,omitempty"`
	Data                              []byte                    `json:"data,omitempty"`
	CodeMetadata                      []byte                    `json:"codeMetadata,omitempty"`
	Code                              string                    `json:"code,omitempty"`
	PreviousTransactionHash           string                    `json:"previousTransactionHash,omitempty"`
	OriginalTransactionHash           string                    `json:"originalTransactionHash,omitempty"`
	ReturnMessage                     string                    `json:"returnMessage,omitempty"`
	OriginalSender                    string                    `json:"originalSender,omitempty"`
	Signature                         string                    `json:"signature,omitempty"`
	SourceShard                       uint32                    `json:"sourceShard"`
	DestinationShard                  uint32                    `json:"destinationShard"`
	BlockNonce                        uint64                    `json:"blockNonce,omitempty"`
	BlockHash                         string                    `json:"blockHash,omitempty"`
	NotarizedAtSourceInMetaNonce      uint64                    `json:"notarizedAtSourceInMetaNonce,omitempty"`
	NotarizedAtSourceInMetaHash       string                    `json:"NotarizedAtSourceInMetaHash,omitempty"`
	NotarizedAtDestinationInMetaNonce uint64                    `json:"notarizedAtDestinationInMetaNonce,omitempty"`
	NotarizedAtDestinationInMetaHash  string                    `json:"notarizedAtDestinationInMetaHash,omitempty"`
	MiniBlockType                     string                    `json:"miniblockType,omitempty"`
	MiniBlockHash                     string                    `json:"miniblockHash,omitempty"`
	HyperblockNonce                   uint64                    `json:"hyperblockNonce,omitempty"`
	HyperblockHash                    string                    `json:"hyperblockHash,omitempty"`
	Timestamp                         int64                     `json:"timestamp,omitempty"`
	Receipt                           *ApiReceipt               `json:"receipt,omitempty"`
	SmartContractResults              []*ApiSmartContractResult `json:"smartContractResults,omitempty"`
	Logs                              *ApiLogs                  `json:"logs,omitempty"`
	Status                            TxStatus                  `json:"status,omitempty"`
	Tokens                            []string                  `json:"tokens,omitempty"`
	ESDTValues                        []string                  `json:"esdtValues,omitempty"`
	Receivers                         []string                  `json:"receivers,omitempty"`
	ReceiversShardIDs                 []uint32                  `json:"receiversShardIDs,omitempty"`
	Operation                         string                    `json:"operation,omitempty"`
	Function                          string                    `json:"function,omitempty"`
	InitiallyPaidFee                  string                    `json:"initiallyPaidFee,omitempty"`
	Fee                               string                    `json:"fee,omitempty"`
	IsRelayed                         bool                      `json:"isRelayed,omitempty"`
	IsRefund                          bool                      `json:"isRefund,omitempty"`
	CallType                          string                    `json:"callType,omitempty"`
	RelayerAddress                    string                    `json:"relayerAddress,omitempty"`
	RelayedValue                      string                    `json:"relayedValue,omitempty"`
	ChainID                           string                    `json:"chainID,omitempty"`
	Version                           uint32                    `json:"version,omitempty"`
	Options                           uint32                    `json:"options"`
	GuardianAddr                      string                    `json:"guardian,omitempty"`
	GuardianSignature                 string                    `json:"guardianSignature,omitempty"`
}

// ApiSmartContractResult represents a smart contract result with changed fields' types in order to make it friendly for API's json
type ApiSmartContractResult struct {
	Hash              string      `json:"hash,omitempty"`
	Nonce             uint64      `json:"nonce"`
	Value             *big.Int    `json:"value"`
	RcvAddr           string      `json:"receiver"`
	SndAddr           string      `json:"sender"`
	RelayerAddr       string      `json:"relayerAddress,omitempty"`
	RelayedValue      *big.Int    `json:"relayedValue,omitempty"`
	Code              string      `json:"code,omitempty"`
	Data              string      `json:"data,omitempty"`
	PrevTxHash        string      `json:"prevTxHash"`
	OriginalTxHash    string      `json:"originalTxHash"`
	GasLimit          uint64      `json:"gasLimit"`
	GasPrice          uint64      `json:"gasPrice"`
	CallType          vm.CallType `json:"callType"`
	CodeMetadata      string      `json:"codeMetadata,omitempty"`
	ReturnMessage     string      `json:"returnMessage,omitempty"`
	OriginalSender    string      `json:"originalSender,omitempty"`
	Logs              *ApiLogs    `json:"logs,omitempty"`
	Tokens            []string    `json:"tokens,omitempty"`
	ESDTValues        []string    `json:"esdtValues,omitempty"`
	Receivers         []string    `json:"receivers,omitempty"`
	ReceiversShardIDs []uint32    `json:"receiversShardIDs,omitempty"`
	Operation         string      `json:"operation,omitempty"`
	Function          string      `json:"function,omitempty"`
	IsRelayed         bool        `json:"isRelayed,omitempty"`
	IsRefund          bool        `json:"isRefund,omitempty"`
}

// ApiReceipt represents a receipt with changed fields' types in order to make it friendly for API's json
type ApiReceipt struct {
	Value   *big.Int `json:"value"`
	SndAddr string   `json:"sender"`
	Data    string   `json:"data,omitempty"`
	TxHash  string   `json:"txHash"`
}

// ApiLogs represents logs with changed fields' types in order to make it friendly for API's json
type ApiLogs struct {
	Address string    `json:"address"`
	Events  []*Events `json:"events"`
}

// Events represents the events generated by a transaction with changed fields' types in order to make it friendly for API's json
type Events struct {
	Address        string   `json:"address"`
	Identifier     string   `json:"identifier"`
	Topics         [][]byte `json:"topics"`
	Data           []byte   `json:"data"`
	AdditionalData [][]byte `json:"additionalData"`
}

// CostResponse is structure used to return the transaction cost in gas units
type CostResponse struct {
	GasUnits             uint64                             `json:"txGasUnits"`
	ReturnMessage        string                             `json:"returnMessage"`
	SmartContractResults map[string]*ApiSmartContractResult `json:"smartContractResults"`
	Logs                 *ApiLogs                           `json:"logs,omitempty"`
}

// SimulationResults is the data transfer object which will hold results for simulating a transaction's execution
type SimulationResults struct {
	Hash       string                             `json:"hash,omitempty"`
	FailReason string                             `json:"failReason,omitempty"`
	Status     TxStatus                           `json:"status,omitempty"`
	ScResults  map[string]*ApiSmartContractResult `json:"scResults,omitempty"`
	Receipts   map[string]*ApiReceipt             `json:"receipts,omitempty"`
	Logs       *ApiLogs                           `json:"logs,omitempty"`
}
