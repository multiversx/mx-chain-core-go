package integrationTests

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/stretchr/testify/require"
)

func TestStartServerAddClientAndSendData(t *testing.T) {
	url := "localhost:8833"
	wsServer, err := createServer(url, &mock.LoggerMock{})
	require.Nil(t, err)

	wg := &sync.WaitGroup{}

	wg.Add(1)

	_ = wsServer.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte) error {
			require.Equal(t, []byte("test"), payload)
			wg.Done()
			return nil
		},
	})

	wsServer.Start()

	wsClient1, err := createClient(url)
	require.Nil(t, err)

	wsClient1.Start()

	for {
		err = wsClient1.Send(data.WsSendArgs{
			Payload: []byte("test"),
			OpType:  data.OperationSaveBlock,
		})
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	require.Nil(t, err)

	_ = wsClient1.Close()
	_ = wsServer.Close()
	wg.Wait()
}

func TestStartServerAddClientAndCloseClientAndServerShouldReceiveClose(t *testing.T) {
	url := "localhost:8833"

	wg1, wg2 := &sync.WaitGroup{}, &sync.WaitGroup{}
	wg1.Add(1)
	wg2.Add(1)
	serverReceivedCloseMessage := false
	log := &mock.LoggerStub{
		InfoCalled: func(message string, args ...interface{}) {
			if strings.Contains(message, "connection closed") {
				serverReceivedCloseMessage = true
				wg2.Done()
			}
		},
	}

	wsServer, err := createServer(url, log)
	require.Nil(t, err)

	_ = wsServer.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte) error {
			require.Equal(t, []byte("test"), payload)
			wg1.Done()
			return nil
		},
	})

	wsServer.Start()

	wsClient1, err := createClient(url)
	require.Nil(t, err)
	wsClient1.Start()
	time.Sleep(time.Second)

	for {
		err = wsClient1.Send(data.WsSendArgs{
			Payload: []byte("test"),
			OpType:  data.OperationSaveBlock,
		})
		if err == nil {
			break
		}
	}

	err = wsClient1.Close()
	require.Nil(t, err)
	wg2.Wait()
	_ = wsServer.Close()
	wg1.Wait()
	require.True(t, serverReceivedCloseMessage)
}

// TODO add a scenario where server will be closed first
