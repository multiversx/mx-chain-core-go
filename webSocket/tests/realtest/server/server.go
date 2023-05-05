package main

import (
	"fmt"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/block"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/multiversx/mx-chain-core-go/webSocket/factory"
	logger "github.com/multiversx/mx-chain-logger-go"
)

var (
	jsonMarshaller = &marshal.JsonMarshalizer{}
	log            = logger.GetOrCreate("test-server")
)

func main() {
	server, err := createServer()
	if err != nil {
		fmt.Println("cannot create server: ", err.Error())
		return
	}

	timeoutChan := make(chan bool)
	go func(tChan chan bool) {
		time.Sleep(1 * time.Minute)
		tChan <- true
	}(timeoutChan)

	funcCloseServer := func() {
		err = server.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	for {
		select {
		case <-timeoutChan:
			funcCloseServer()
		default:
			time.Sleep(2 * time.Second)
			doAction(server)
		}
	}
}

func doAction(server Driver) {
	metaHeader := &block.MetaBlock{
		Nonce: 100,
	}
	metaHeaderBytes, _ := jsonMarshaller.Marshal(metaHeader)

	err := server.SaveBlock(&outport.OutportBlock{
		BlockData: &outport.BlockData{
			HeaderType:  string(core.MetaHeader),
			HeaderHash:  []byte("header hash"),
			HeaderBytes: metaHeaderBytes,
		},
	},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = server.SaveAccounts(&outport.Accounts{BlockTimestamp: 1155})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = server.FinalizedBlock(&outport.FinalizedBlock{HeaderHash: []byte("reverted header hash")})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = server.SaveRoundsInfo(nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createServer() (Driver, error) {
	return factory.NewWebSocketDriver(factory.ArgsWebSocketDriverFactory{
		Marshaller: jsonMarshaller,
		WebSocketConfig: data.WebSocketConfig{
			URL:                "127.0.0.1:21112",
			RetryDurationInSec: 5,
			IsServer:           true,
			WithAcknowledge:    true,
		},
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      log,
	})
}
