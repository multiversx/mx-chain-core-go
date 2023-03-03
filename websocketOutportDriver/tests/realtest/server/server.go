package main

import (
	"fmt"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	"github.com/multiversx/mx-chain-core-go/marshal"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/factory"
)

var jsonMarshaller = &marshal.JsonMarshalizer{}

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
	err := server.SaveBlock(&outport.OutportBlock{BlockData: &outport.BlockData{HeaderHash: []byte("header hash")}})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = server.SaveAccounts(nil)
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
	wsFactory, err := factory.NewOutportDriverWebSocketSenderFactory(factory.OutportDriverWebSocketSenderFactoryArgs{
		Marshaller: jsonMarshaller,
		WebSocketConfig: data.WebSocketConfig{
			URL: "127.0.0.1:21112",
		},
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      &mock.LoggerMock{},
		WithAcknowledge:          true,
	})
	if err != nil {
		return nil, err
	}

	return wsFactory.Create()
}
