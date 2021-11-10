package main

import (
	"fmt"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/indexer"
	"github.com/ElrondNetwork/elrond-go-core/data/typeConverters/uint64ByteSlice"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/data"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/factory"
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
		time.Sleep(5 * time.Minute)
		tChan <- true
	}(timeoutChan)

	for {
		select {
		case <-timeoutChan:
			return
		default:
			time.Sleep(1 * time.Second)
			doAction(server)
		}
	}
}

func doAction(server Driver) {
	fmt.Println("called SaveBlock")
	err := server.SaveBlock(&indexer.ArgsSaveBlockData{HeaderHash: []byte("header hash")})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = server.SaveAccounts(1155, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = server.FinalizedBlock([]byte("reverted header hash"))
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
			URL: "127.0.0.1:21111",
		},
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		Log:                      &mock.LoggerMock{},
	})
	if err != nil {
		return nil, err
	}

	return wsFactory.Create()
}